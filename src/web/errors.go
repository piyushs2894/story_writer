package web

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("Not Found")
	ErrUnexpected   = errors.New("Unexpected error has occured")
	ErrConnectivity = errors.New("Please try again after sometime.")
)
