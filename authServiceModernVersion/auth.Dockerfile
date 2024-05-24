FROM golang:1.22-alpine AS builder

RUN mkdir /app
RUN mkdir /app/authServiceModernVersion
RUN mkdir /app/postService

COPY authServiceModernVersion /app/authServiceModernVersion
COPY postService /app/postService

WORKDIR /app/authServiceModernVersion

COPY ./authServiceModernVersion/go.mod .
COPY ./authServiceModernVersion/go.sum .

RUN go mod download && go mod verify

COPY ./authServiceModernVersion/ .

RUN go build -o auth-service ./cmd/main.go

# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /authService/auth-service .
# COPY --from=builder /authService/config/.env /usr/local/bin/config/.env

EXPOSE 8080
EXPOSE 5432

CMD ["./auth-service"]