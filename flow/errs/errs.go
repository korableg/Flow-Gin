package errs

import (
	"encoding/json"
	"errors"
)

var (
	ERR_HUB_NAME_ISEMPTY              = errors.New("hub name is empty")
	ERR_HUB_NAME_OVER100              = errors.New("hub name over 100 symbols")
	ERR_HUB_IS_ALREADY_EXISTS         = errors.New("hub is already exists")
	ERR_HUB_NAME_NOT_MATCHED_PATTERN  = errors.New("the hub name should be contain only letters, digits, ., -, _")
	ERR_HUB_NOT_FOUND                 = errors.New("hub not found")
	ERR_NODE_NAME_ISEMPTY             = errors.New("node name is empty")
	ERR_NODE_NAME_OVER100             = errors.New("node name over 100 symbols")
	ERR_NODE_NOT_FOUND                = errors.New("node not found")
	ERR_NODE_IS_ALREADY_EXISTS        = errors.New("node is already exists")
	ERR_NODE_NAME_NOT_MATCHED_PATTERN = errors.New("the node name should be contain only letters, digits, ., -, _")
	ERR_PAGE_NOT_FOUND                = errors.New("page not found")
)

type Error struct {
	error string
}

func New(err error) *Error {
	return &Error{error: err.Error()}
}

func (e *Error) Error() string {
	return e.error
}

func (e *Error) MarshalJSON() ([]byte, error) {

	errormap := make(map[string]interface{})
	errormap["error"] = e.error

	return json.Marshal(errormap)

}
