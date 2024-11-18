package http

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"

	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
	// "src/pkg/auth"
)

var Issuer = "RacketShop"

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

// func (a *AuthController) Login(c *gin.Context) {

// 	var req dto.LoginReq

// 	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
// 		a.l.Infof(err.Error())
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token, err := a.authService.Login(c, &req)
// 	if err != nil {
// 		a.l.Infof(err.Error())
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.LoginRes{AccessToken: token})
// }

// Login godoc
//
//	@Summary		Вход в личный кабинет пользователя
//	@Description	Метод для входа в личный кабинет пользователя
//	@Tags			Auth
//	@Produce		jpeg
//	@Produce		json
//	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
//	@Success		200	{object}	dto.LoginRes			"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      Issuer,
		AccountName: req.Email,
	})

	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	png.Encode(&buf, img)

	c.Header("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

// VerifyLogin godoc
//
//	@Summary		Вход в личный кабинет пользователя
//	@Description	Метод для входа в личный кабинет пользователя
//	@Tags			Auth
//	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
//	@Success		200	{object}	dto.LoginRes			"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/auth/login [post]
func (a *AuthController) VerifyLogin(c *gin.Context) {

	var req dto.LoginVerifyReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keySecret, err := getKeySecret(c)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := totp.Validate(req.Code, keySecret)
	if !valid {
		a.l.Infof("error")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error"})
		return
	}

	userID, _ := getUserID(c)
	token, err := a.authService.GenerateToken(userID)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.LoginRes{AccessToken: token})
}

// Register godoc
//
//	@Summary		Вход в аккаунт пользователя
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
//	@Summary		Вход в личный кабинет пользователя
//	@Description	Метод для входа в личный кабинет пользователя
//	@Tags			Auth
//	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
//	@Success		200	{object}	dto.LoginRes			"Пользователь успешно авторизовался"
//	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
//	@Router			/auth/login [post]
func (a *AuthController) ResetPassword(c *gin.Context) {

	var req dto.ResetPasswordReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqLogin := *&dto.LoginReq{
		Email:    req.Email,
		Password: req.OldPassword,
	}

	err := a.authService.Login(c, &reqLogin)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      Issuer,
		AccountName: req.Email,
	})

	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	png.Encode(&buf, img)

	c.Header("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

// // VerifyResetPassword godoc
// //
// //	@Summary		Вход в личный кабинет пользователя
// //	@Description	Метод для входа в личный кабинет пользователя
// //	@Tags			Auth
// //	@Param			req	body		dto.LoginReq			true	"Вход пользователя"
// //	@Success		200	{object}	dto.LoginRes			"Пользователь успешно авторизовался"
// //	@Failure		400	{object}	http.StatusBadRequest	"Некорректное тело запроса"
// //	@Failure		401	{object}	http.StatusUnauthorized	"Вход неуспешен"
// //	@Router			/auth/login [post]
// func (a *AuthController) VerifyResetPassword(c *gin.Context) {

// 	var req dto.LoginReq

// 	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
// 		a.l.Infof(err.Error())
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token, err := a.authService.Login(c, &req)
// 	if err != nil {
// 		a.l.Infof(err.Error())
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, dto.LoginRes{AccessToken: token})
// }
