package file_processer

import filePkg "github.com/nekizz/vmo-demo-project/pkg/file"

type service struct {
	filer FileService
}

func NewService() *service {
	return &service{filer: filePkg.NewService()}
}

func (s *service) CountFileLine(path string) (int, error) {
	return s.filer.CountLine(path)
}

func (s *service) CheckSum(path, algo string) (string, error) {
	return s.filer.CheckSum(path, algo)
}

type FileService interface {
	CountLine(string) (int, error)
	CheckSum(string, string) (string, error)
}
