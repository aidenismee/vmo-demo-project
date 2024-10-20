package file_processor

import (
	"io"
	"os"
)

type service struct {
	filer  Filer
	hasher Hasher
}

func NewService(hasher Hasher, filer Filer) *service {
	return &service{
		hasher: hasher,
		filer:  filer,
	}
}

func (s *service) CountFileLine(path string) (int, error) {
	file, err := s.filer.OpenFile(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	resp, err := s.filer.CountLine(file)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (s *service) CheckSum(path string) (string, error) {
	file, err := s.filer.OpenFile(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	resp, err := s.hasher.HashFile(file)
	if err != nil {
		return "", err
	}

	return resp, nil
}

type Filer interface {
	CountLine(*os.File) (int, error)
	OpenFile(filePath string) (*os.File, error)
}

type Hasher interface {
	HashFile(file io.Reader) (string, error)
}
