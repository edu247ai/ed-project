# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /out && CGO_ENABLED=0 GOOS=linux go build -o /out/smtpd ./cmd/smtpd

# Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /out/smtpd /usr/local/bin/smtpd

ENV SMTP_LISTEN_ADDRESS=:2525

EXPOSE 2525

ENTRYPOINT ["/usr/local/bin/smtpd"]
