package contract

import (
	appError "github.com/nekizz/vmo-demo-project/pkg/utils/app_error"
)

type CountLineInput struct {
	Path string
}

func (c *CountLineInput) Validate() error {
	if c.Path == "" {
		return appError.ErrEmptyPath
	}

	return nil
}
