# Use the official golang image as the base
FROM golang:latest

# Set the working directory for the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod ./
COPY go.sum ./

# Download project dependencies
RUN go mod download

# Copy all project files (excluding .git) into the working directory
COPY . .

# Build the Go binary using the specified command (assuming it's in cmd/main/main.go)
RUN go build -o bin/main cmd/main/main.go

# Expose port 8080 (or any other desired port)
EXPOSE 8080

# Set the default command to run the built binary
CMD ["./bin/main"]

# docker build -t go-app .
