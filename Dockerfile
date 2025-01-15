# Use the official Golang image as a builder
FROM golang:1.23 AS builder
ARG CGO_ENABLED=0

# Set working directory
WORKDIR /app

# Copy source files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN go build -o webhook

# Use a minimal base image
FROM alpine:latest

# Copy the compiled binary
COPY --from=builder /app/webhook .

# Expose application port
EXPOSE 8080

# Start the application
CMD ["/webhook"]
