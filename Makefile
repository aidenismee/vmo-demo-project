.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/vmo-demo-project .

gen-mock:
	mockgen -destination pkg/file/mock.go -package=file github.com/nekizz/vmo-demo-project/pkg/file Service
	mockgen -destination pkg/hasher/mock.go -package=hasher github.com/nekizz/vmo-demo-project/pkg/hasher Service
	mockgen -destination internal/file_processor/mock.go -package=file_processor github.com/nekizz/vmo-demo-project/internal/file_processor Service

run-test:
	go test ./...