# Use an official Golang image as the base
FROM golang:1.22.3

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the application source code
COPY app/ .

# Build the Go app
RUN go build -o main .

# Expose the port on which the app runs
EXPOSE 8080

# Start the Go app
CMD ["./main"]

