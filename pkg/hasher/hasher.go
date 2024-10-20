package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"

	"github.com/nekizz/vmo-demo-project/internal/enum"
)

var mappingAlgoHasher = map[enum.Algorithm]hash.Hash{
	enum.Md5:    md5.New(),
	enum.Sha1:   sha1.New(),
	enum.Sha256: sha256.New(),
}

type Hasher interface {
	HashFile(file io.Reader) (string, error)
}

type hasher struct {
	hash hash.Hash
}

func NewHasher(algo string) Hasher {
	return &hasher{
		hash: mappingAlgoHasher[enum.Algorithm(algo)],
	}
}

func (s *hasher) HashFile(file io.Reader) (string, error) {
	if _, err := io.Copy(s.hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(s.hash.Sum(nil)), nil
}
