package file

import (
	"bufio"
	"os"
)

type Service interface {
	CountLine(*os.File) (int, error)
	OpenFile(filePath string) (*os.File, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CountLine(file *os.File) (int, error) {
	var counter int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}

func (s *service) OpenFile(filePath string) (file *os.File, err error) {
	if filePath == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			return nil, err
		}
	}

	return
}
