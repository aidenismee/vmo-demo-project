package contract

import (
	"errors"

	"github.com/nekizz/vmo-demo-project/internal/enum"
	arrayUtils "github.com/nekizz/vmo-demo-project/pkg/utils/array"
)

type CheckSumInput struct {
	Path string
	Algo string
}

func (c *CheckSumInput) Validate() error {
	if c.Path == "" {
		return errors.New("empty path")
	}

	if !arrayUtils.Contains[enum.Algorithm]([]enum.Algorithm{enum.Sha256, enum.Sha1, enum.Md5}, enum.Algorithm(c.Algo)) {
		return errors.New("invalid algorithm")
	}

	return nil
}
