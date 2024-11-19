package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/exp/rand"

	"src/config"
	"src/internal/controller"
	mypostgres "src/internal/repository/postgres"
	"src/internal/service"
	"src/pkg/logging"
	httpserver "src/pkg/server/http"
	"src/pkg/storage/postgres"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		cfg.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	l := logging.New(cfg.Logger.Level, loggerFile)

	db, _ := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.Database,
	))

	userRepo := mypostgres.NewUserRepository(db)
	racketRepo := mypostgres.NewRacketRepository(db)
	cartRepo := mypostgres.NewCartRepository(db)
	orderRepo := mypostgres.NewOrderRepository(db)

	userService := service.NewUserService(l, userRepo)
	racketService := service.NewRacketService(l, racketRepo)
	cartService := service.NewCartService(l, cartRepo, racketRepo)
	authService := service.NewAuthService(l, userRepo, cfg.Auth.SigningKey, cfg.Auth.AccessTokenTTL)
	orderService := service.NewOrderService(l, orderRepo, cartRepo, racketRepo)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// recepientName := os.Getenv("RECEPIENT_EMAIL_ADDRESS")
	// user := &model.User{
	// 	Name:     "Admin",
	// 	Surname:  "Admin",
	// 	Email:    recepientName,
	// 	Password: "123456",
	// 	Role:     "Admin",
	// }

	// userRepo.Create(context.TODO(), user)

	// Create controller
	handler := gin.New()
	con := controller.NewRouter(handler)

	// Set routes
	con.SetV2Routes(l, authService, userService, racketService, cartService, orderService)

	// Create router
	router := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	// Starting server
	err = router.Start()
	if err != nil {
		log.Fatal(err)
	}
}
