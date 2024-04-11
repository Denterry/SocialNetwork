FROM golang:1.22-alpine AS builder


# Set the current working directory inside the container
WORKDIR /app

# Copy the entire directory contents into the container
COPY . .

# Build the Go application
RUN go build -o post-service ./cmd

# Use a smaller base image for the final container
FROM alpine:latest

# Set environment variables for database connection
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=todorov_want
ENV DB_PASSWORD=postgres
ENV DB_NAME=auth

# Copy the built executable from the builder stage to the final image
COPY --from=builder /app/post-service /usr/local/bin/post-service

# Expose the port on which the service will run
EXPOSE 50051

# Command to run the service
CMD ["post-service"]
