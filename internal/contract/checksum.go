package contract

import (
	appError "github.com/nekizz/vmo-demo-project/pkg/utils/app_error"
)

type CheckSumInput struct {
	Path string
}

func (c *CheckSumInput) Validate() error {
	if c.Path == "" {
		return appError.ErrEmptyPath
	}

	return nil
}
