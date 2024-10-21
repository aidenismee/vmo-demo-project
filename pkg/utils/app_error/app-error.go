package app_error

import (
	"errors"
)

var (
	ErrEmptyPath            = errors.New("empty file path")
	ErrEmptyAlgorithm       = errors.New("empty algorithm")
	ErrMoreThanOneAlgorithm = errors.New("more than one algorithm")
	ErrInvalidAlgorithm     = errors.New("invalid algorithm")
)
