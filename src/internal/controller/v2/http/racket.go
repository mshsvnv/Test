package http

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"src/internal/dto"
	"src/internal/model"
	"src/internal/service"
	"src/pkg/logging"
	"src/pkg/storage/postgres"
)

const multiFormSizeDefault = 10000000

type RacketController struct {
	l             logging.Interface
	racketService service.IRacketService
	userService   service.IUserService
}

func NewRacketController(
	l logging.Interface,
	racketService service.IRacketService,
	userService service.IUserService) *RacketController {
	return &RacketController{
		l:             l,
		racketService: racketService,
		userService:   userService,
	}
}

type RacketRes struct {
	Racket *model.Racket `json:"racket"`
}

type RacketsRes struct {
	Rackets []*model.Racket `json:"rackets"`
}

// ListAllRackets godoc
//
//	@Summary		Получение списка всех ракеток в магазине
//	@Description	Метод для получения списка всех ракеток в магазине
//	@Tags			All
//	@Param			pattern	query		string							false	"Значение для фильтрации"
//	@Param			field	query		string							false	"Поле для фильтрации и сортировки"
//	@Param			sort	query		string							false	"Направление сортировки"
//	@Success		200		{object}	RacketsRes						"Список ракеток благополучно получен"
//	@Failure		500		{object}	http.StatusInternalServerError	"Внутренняя ошибка получения списка ракеток"
//	@Router			/rackets [get]
func (r *RacketController) ListsAllRackets(c *gin.Context) {

	rackets, err := r.racketService.GetAllRackets(c, &dto.ListRacketsReq{
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
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
		return
	}

	c.JSON(http.StatusOK, RacketsRes{Rackets: rackets})
}

// GetRacketByID godoc
//
//	@Summary		Получение информации о конкретной ракетке
//	@Description	Метод для получения информации о конкретной ракетке
//	@Tags			All
//	@Param			id	path		int								true	"Идентификатор ракетки"
//	@Success		200	{object}	RacketFeedbacksRes				"Список ракеток благополучно получен"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка получения списка ракеток"
//	@Router			/rackets/{id} [get]
func (r *RacketController) GetRacketByID(c *gin.Context) {

	racketID, _ := strconv.Atoi(c.Param("id"))

	racket, err := r.racketService.GetRacketByID(c, racketID)
	if err != nil {
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RacketRes{
		Racket: racket,
	})
}

type CreateRacketReq struct {
	Brand     string  `json:"brand"`
	Weight    float32 `json:"weight"`
	Balance   float32 `json:"balance"`
	HeadSize  float32 `json:"head_size"`
	Avaliable bool    `json:"avaliable"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}

// AddRacket godoc
//
//	@Summary		Добавление ракетки в магазин
//	@Description	Метод для добавления ракетки в интернет-магазин
//	@Tags			Admin
//	@Accept			mpfd
//	@Accept			json
//	@Produce		jpeg
//	@Produce		json
//	@Param			req		formData	CreateRacketReq					true	"Информация о ракетке на добавление"
//	@Param			image	formData	file							true	"Изображение ракетки"
//	@Success		200		{object}	RacketRes						"Ракетка благополучна добавлена"
//	@Failure		400		{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		500		{object}	http.StatusInternalServerError	"Внутренняя ошибка добавления ракетки"
//	@Security		BearerAuth
//	@Router			/rackets [post]
func (r *RacketController) AddRacket(c *gin.Context) {

	err := c.Request.ParseMultipartForm(multiFormSizeDefault)
	if err != nil {
		r.l.Errorf("failed to parse multipart form: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect request"})
		return
	}

	err = c.Request.ParseForm()
	if err != nil {
		r.l.Errorf("failed to read profile data: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Incorrect profile data"})
		return
	}
	req := c.Request.Form

	weight, _ := strconv.ParseFloat(req.Get("weight"), 32)
	balance, _ := strconv.ParseFloat(req.Get("balance"), 32)
	headSize, _ := strconv.ParseFloat(req.Get("head_size"), 32)
	price, _ := strconv.ParseFloat(req.Get("price"), 32)
	quantity, _ := strconv.ParseInt(req.Get("quantity"), 10, 64)

	reqNew := &dto.CreateRacketReq{
		Brand:    req.Get("brand"),
		Weight:   float32(weight),
		Balance:  float32(balance),
		HeadSize: float32(headSize),
		Price:    float32(price),
		Quantity: int(quantity),
	}

	f, err := c.FormFile("image")
	if err != nil {
		r.l.Errorf("failed to get image from form file: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t get image from form file"})
		return
	}

	image, err := f.Open()
	if err != nil {
		r.l.Errorf("failed to open form file: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Can`t open form file"})
		return
	}
	defer image.Close()

	imageData, err := io.ReadAll(image)
	if err != nil {
		r.l.Errorf("failed to read photo data: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can`t read photo data"})
		return
	}

	racket, err := r.racketService.CreateRacket(c, reqNew)
	if err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", imageData)
	c.JSON(http.StatusOK, RacketRes{
		Racket: racket,
	})
}

type UpdateRacketReq struct {
	Quantity int `json:"quantity"`
}

// UpdateRacket godoc
//
//	@Summary		Изменение статуса ракетки
//	@Description	Метод для изменения статуса ракетки в интернет-магазин
//	@Tags			Admin
//	@Param			id	path		int								true	"Индентефикатор ракетки"
//	@Param			req	body		UpdateRacketReq					true	"Информация о ракетке на изменение"
//	@Success		200	{object}	string							"Ракетка благополучна изменена"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка добавления ракетки"
//	@Security		BearerAuth
//
//	@Router			/rackets/{id} [patch]
func (r *RacketController) UpdateRacket(c *gin.Context) {

	var req UpdateRacketReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	racketID, _ := strconv.Atoi(c.Param("id"))

	err := r.racketService.UpdateRacket(c, &dto.UpdateRacketReq{
		ID:       racketID,
		Quantity: req.Quantity,
	})
	if err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
