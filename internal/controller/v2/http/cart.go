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

type CartController struct {
	l       logging.Interface
	service service.ICartService
}

func NewCartController(l logging.Interface, service service.ICartService) *CartController {
	return &CartController{
		l:       l,
		service: service,
	}
}

type Cart struct {
	TotalPrice float32           `json:"total_price"`
	Quantity   int               `json:"quantity"`
	Lines      []*model.CartLine `json:"lines"`
}

type CartRes struct {
	Cart *Cart `json:"cart"`
}

// GetMyCart godoc
//
//	@Summary		Получение содержимого корзины
//	@Description	Метод для получения содержимого корзины
//	@Tags			User
//	@Success		200	{object}	CartRes							"Корзина благополучно получена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения корзины"
//	@Security		BearerAuth
//
//	@Router			/cart [get]
func (cc *CartController) GetMyCart(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	cart, err := cc.service.GetCartByID(c, userID)
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	var cartRes Cart
	utils.Copy(&cartRes, cart)

	c.JSON(http.StatusOK, CartRes{
		Cart: &cartRes,
	})
}

type AddRacketCartReq struct {
	Quantity int `json:"quantity"`
	RacketID int `json:"racket_id"`
}

// AddRacket godoc
//
//	@Summary		Добавление ракетки в корзину
//	@Description	Метод для добавления ракетки в корзину
//	@Tags			User
//	@Param			req	body		AddRacketCartReq				true	"Информация для добавления ракетки в корзину"
//	@Success		200	{object}	CartRes							"Корзина благополучно изменена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка изменения корзины"
//	@Security		BearerAuth
//
//	@Router			/cart [post]
func (cc *CartController) AddRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var req AddRacketCartReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		cc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cart, err := cc.service.AddRacket(c, &dto.AddRacketCartReq{
		UserID:   userID,
		RacketID: req.RacketID,
		Quantity: req.Quantity,
	})
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	var cartRes Cart
	utils.Copy(&cartRes, cart)

	c.JSON(http.StatusOK, CartRes{
		Cart: &cartRes,
	})
}

// RemoveRacket godoc
//
//	@Summary		Удаление ракетки из корзины
//	@Description	Метод для удаления ракетки в корзину
//	@Tags			User
//	@Param			id	path		int								true	"Индентефикатор ракетки"
//	@Success		200	{object}	CartRes							"Корзина благополучно изменена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка удаления из корзины"
//	@Security		BearerAuth
//
//	@Router			/cart/rackets/{id} [delete]
func (cc *CartController) RemoveRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var req dto.RemoveRacketCartReq

	racketID, _ := strconv.Atoi(c.Param("id"))

	req.RacketID = racketID
	req.UserID = userID

	cart, err := cc.service.RemoveRacket(c, &req)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	var cartRes Cart
	utils.Copy(&cartRes, cart)

	c.JSON(http.StatusOK, CartRes{
		Cart: &cartRes,
	})
}

type UpdateRacketCartReq struct {
	Quantity int `json:"quantity"`
}

// UpdateRacket godoc
//
//	@Summary		Изменения количества ракеток в корзине
//	@Description	Метод изменения количества ракеток в корзине
//	@Tags			User
//	@Param			id	path		int								true	"Индентефикатор ракетки"
//	@Param			req	body		UpdateRacketCartReq				true	"Информация для удаления ракетки из корзины"
//	@Success		200	{object}	CartRes							"Корзина благополучно изменена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка удаления из корзины"
//	@Security		BearerAuth
//
//	@Router			/cart/rackets/{id} [put]
func (cc *CartController) UpdateRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req UpdateRacketCartReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		cc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	racketID, _ := strconv.Atoi(c.Param("id"))

	cart, err := cc.service.UpdateRacket(c, &dto.UpdateRacketCartReq{
		UserID:   userID,
		RacketID: racketID,
		Quantity: req.Quantity,
	})
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	var cartRes Cart
	utils.Copy(&cartRes, cart)

	c.JSON(http.StatusOK, CartRes{
		Cart: &cartRes,
	})
}
