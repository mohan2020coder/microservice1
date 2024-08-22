# Use the official Golang image as a base image for building
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o user_service .

# Use a minimal base image for the final container
FROM alpine:latest  

# Set the current working directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/user_service .

# Ensure the binary has execute permissions
RUN chmod +x ./user_service

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./user_service"]
