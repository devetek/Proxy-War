FROM golang:1.14.3-alpine

WORKDIR /terpusat

COPY http2/main.go /terpusat/http2/main.go
COPY go.sum /terpusat/go.sum
COPY go.mod /terpusat/go.mod

RUN go mod tidy;

CMD ["go", "run", "http2/main.go"]