FROM golang:1.22.3 AS builder

RUN mkdir /app
RUN mkdir /app/statisticsService
RUN mkdir /app/postService

COPY statisticsService /app/statisticsService
COPY postService /app/postService

WORKDIR /app/statisticsService

COPY ./statisticsService/go.mod .
COPY ./statisticsService/go.sum .

RUN go mod download && go mod verify

COPY ./statisticsService/ .

# COPY ./statisticsService/wait-for-it.sh /usr/local/bin/wait-for-it.sh
# RUN chmod +x /usr/local/bin/wait-for-it.sh

RUN go build -o statistics-service ./cmd/main.go

EXPOSE 50053
EXPOSE 8082

CMD ["./statistics-service"]
