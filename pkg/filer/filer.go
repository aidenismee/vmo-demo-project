package filer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/nekizz/vmo-demo-project/internal/enum"
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
		return os.Stdin, nil
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("no such file '%s'", filePath)
		}

		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("expected file got directory '%s'", filePath)
	}

	detectedMIME, err := mimetype.DetectFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot detech file '%s'", filePath)
	}

	if detectedMIME.Is(enum.ExecutableApplication) {
		return nil, fmt.Errorf("cannot do linecount for binary file '%s'", fileInfo.Name())
	}

	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
