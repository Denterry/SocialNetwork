FROM golang:1.22-alpine AS builder

RUN mkdir /app
RUN mkdir /app/postService

COPY postService /app/postService

WORKDIR /app/postService

COPY ./postService/go.mod .
COPY ./postService/go.sum .

RUN go mod download && go mod verify

COPY ./postService/ .

RUN go build -o post-service ./cmd/main.go

EXPOSE 8081

CMD ["./post-service"]
