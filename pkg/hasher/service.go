package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/nekizz/vmo-demo-project/internal/enum"
	"hash"
	"io"
)

var mappingAlgoHasher = map[enum.Algorithm]hash.Hash{
	enum.Md5:    md5.New(),
	enum.Sha1:   sha1.New(),
	enum.Sha256: sha256.New(),
}

type Service interface {
	HashFile(file io.Reader) (string, error)
}

type service struct {
	hasher hash.Hash
}

func NewService(algo string) Service {
	return &service{
		hasher: mappingAlgoHasher[enum.Algorithm(algo)],
	}
}

func (s *service) HashFile(file io.Reader) (string, error) {
	if _, err := io.Copy(s.hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(s.hasher.Sum(nil)), nil
}
