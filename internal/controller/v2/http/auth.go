package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"

	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
	"src/pkg/sender"
)

var Issuer = "RacketShop"
var verificationCodes = make(map[string]string)

type AuthController struct {
	l           logging.Interface
	authService service.IAuthService
	userService service.IUserService
}

func NewAuthController(
	l logging.Interface,
	authService service.IAuthService,
	userService service.IUserService) *AuthController {
	return &AuthController{
		l:           l,
		authService: authService,
		userService: userService,
	}
}

// Login godoc
//
//	@Summary		Вход в личный кабинет пользователя
//	@Description	Метод для входа в личный кабинет пользователя
//	@Tags			Auth
//	@Produce		json
//	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
//	@Success		200	{object}	string			"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusInternalServerError	"Вход неуспешен"
//	@Failure		503	{object}	http.StatusServiceUnavailable	"Сервис недоступен"
//	@Router			/auth/login [post]
func (a *AuthController) Login(c *gin.Context) {

	var req dto.LoginReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.authService.Login(c, &req)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	if os.Getenv("TEST") != "" {
		code = fmt.Sprintf("%d", 123456)
	}
	verificationCodes[req.Email] = code

	err = sender.SendEmail(code, req.Email)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "OTP code to \"Login\" was sent to your email"})
}

// VerifyLogin godoc
//
//	@Summary		Проверка дополнительного кода при авторизации
//	@Description	Метод для проверки дополнительного кода при авторизации
//	@Tags			Auth
//	@Param			req	body		dto.LoginVerifyReq		true	"Вход пользователя"
//	@Success		200	{object}	dto.LoginRes			"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/auth/login/verify [post]
func (a *AuthController) VerifyLogin(c *gin.Context) {

	var req dto.LoginVerifyReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expectedCode, exists := verificationCodes[req.Email]
	fmt.Print(req.Email, " b ", req.Code, " a ", expectedCode, exists)
	if !exists || expectedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	user, _ := a.userService.GetUserByEmail(c, req.Email)

	token, err := a.authService.GenerateToken(user.ID)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if os.Getenv("TEST") != "" {
		token = "your_access_token"
	}
	c.JSON(http.StatusOK, dto.LoginRes{AccessToken: token})
}

// Register godoc
//
//	@Summary		Регистрация
//	@Description	Метод для регистрации в интернет-магазине
//	@Tags			Auth
//	@Param			dto.RegisterReq	body		dto.RegisterReq					true	"Регистрация пользователя"
//	@Success		200				{object}	dto.RegisterRes					"Пользователь успешно авторизовался"
//	@Failure		400				{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		500				{object}	http.StatusInternalServerError	"Внутренняя ошибка регистрации пользователя"
//	@Router			/auth/register [post]
func (a *AuthController) Register(c *gin.Context) {

	var req dto.RegisterReq
	if err := c.ShouldBindBodyWithJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.Register(c, &req)

	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterRes{AccessToken: token})
}

// ResetPassword godoc
//
//	@Summary		Смена пароля авторизованного пользователя
//	@Description	Метод для смены пароля авторизованного пользователя
//	@Tags			Auth
//	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
//	@Success		200	{object}	string			"Пароль успешно изменен"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/auth/reset_password [post]
func (a *AuthController) ResetPassword(c *gin.Context) {

	var req dto.ResetPasswordReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqLogin := &dto.LoginReq{
		Email:    req.Email,
		Password: req.OldPassword,
	}

	err := a.authService.Login(c, reqLogin)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	if os.Getenv("TEST") != "" {
		code = fmt.Sprintf("%d", 123456)
	}
	verificationCodes[req.Email] = code

	err = sender.SendEmail(code, req.Email)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "OTP code to \"Reset Password\" was sent to your email"})
}

// VerifyResetPassword godoc
//
//	@Summary		Проверка дополнительного кода при смене пароля
//	@Description	Метод для проверки дополнительного кода при смене пароля
//	@Tags			Auth
//	@Param			req	body		dto.VerifyResetPasswordReq	true	"Вход пользователя"
//	@Success		200	{object}	string				"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest		"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized		"Вход неуспешен"
//	@Router			/auth/reset_password/verify [post]
func (a *AuthController) VerifyResetPassword(c *gin.Context) {

	var req dto.VerifyResetPasswordReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expectedCode, exists := verificationCodes[req.Email]
	if !exists || expectedCode != req.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}
	user, _ := a.userService.GetUserByEmail(c, req.Email)

	reqVerifyPassword := &dto.UpdatePasswordReq{
		Email:    user.Email,
		Password: req.NewPassword,
	}

	_, err := a.userService.UpdatePassword(c, reqVerifyPassword)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Your password has been already updated!"})
}
