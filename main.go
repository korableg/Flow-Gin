package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/korableg/mini-gin/Config"
	"github.com/korableg/mini-gin/Errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func defaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", fmt.Sprintf("Mini:%s", Config.Version()))
	}
}

func shutdown() {
	log.Println("Shutting down mini...")

	// TODO Save messages, nodes, routers to disk
}

func main() {

	log.Println("Starting mini...")

	//gin.SetMode(gin.ReleaseMode)

	defer shutdown()

	engine := gin.New()
	engine.Use(defaultHeaders())

	engine.NoRoute(Errors.PageNotFoundHandler)
	engine.NoMethod(Errors.MethodNotAllowed)

	go func() {
		engine.Run(Config.Address())
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

}
