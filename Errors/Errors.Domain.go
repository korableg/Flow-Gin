package Errors

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ERR_HUB_NAME_ISEMPTY      = errors.New("hub name is empty")
	ERR_HUB_IS_ALREADY_EXISTS = errors.New("hub is already exists")
	ERR_NODE_NAME_ISEMPTY     = errors.New("Node name is empty")
	ERR_NODE_NOT_FOUND        = errors.New("Node not found")
	ERR_PAGE_NOT_FOUND        = errors.New("Page not found")
	ERR_METHOD_NOT_ALLOWED    = errors.New("Method not allowed")
)

type Error struct {
	Error string
}

func NewError(err error) *Error {
	return &Error{Error: err.Error()}
}

func (e *Error) ToGin() *gin.H {
	return &gin.H{"error": e.Error}
}
