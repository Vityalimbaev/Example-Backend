package server

import (
	"context"
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(router *gin.Engine) error {

	serverConfig := config.GetServerConfig()

	server := &http.Server{
		Addr:    serverConfig.Bind + ":" + serverConfig.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatalf("Failed to listen and serve: " + err.Error())
		}
	}()

	quet := make(chan os.Signal, 1)
	signal.Notify(quet, os.Interrupt, os.Interrupt)

	<-quet

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return server.Shutdown(ctx)
}

func GetRouter() *gin.Engine {
	var router *gin.Engine

	router = gin.Default()
	//router.Use(enableCors())

	return router
}

func enableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}
