package contract

import "errors"

type CountLineInput struct {
	Path string
}

func (c *CountLineInput) Validate() error {
	if c.Path == "" {
		return errors.New("empty path")
	}

	return nil
}
