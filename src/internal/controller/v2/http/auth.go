package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
)

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

	token, err := a.authService.Login(c, &req)
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
