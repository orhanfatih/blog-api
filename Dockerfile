FROM golang:alpine AS builder

WORKDIR /app

# Copy source code (including go.mod and go.sum)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application code (including main.go)
COPY . .

# Build the Go binary
RUN go build -o blogapi .

# Create a new image for running the application
FROM alpine:latest

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/blogapi /app

# Copy the .env file
COPY .env /app/.env

# Expose the server port
EXPOSE 3000

# Install PostgreSQL client for development (if needed)
RUN apk add --no-cache postgresql-client

# Run the application
CMD ["./blogapi"]
