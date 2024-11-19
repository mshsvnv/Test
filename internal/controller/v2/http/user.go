package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/model"
	"src/internal/service"
	"src/pkg/logging"
	"src/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	l            logging.Interface
	userService  service.IUserService
	cartService  service.ICartService
	orderService service.IOrderService
}

func NewUserController(
	l logging.Interface,
	userService service.IUserService,
	cartService service.ICartService,
	orderService service.IOrderService) *UserController {
	return &UserController{
		l:            l,
		userService:  userService,
		cartService:  cartService,
		orderService: orderService,
	}
}

type UserRes struct {
	User User `json:"user"`
}

type UsersRes struct {
	User []*User `json:"users"`
}

type User struct {
	Name    string         `json:"name"`
	Surname string         `json:"surname"`
	Email   string         `json:"email"`
	Role    model.UserRole `json:"role"`
}

// GetMyProfile godoc
//
//	@Summary		Получение информации об авторизованном пользователе
//	@Description	Метод для получения информации об авторизованном пользователе
//	@Tags			User
//	@Success		200	{object}	UserRes							"Информация о конеретном пользователе"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения информации о пользователе"
//	@Security		BearerAuth
//	@Router			/profile [get]
func (u *UserController) GetMyProfile(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	user, err := u.userService.GetUserByID(c, userID)
	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var userRes User
	utils.Copy(&userRes, user)

	c.JSON(http.StatusOK, UserRes{
		User: userRes,
	})
}

// GetUserByID godoc
//
//	@Summary		Получение информации о конкретном пользователе
//	@Description	Метод для получения информации о конкретном пользователе
//	@Tags			Admin
//	@Param			id	path		int								true	"Идентификатор пользователя"
//	@Success		200	{object}	UserRes							"Информация о конеретном пользователе"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения списка ракеток"
//	@Security		BearerAuth
//	@Router			/users/{id} [get]
func (u *UserController) GetUserByID(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := u.userService.GetUserByID(c, userID)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var userRes User
	utils.Copy(&userRes, user)

	c.JSON(http.StatusOK, UserRes{
		User: userRes,
	})
}

// GetAllUsers godoc
//
//	@Summary		Получение информации о всех пользователех
//	@Description	Метод для получения информации информации о всех пользователех
//	@Tags			Admin
//	@Success		200	{object}	UsersRes						"Информация о всех пользователях"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения списка пользователей"
//	@Security		BearerAuth
//	@Router			/users [get]
func (u *UserController) GetAllUsers(c *gin.Context) {

	users, err := u.userService.GetAllUsers(c)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var usersRes []*User
	for _, user := range users {
		var userRes User
		utils.Copy(&userRes, user)
		usersRes = append(usersRes, &userRes)
	}

	c.JSON(http.StatusOK, UsersRes{
		User: usersRes,
	})
}

type UpdateReq struct {
	Role model.UserRole `json:"role"`
}

// UpdateUser godoc
//
//	@Summary		Изменение роли пользователя
//	@Description	Метод для Изменение роли пользователя
//	@Tags			Admin
//
//	@Param			id	path		int								true	"Идентефикатор пользователя"
//	@Param			req	body		UpdateReq						true	"Роль пользователя"
//
//	@Success		200	{object}	UserRes							"Информация о всех пользователях"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения списка пользователей"
//	@Security		BearerAuth
//	@Router			/users/{id} [put]
func (u *UserController) UpdateUser(c *gin.Context) {

	var req UpdateReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	user, _ := u.userService.GetUserByID(c, userID)

	user2, err := u.userService.UpdateRole(c, &dto.UpdateReq{
		Email: user.Email,
		Role:  req.Role,
	})
	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userRes User
	utils.Copy(&userRes, user2)

	c.JSON(http.StatusOK, UserRes{
		User: userRes,
	})
}
