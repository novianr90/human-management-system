# ---- Build Stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .

# Ensure all dependencies are available (remove -mod=vendor if you don’t use vendoring)
RUN go build -o /app/main

# ---- Runtime Stage ----
FROM alpine:latest

WORKDIR /app

# Install required runtime dependencies (like SSL certificates)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the application port
EXPOSE 8080

# Ensure the binary is executable
RUN chmod +x /app/main

# Set entrypoint
ENTRYPOINT [ "/app/main" ]
