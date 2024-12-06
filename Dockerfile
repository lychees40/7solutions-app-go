# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .
RUN CGO_ENABLED=0 go build -o main .
# Stage 2: Create a smaller image to run the application
# Track version to ease to track the version of the base image
FROM alpine:3.21

# Install necessary packages

# Set the working directory inside the container
WORKDIR /app


# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

CMD ["./main"]