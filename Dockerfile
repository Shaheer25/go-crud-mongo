FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./

# RUN apk update && apk add --no-cache git

# RUN go install github.com/cosmtrek/air@latest

COPY . .

EXPOSE 8080

ARG ENV_FILE
ENV ENV_FILE ${ENV_FILE:-.env}
COPY $ENV_FILE .env

CMD go run main.go
