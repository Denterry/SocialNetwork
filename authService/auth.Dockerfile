# Use the official Golang image as a base image
FROM golang:1.22-alpine AS builder

RUN mkdir /authService

WORKDIR /authService

COPY go.mod /authService
COPY go.sum /authService

RUN go mod download

COPY . /authService
COPY ./.env /authService/.env
COPY ./config/database.env /authService/config/database.env

RUN go build -o auth-service ./cmd/authService

# COPY --from=builder /authService/auth-service /usr/local/bin/auth-service
# COPY --from=builder /authService/config/database.env /usr/local/bin/config/database.env

EXPOSE 8080
EXPOSE 5432

CMD ["./auth-service"]
