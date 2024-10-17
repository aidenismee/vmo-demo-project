package hash

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

func HashFile(file io.Reader, algo enum.Algorithm) (string, error) {
	hasher := mappingAlgoHasher[algo]
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
