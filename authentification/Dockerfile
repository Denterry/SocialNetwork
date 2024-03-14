# Use the official Golang image as a base image
FROM golang:1.21-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# # Build the Go application
RUN go build -o main ./cmd

# Use a smaller base image for the final container
FROM alpine:latest

# Set environment variables from .env file
COPY .env .env

# Set environment variables for database connection
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=todorov_want
ENV DB_PASSWORD=postgres
ENV DB_NAME=auth

# Copy the built executable from the builder stage to the final image
COPY --from=builder /app/main /usr/local/bin/main

# Expose the port on which the service will run
EXPOSE 8080

# Command to run the service
CMD ["main"]
