package filer

import (
	"bufio"
	"os"
)

type Filer interface {
	CountLine(*os.File) (int, error)
	OpenFile(filePath string) (*os.File, error)
}

type filer struct{}

func NewFiler() Filer {
	return &filer{}
}

func (s *filer) CountLine(file *os.File) (int, error) {
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

func (s *filer) OpenFile(filePath string) (file *os.File, err error) {
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
