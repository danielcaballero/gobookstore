# Build stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Ensure dependencies are up-to-date
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /home/appuser

# Install certificates if your app makes HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D appuser

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Change ownership to the non-root user
RUN chown appuser:appuser ./main

# Switch to the non-root user
USER appuser

# Expose the port that your application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
