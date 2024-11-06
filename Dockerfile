# Use an official Go image as the base image
FROM golang:1.23.2-alpine3.20
LABEL authors="blank"

# Install GCC and other dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set the entrypoint
ENTRYPOINT ["./app"]
ENTRYPOINT ["top", "-b"]