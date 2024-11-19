package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docsv2 "src/docs/v2"
	routesv2 "src/internal/controller/v2/http"
	"src/internal/service"
	"src/pkg/logging"
)

type Controller struct {
	handler      *gin.Engine
	routerGroups map[string]*gin.RouterGroup
}

func NewRouter(handler *gin.Engine) *Controller {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	docsv2.SwaggerInfov2.BasePath = "/api/v2"

	v2 := handler.Group("/api/v2")
	{
		v2.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
		v2.GET("/healthcheck", healthCheck)
	}

	return &Controller{
		handler: handler,
		routerGroups: map[string]*gin.RouterGroup{
			"v2": v2,
		},
	}
}

// healthCheck godoc
//
//	@Summary		Проверка здоровья
//	@Description	Проверка на жизнеспособность
//	@Tags			system
//	@Success		200	{string} string "Сервис жив"
//	@Failure		404	"Сервис мертв"
//	@Router			/healthcheck [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, time.Now().String())
}

func (c *Controller) SetV2Routes(
	l logging.Interface,
	authService service.IAuthService,
	userService service.IUserService,
	racketService service.IRacketService,
	cartService service.ICartService,
	orderService service.IOrderService,
) {
	routesv2.SetRoutes(
		c.routerGroups["v2"],
		l,
		authService,
		userService,
		racketService,
		cartService,
		orderService,
	)
}
