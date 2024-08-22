# Use the official Golang image as a base image
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app for Linux OS
RUN CGO_ENABLED=0 GOOS=linux go build -o order_service .

# Start a new stage with a minimal base image
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/order_service .

# Ensure the binary has execute permissions
RUN chmod +x ./order_service

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./order_service"]
