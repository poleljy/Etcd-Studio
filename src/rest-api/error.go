package rest_api

import (
	"encoding/json"
	"net/http"
)

// Error implements the error interface.
type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"error"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func NewError(detail string, code int32) error {
	return &Error{
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}
