package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
	"src/pkg/utils"
)

type FeedbackController struct {
	l       logging.Interface
	service service.IFeedbackService
}

func NewFeedbackController(l logging.Interface, service service.IFeedbackService) *FeedbackController {
	return &FeedbackController{
		l:       l,
		service: service,
	}
}

type Feedback struct {
	RacketID int       `json:"racket_id"`
	Feedback string    `json:"feedback"`
	Date     time.Time `json:"date"`
	Rating   int       `json:"rating"`
}

type FeedbackRes struct {
	Feedback *Feedback `json:"feedback"`
}

type FeedbacksRes struct {
	Feedbacks []*Feedback `json:"feedbacks"`
}

// GetFeedbacksByRacketID godoc
//
//	@Summary		Получение отзывов на ракетку
//	@Description	Метод для получение отзывов на ракетку
//	@Tags			All
//
//	@Param			id	path		int								true	"Идентефикатор ракетки"
//
//	@Success		200	{object}	FeedbacksRes					"Список ракеток благополучно получен"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при получении отзывов"
//	@Router			/feedbacks/rackets/{id} [get]
func (fc *FeedbackController) GetFeedbacksByRacketID(c *gin.Context) {

	racketID, _ := strconv.Atoi(c.Param("id"))

	feedbacks, err := fc.service.GetFeedbacksByRacketID(c, racketID)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var feedbacksRes []*Feedback
	for _, feedback := range feedbacks {

		var feedbackRes Feedback
		utils.Copy(&feedbackRes, feedback)

		feedbacksRes = append(feedbacksRes, &feedbackRes)
	}

	c.JSON(http.StatusOK, FeedbacksRes{
		Feedbacks: feedbacksRes,
	})
}

// GetFeedbacksByUserID godoc
//
//	@Summary		Получение отзывов пользователя
//	@Description	Метод для получения отзывов пользователя
//	@Tags			User
//	@Success		200	{object}	FeedbacksRes					"Список ракеток благополучно получен"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при получении отзывов"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Security		BearerAuth
//	@Router			/feedbacks [get]
func (fc *FeedbackController) GetFeedbacksByUserID(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	feedbacks, err := fc.service.GetFeedbacksByUserID(c, userID)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var feedbacksRes []*Feedback
	for _, feedback := range feedbacks {

		var feedbackRes Feedback
		utils.Copy(&feedbackRes, feedback)

		feedbacksRes = append(feedbacksRes, &feedbackRes)
	}

	c.JSON(http.StatusOK, FeedbacksRes{
		Feedbacks: feedbacksRes,
	})
}

type CreateFeedbackReq struct {
	Feedback string `json:"feedback"`
	Rating   int    `json:"rating"`
	RacketID int    `json:"racket_id"`
}

// CreateFeedback godoc
//
//	@Summary		Создание отзыва
//	@Description	Метод для создания отзыва
//	@Tags			User
//
//	@Param			req	body		CreateFeedbackReq				true	"Информация для создания отзыва"
//
//	@Success		200	{object}	FeedbackRes						"Созданный отзыв"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при создании отзыва"
//	@Security		BearerAuth
//	@Router			/feedbacks [post]
func (fc *FeedbackController) CreateFeedback(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var req CreateFeedbackReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		fc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	feedback, err := fc.service.CreateFeedback(c, &dto.CreateFeedbackReq{
		RacketID: req.RacketID,
		UserID:   userID,
		Feedback: req.Feedback,
		Rating:   req.Rating,
	})

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var feedbackRes Feedback
	utils.Copy(&feedbackRes, feedback)

	c.JSON(http.StatusOK, FeedbackRes{
		Feedback: &feedbackRes,
	})
}

// DeleteFeedback godoc
//
//	@Summary		Удаление отзыва
//	@Description	Метод для удаления отзыва
//	@Tags			User
//
//	@Param			id	path		int								true	"Идентефикатор ракетки"
//
//	@Success		200	{string}	string							"Созданный отзыв"
//	@Failure		400	{object}	http.StatusBadRequest			"Некорректное тело запроса"
//	@Failure		401	{object}	http.StatusUnauthorized			"Пользователь не авторизован"
//	@Failure		500	{object}	http.StatusInternalServerError	"Внутренняя ошибка при удалении отзыва"
//	@Security		BearerAuth
//	@Router			/feedbacks/{id} [delete]
func (fc *FeedbackController) DeleteFeedback(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	racketID, _ := strconv.Atoi(c.Param("id"))

	var req dto.DeleteFeedbackReq

	req.UserID = userID
	req.RacketID = racketID

	err = fc.service.DeleteFeedback(c, &req)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
