# Use the official Golang image as a base image
FROM golang:1.22-alpine AS builder

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
RUN go build -o auth-service ./cmd/authService

# Use a smaller base image for the final container
FROM alpine:latest

# Set environment variables for database connection
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=todorov_want
ENV DB_PASSWORD=pass
ENV DB_NAME=auth

# Copy the built executable from the builder stage to the final image
COPY --from=builder /app/auth-service /usr/local/bin/auth-service
COPY --from=builder /app/.env /.env

# Expose the port on which the service will run
EXPOSE 8080

# Command to run the service
CMD ["auth-service"]
