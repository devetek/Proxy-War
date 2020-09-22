FROM golang:1.14.3-alpine

WORKDIR /terpusat

COPY http1/main.go /terpusat/http1/main.go
COPY types/types.go /terpusat/types/types.go
COPY go.sum /terpusat/go.sum
COPY go.mod /terpusat/go.mod

RUN go mod tidy;

CMD ["go", "run", "http1/main.go"]