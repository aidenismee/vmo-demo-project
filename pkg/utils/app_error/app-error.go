package app_error

import (
	"errors"
)

var (
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrEmptyPath            = errors.New("empty file path")
	ErrEmptyAlgorithm       = errors.New("empty algorithm")
	ErrMoreThanOneAlgorithm = errors.New("more than one algorithm")
	ErrInvalidAlgorithm     = errors.New("invalid algorithm")
	ErrInvalidInput         = errors.New("invalid input")
	ErrInternal             = errors.New("internal error")
)
