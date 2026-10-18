# ---------------------------------------------------------
# Stage 1: Builder
# ---------------------------------------------------------
FROM golang:1.27.1-alpine AS builder

# Install build tools
RUN apk add --no-cache git bash

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the IPFS server binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o ipfs-server ./cmd/ipfs-node/main.go


# ---------------------------------------------------------
# Stage 2: Runtime
# ---------------------------------------------------------
FROM alpine:3.19

# Install minimal runtime dependencies
RUN apk add --no-cache bash ca-certificates

# Create working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/ipfs-server /usr/local/bin/ipfs-server

# Copy config profiles
COPY config/profiles /app/config/profiles

# Create storage directory
RUN mkdir -p /var/ipfs/dev && \
    mkdir -p /var/ipfs/prod

# Expose ports
# API
EXPOSE 8082
# Dashboard
EXPOSE 8083
# Swarm (future libp2p)
EXPOSE 4002

# Run the server
ENTRYPOINT ["/usr/local/bin/ipfs-server"]
