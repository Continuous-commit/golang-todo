FROM golang:1.19-alpine AS builder

WORKDIR /go/src

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main /go/src/main.go

CMD ["air", "-c", ".air.toml"]
