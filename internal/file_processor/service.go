package file_processor

import (
	"io"
	"os"
)

type service struct {
	filer  FileService
	hasher HasherService
}

func NewService(hasher HasherService, filer FileService) *service {
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

	return s.filer.CountLine(file)
}

func (s *service) CheckSum(path string) (string, error) {
	file, err := s.filer.OpenFile(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return s.hasher.HashFile(file)
}

type FileService interface {
	CountLine(*os.File) (int, error)
	OpenFile(filePath string) (*os.File, error)
}

type HasherService interface {
	HashFile(file io.Reader) (string, error)
}
