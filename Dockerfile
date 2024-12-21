# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod  ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

RUN  VAULT_IPS="apple, banana, cherry, 127.0.0.1"
# Ensure the binary is built for Linux and the right architecture (disable CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/telnet-server

# Stage 2: Final stage - using a minimal Alpine image
FROM alpine:alpine

WORKDIR /app/

# Copy the built application from the builder stage
COPY --from=builder /app/telnet-server .

COPY ./static ./static

# Ensure the binary is executable
RUN chmod +x ./telnet-server

# Expose the port the app will run on
EXPOSE 8080

# Run the Go application
CMD ["./telnet-server"]