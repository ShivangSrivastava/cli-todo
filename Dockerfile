FROM golang:1.24-alphine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .
