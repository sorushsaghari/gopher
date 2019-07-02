package models

import "errors"

var (
	ErrNotFound = errors.New("models: resource not found")
	ErrBadRequest = errors.New("models:bad request")
)
