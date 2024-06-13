FROM golang:1.22-alpine AS builder

RUN mkdir /app
RUN mkdir /app/mainService
RUN mkdir /app/postService
RUN mkdir /app/statisticsService

COPY mainService /app/mainService
COPY postService /app/postService
COPY statisticsService /app/statisticsService

WORKDIR /app/mainService

COPY ./mainService/go.mod .
COPY ./mainService/go.sum .

RUN go mod download && go mod verify

COPY ./mainService/ .

RUN go build -o auth-service ./cmd/main.go

# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /authService/auth-service .
# COPY --from=builder /authService/config/.env /usr/local/bin/config/.env

EXPOSE 8080
EXPOSE 5432

CMD ["./auth-service"]