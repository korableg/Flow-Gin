package Errors

import (
	"github.com/gin-gonic/gin"
)

func PageNotFoundHandler(c *gin.Context) {
	c.JSON(404, NewError(ERR_PAGE_NOT_FOUND).ToGin())
}

func MethodNotAllowed(c *gin.Context) {
	c.JSON(404, NewError(ERR_METHOD_NOT_ALLOWED).ToGin())
}
