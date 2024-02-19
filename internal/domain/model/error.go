package model

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest = &StatusError{Code: http.StatusBadRequest, Err: errors.New("bad request")}
	ErrNotFound   = &StatusError{Code: http.StatusNotFound, Err: errors.New("not found")}
)

type StatusError struct {
	Code int
	Err  error
}

func (e *StatusError) Error() string {
	if e == nil {
		return "nil"
	}

	if e.Err != nil {
		return fmt.Sprintf("error: %s", e.Err.Error())
	}

	return "error: unexpected error"
}
