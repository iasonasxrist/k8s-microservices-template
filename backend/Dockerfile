# Use the official Golang image to build the Go application
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files if available
COPY go.mod ./
#COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download || true

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
