package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// func prometheusHandler() gin.HandlerFunc {
// 	h := promhttp.Handler()

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	router.Use(gin.RecoveryWithWriter(os.Stdout))
	router.Use(gin.LoggerWithWriter(os.Stdout))

	// router.GET("/metrics", prometheusHandler())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	s := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
