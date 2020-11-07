package Engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/korableg/mini-gin/Config"
	"github.com/korableg/mini-gin/Errors"
)

var instance *gin.Engine

func init() {

	if Config.Debug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	instance = gin.New()
	instance.Use(defaultHeaders())

	instance.NoRoute(Errors.PageNotFoundHandler)
	instance.NoMethod(Errors.MethodNotAllowed)

}

func GetEngine() *gin.Engine {
	return instance
}

func Run() {
	go func() {
		instance.Run(Config.Address())
	}()
}

func defaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", fmt.Sprintf("Mini:%s", Config.Version()))
	}
}
