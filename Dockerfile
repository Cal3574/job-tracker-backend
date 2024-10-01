# Specify the platform for AMD architecture
FROM --platform=linux/amd64 golang:1.20-alpine AS builder

# Set environment variables for Go build
ENV GO111MODULE=on
ENV CGO_ENABLED=0

# Create app directory in container
WORKDIR /app

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build the Go API (output as binary called 'app')
RUN go build -o app ./cmd/main.go

# Use a smaller image for the final deployment stage
FROM --platform=linux/amd64 alpine:3.18

# Set working directory in the container
WORKDIR /app

# Copy the built binary from the previous build stage
COPY --from=builder /app/app .


# Expose the port the Go API listens on
EXPOSE 8080

# Command to run the Go API
CMD ["./app"]
