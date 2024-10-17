package file

import (
	"bufio"
	"github.com/nekizz/vmo-demo-project/internal/enum"
	"github.com/nekizz/vmo-demo-project/pkg/utils/hash"
	"os"
)

type Service interface {
	CountLine(string) (int, error)
	CheckSum(string, string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CountLine(filePath string) (int, error) {
	var (
		counter int
		file    *os.File
	)

	if filePath == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(filePath)
		if err != nil {
			return 0, err
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func (s *service) CheckSum(filePath, algo string) (string, error) {
	var file *os.File

	if filePath == "-" {
		file = os.Stdin
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
	}

	hashStr, err := hash.HashFile(file, enum.Algorithm(algo))
	if err != nil {
		return "", err
	}

	return hashStr, nil
}
