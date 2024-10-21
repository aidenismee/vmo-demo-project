ARG GO_IMAGE=golang:1.23.2-bookworm
ARG DEBIAN_IMAGE=gcr.io/distroless/base-debian12:latest

FROM ${GO_IMAGE} AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /futil .

FROM ${DEBIAN_IMAGE} AS app

WORKDIR /

COPY --from=builder ["/futil", "/"]

ENTRYPOINT ["./futil"]
CMD ["version"]