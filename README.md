# VMO Demo Project

Hi 👋, This is an demo project for software engineering team members


## About
Go programming challenge for `futil`.
This is an CLI `application` for utility tool for file processing. With easy-to-use features, it allows users to perform common tasks without much effort.


## Prerequisite:
- Linux based OS is preferred
- Docker (version 20.10.21 is preferred) with compose plugin or *docker-compose*
- Go (version 1.20+)

## How to Build and Run the Project
1. Clone the Repository:
2. Build the project: make build
3. Run the application: ./bin/vmo-demo-project help

## Technical choice:
- Cobra: [Golang Cobra](https://github.com/spf13/cobra/)
- Go Mock: [Go Mock](https://github.com/uber-go/mock/)
- Ginkgo: [Ginkgo](https://onsi.github.io/ginkgo/)


## Architecture
### Project Architecture
- Domain based approach with a loosely applied layered architecture is utilized for this project.
- `cmd` directory is the default directory created by `Cobra` for a command based project. This folder includes command implementations and serves as the entry point for different application commands
- The `internal` directory contains packages that encapsulate business logic or application-specific functionality. This could also include utility functions used across the project.
- Within the `internal` directory, each entity has its own directory containing its specific logic and handler.
- `contract` directory define the request and response of handler.
- `enum`, `utils` directory define the shared functions, errors, constant through the projects. 
- The library is considered external is located on `pkg` directory.


### Project Structure
```
├── .github
│  └── workflows
│  │  ├── build_code.yaml 
│  │  ├── build_docker.yaml
│  │  └── CI.yaml
├── cmd
│  ├── checksum.go
│  ├── linecount.go
│  ├── root.go
│  └── version.go
├── deployments
│  ├── docker-compose.yaml
│  └── Dockerfile
├── internal
│  ├── contract
│  │  ├── checksum.go
│  │  └── linecount.go
│  ├── enum
│  │  └── algorithm.go
│  ├── file_processor
│  │   ├── handler.go
│  │   ├── handler_test.go
│  │   ├── mock.go
│  │   ├── service.go
│  │   └── service_test.go
├── pkg
│  ├── filer
│  │  ├── filer.go
│  │  └── mock.go
│  ├── hasher
│  │  ├── hasher.go
│  │  └── mock.go
│  ├── utils
│  │   ├── app_error
│  │   │  └── app_error.go
│  │   └── array
│  │   │  └── array.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── README.md
```

### Assessments
* **Architecture**: The logic is divided into 3 separate layer: handler, and service layer. Each layer has one single purpose (handler to validate and handle requests from clients, service for data handling and conversion) and independent of one and another.
* **Clarity**: The project purpose, architecture, structure and assessment criteria are explained clearly in the `README.md` file.
* **Code quality**: The source code is well organized with few nested functions and directory along with clear structure, logic block and easy to navigate structure.
* **Usability**: the project is suitable for multiple clients, even with physical constraint as the source code is relatively lightweight and required no extensive CPU or memory to execute.
* **Security**: No obvious security flaw is witness during the development and testing process.
* **Testing**: basic unit test for some important functions is included to demonstrate testing coverage of the project.
* **Technical choices**: the decision of using Cobra and Ginkgo are both lightweight and suitable for this kind of project given the time-constraint and project complexity.
* **Production-readiness**: error handling and logging are included. Monitoring and tracing are not implemented due to the time-constraint.
* **Scalability**: The scalability can be achieved relative easily thanks to the separation of important layers.


### Future direction
- Support for calculating checksums for multiple files at once.
- Output file writing for results.

