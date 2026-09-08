# ------------------------------------------------------------
# BUILD STAGE
# ------------------------------------------------------------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install build tools
RUN apk add --no-cache git build-base

# Copy Go modules first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o edu247api ./cmd/edugo

# ------------------------------------------------------------
# RUNTIME STAGE
# ------------------------------------------------------------
FROM alpine:3.19

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/edu247api /app/edu247api

# Create uploads directory
RUN mkdir -p /app/uploads

# Expose API port
EXPOSE 8080

# Run server
CMD ["/app/edu247api"]
