package http

import (
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(
	handler *gin.RouterGroup,
	l logging.Interface,
	authService service.IAuthService,
	userService service.IUserService) {

	authController := NewAuthController(l, authService, userService)

	auth := handler.Group("auth")

	auth.POST("/register", authController.Register)

	auth.POST("/login", authController.Login)
	auth.POST("/login/verify", authController.VerifyLogin)

	// auth.POST("/reset_password", authController.ResetPassword)
	// auth.POST("/reset_password/verify", authController.VerifyResetPassword)
}

func setRacketRoute(
	handler *gin.RouterGroup,
	l logging.Interface,
	racketService service.IRacketService,
	authService service.IAuthService,
	userService service.IUserService) {

	authController := NewAuthController(l, authService, userService)
	racketController := NewRacketController(l, racketService, userService)

	handler.GET("/rackets", racketController.ListsAllRackets)
	handler.GET("/rackets/:id", racketController.GetRacketByID)
	handler.POST("/rackets", authController.AdminIdentity, racketController.AddRacket)

	handler.PATCH("/rackets/:id", authController.AdminIdentity, racketController.UpdateRacket)
}

func setUserRoute(
	handler *gin.RouterGroup,
	l logging.Interface,
	cartService service.ICartService,
	authService service.IAuthService,
	userService service.IUserService,
	orderService service.IOrderService) {

	authController := NewAuthController(l, authService, userService)
	userController := NewUserController(l, userService, cartService, orderService)

	handler.GET("/profile", authController.UserIdentity, userController.GetMyProfile)

	handler.GET("/users", authController.AdminIdentity, userController.GetAllUsers)
	handler.GET("/users/:id", authController.AdminIdentity, userController.GetUserByID)
	handler.PUT("/users/:id", authController.AdminIdentity, userController.UpdateUser)
}

// cart
func setCartRoute(
	handler *gin.RouterGroup,
	l logging.Interface,
	cartService service.ICartService,
	authService service.IAuthService,
	userService service.IUserService,
) {

	cartController := NewCartController(l, cartService)
	authController := NewAuthController(l, authService, userService)

	handler.GET("/cart", authController.UserIdentity, cartController.GetMyCart)
	handler.POST("/cart", authController.UserIdentity, cartController.AddRacket)
	handler.PUT("/cart/rackets/:id", authController.UserIdentity, cartController.UpdateRacket)
	handler.DELETE("/cart/rackets/:id", authController.UserIdentity, cartController.RemoveRacket)
}

// order
func setOrderRoute(
	handler *gin.RouterGroup,
	l logging.Interface,
	authService service.IAuthService,
	orderService service.IOrderService,
	userService service.IUserService) {

	authController := NewAuthController(l, authService, userService)
	orderController := NewOrderController(l, orderService, userService)

	handler.GET("/orders", authController.UserIdentity, orderController.GetAllOrders)

	handler.POST("/orders", authController.UserIdentity, orderController.CreateOrder)
	handler.GET("/orders/:id", authController.AdminIdentity, orderController.GetOrderByID)
	handler.PATCH("/orders/:id", authController.AdminIdentity, orderController.UpdateOrder)
}

func SetRoutes(
	handler *gin.RouterGroup,
	l logging.Interface,
	authService service.IAuthService,
	userService service.IUserService,
	racketService service.IRacketService,
	cartService service.ICartService,
	orderService service.IOrderService,
) {

	setAuthRoute(
		handler,
		l,
		authService,
		userService,
	)

	setUserRoute(
		handler,
		l,
		cartService,
		authService,
		userService,
		orderService,
	)

	setRacketRoute(
		handler,
		l,
		racketService,
		authService,
		userService,
	)

	setCartRoute(
		handler,
		l,
		cartService,
		authService,
		userService,
	)

	setOrderRoute(
		handler,
		l,
		authService,
		orderService,
		userService,
	)
}
