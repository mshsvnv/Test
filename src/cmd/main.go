package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"src/config"
	"src/internal/controller"
	mypostgres "src/internal/repository/postgres"
	"src/internal/service"
	"src/pkg/logging"
	httpserver "src/pkg/server/http"
	"src/pkg/storage/postgres"
)

func main() {

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
// package main

// import (
// 	"github.com/pquerna/otp"
// 	"github.com/pquerna/otp/totp"

// 	"bufio"
// 	"bytes"
// 	"encoding/base32"
// 	"fmt"
// 	"image/png"
// 	"os"
// 	"time"
// )

// func display(key *otp.Key, data []byte) {
// 	fmt.Printf("Issuer:       %s\n", key.Issuer())
// 	fmt.Printf("Account Name: %s\n", key.AccountName())
// 	fmt.Printf("Secret:       %s\n", key.Secret())
// 	fmt.Println("Writing PNG to qr-code.png....")
// 	os.WriteFile("qr-code.png", data, 0644)
// 	fmt.Println("")
// 	fmt.Println("Please add your TOTP to your OTP Application now!")
// 	fmt.Println("")
// }

// func promptForPasscode() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter Passcode: ")
// 	text, _ := reader.ReadString('\n')
// 	return text
// }

// // Demo function, not used in main
// // Generates Passcode using a UTF-8 (not base32) secret and custom parameters
// func GeneratePassCode(utf8string string) string {
// 	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
// 	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
// 		Period:    30,
// 		Skew:      1,
// 		Digits:    otp.DigitsSix,
// 		Algorithm: otp.AlgorithmSHA512,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return passcode
// }

// func main() {
// 	key, err := totp.Generate(totp.GenerateOpts{
// 		Issuer:      "Example.com",
// 		AccountName: "alice@example.com",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Convert TOTP key into a PNG
// 	var buf bytes.Buffer
// 	img, err := key.Image(200, 200)
// 	if err != nil {
// 		panic(err)
// 	}
// 	png.Encode(&buf, img)

// 	// display the QR code to the user.
// 	display(key, buf.Bytes())

// 	// Now Validate that the user's successfully added the passcode.
// 	fmt.Println("Validating TOTP...")
// 	passcode := promptForPasscode()
// 	valid := totp.Validate(passcode, key.Secret())
// 	if valid {
// 		println("Valid passcode!")
// 		os.Exit(0)
// 	} else {
// 		println("Invalid passcode!")
// 		os.Exit(1)
// 	}
// }
