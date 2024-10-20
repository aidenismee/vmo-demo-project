.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
DOCKER_USERNAME ?= n3kizz
APPLICATION_NAME ?= vmo-demo-project

export TAG

install:
	go get .

run-test:
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/futil .

gen-mock:
	mockgen -destination pkg/filer/mock.go -package=filer github.com/nekizz/vmo-demo-project/pkg/filer Filer
	mockgen -destination pkg/hasher/mock.go -package=hasher github.com/nekizz/vmo-demo-project/pkg/hasher Hasher
	mockgen -destination internal/file_processor/mock.go -package=file_processor github.com/nekizz/vmo-demo-project/internal/file_processor Service

docker-build:
	docker build --file ./deployments/Dockerfile -t $(DOCKER_USERNAME)/$(APPLICATION_NAME) .

docker-push:
	docker push $(DOCKER_USERNAME)/$(APPLICATION_NAME)

