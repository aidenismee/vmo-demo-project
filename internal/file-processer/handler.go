package file_processer

import "github.com/nekizz/vmo-demo-project/internal/contract"

type Service interface {
	CountFileLine(string) (int, error)
	CheckSum(string, string) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CountFileLine(input *contract.CountLineInput) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	resp, err := h.service.CountFileLine(input.Path)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (h *Handler) CheckSum(input *contract.CheckSumInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", err
	}

	resp, err := h.service.CheckSum(input.Path, input.Algo)
	if err != nil {
		return "", err
	}

	return resp, nil
}
