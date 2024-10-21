package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/model"
	"src/internal/service"
	"src/pkg/logging"
	"src/pkg/storage/postgres"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	l            logging.Interface
	orderService service.IOrderService
	userService  service.IUserService
}

func NewOrderController(
	l logging.Interface,
	orderService service.IOrderService,
	userService service.IUserService,
) *OrderController {
	return &OrderController{
		l:            l,
		orderService: orderService,
		userService:  userService,
	}
}

type OrderRes struct {
	Order *model.Order `json:"order"`
}

type OrdersRes struct {
	Orders []*model.Order `json:"orders"`
}

type PlaceOrderReq struct {
	DeliveryDate  time.Time `json:"delivery_date" format:"2006-01-02T15:07:00Z"`
	Address       string    `json:"address"`
	RecepientName string    `json:"recepient_name"`
}

type UpdateOrderReq struct {
	Status model.OrderStatus `json:"status"`
}

// UpdateOrder godoc
//
//	@Summary		Изменение статуса заказа
//	@Description	Метод для изменения статуса заказа
//	@Tags			Admin
//	@Param			id	path		int								true	"Идентефикатор заказа"
//	@Param			req	body		UpdateOrderReq					true	"Информация о новом статусе"
//	@Success		200	{object}	OrderRes						"Информация об изменении статуса заказа"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при изменении статуса заказа"
//	@Security		BearerAuth
//	@Router			/orders/{id} [patch]
func (o *OrderController) UpdateOrder(c *gin.Context) {

	var req UpdateOrderReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := o.orderService.UpdateOrderStatus(c, &dto.UpdateOrderReq{
		OrderID: orderID,
		Status:  req.Status,
	})
	if err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, OrderRes{
		Order: order,
	})
}

// GetOrderByID godoc
//
//	@Summary		Получение заказа по идентефикатору
//	@Description	Метод для получения заказа по идентефикатору
//	@Tags			Admin
//	@Param			id	query		string							true	"Идентефикатор заказа"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при изменении статуса заказа"
//	@Security		BearerAuth
//	@Router			/orders/{id} [get]
func (o *OrderController) GetOrderByID(c *gin.Context) {

	orderID, _ := strconv.Atoi(c.Param("id"))

	order, err := o.orderService.GetOrderByID(c, orderID)
	if err != nil {
		o.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, OrderRes{
		Order: order,
	})
}

// GetAllOrders godoc
//
//	@Summary		Получение заказов в инетрнет-магазине
//	@Description	Метод для получения заказов в инетрнет-магазине
//	@Tags			Admin & User
//	@Param			pattern	query		string							false	"Значение для фильтрации"
//	@Param			field	query		string							false	"Поле для фильтрации и сортировки"
//	@Param			sort	query		string							false	"Направление сортировки"
//	@Success		200		{object}	OrdersRes						"Список заказов"
//	@Failure		401		{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500		{object}	http.StatusInternalServerError	"Внутренняя ошибка при изменении статуса заказа"
//	@Security		BearerAuth
//	@Router			/orders [get]
func (o *OrderController) GetAllOrders(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := o.userService.GetUserByID(c, userID)
	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if user.Role == model.UserRoleAdmin {
		orders, err := o.orderService.GetAllOrders(c, &dto.ListOrdersReq{
			Pagination: &postgres.Pagination{
				Filter: postgres.FilterOptions{
					Pattern: c.Query("pattern"),
					Column:  c.Query("field"),
				},
				Sort: postgres.SortOptions{
					Direction: postgres.SortDirectionFromString(c.Query("sort")),
					Columns:   []string{c.Query("field")},
				},
			},
		})

		if err != nil {
			o.l.Errorf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, OrdersRes{
			Orders: orders,
		})

		return

	} else if user.Role == model.UserRoleCustomer {
		orders, err := o.orderService.GetMyOrders(c, userID)

		if err != nil {
			o.l.Errorf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, OrdersRes{
			Orders: orders,
		})

		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
}

// CreateOrder godoc
//
//	@Summary		Создание заказа
//	@Description	Метод для создания заказа
//	@Tags			User
//
//	@Param			req	body		PlaceOrderReq					true	"Информация о заказе"
//
//	@Success		200	{object}	string							"Информация об успешности создания заказа"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при размещении заказа"
//	@Security		BearerAuth
//	@Router			/orders [post]
func (o *OrderController) CreateOrder(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var req PlaceOrderReq
	if err := c.ShouldBindBodyWithJSON(&req); c.Request.Body == nil || err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := o.orderService.CreateOrder(c, &dto.PlaceOrderReq{
		UserID:        userID,
		DeliveryDate:  req.DeliveryDate,
		Address:       req.Address,
		RecepientName: req.RecepientName,
	}); err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error 2"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
