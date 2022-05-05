package constants

import "errors"

var (
	ErrServerError = errors.New("server error")
	ErrNotFound    = errors.New("not found")
	ErrBadRequest  = errors.New("bad request")
)

const DeleteSuccess = "Deleted Successfully"
