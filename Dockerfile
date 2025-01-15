# Use the official Golang image as a builder
FROM golang:1.23.3 AS builder

# Set working directory
WORKDIR /app

# Copy source files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN go build -o /artistWebhook

# Use a minimal base image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Create data directory for Unraid config storage
RUN mkdir -p /data

# Copy the compiled binary
COPY --from=builder /artistWebhook .

# Expose application port
EXPOSE 8080

# Start the application
CMD ["./artistWebhook"]
