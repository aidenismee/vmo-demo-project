package contract

import (
	"errors"
)

type CheckSumInput struct {
	Path string
}

func (c *CheckSumInput) Validate() error {
	if c.Path == "" {
		return errors.New("empty path")
	}

	return nil
}
