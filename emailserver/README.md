# Email Server

**A Go-based mail platform under development**

The project brings together services for inbound and outbound email, mailbox
access, webmail, administration, and spam filtering. The SMTP service can run
locally with in-memory storage or use PostgreSQL for persistent messages and
queue entries. Other services need separate validation before production use.

## Platform at a glance

| Component | Purpose |
| --- | --- |
| SMTP | Accepts inbound mail and applies authentication and spam checks |
| Queue worker | Delivers outbound mail and schedules retries |
| IMAP / POP3 | Provides mailbox access |
| Webmail | Browser-based mail client |
| Admin API | Manages domains and users |
| Spam filtering | Combines Rspamd with internal filters |
| Storage | PostgreSQL data with planned S3-compatible attachment storage |

## Project structure

```text
emailserver/
├── smtp/
│   ├── cmd/smtpd/                  # SMTP server entry point
│   ├── internal/
│   │   ├── handlers/               # SMTP commands and sessions
│   │   ├── server/                 # Listener and connection lifecycle
│   │   ├── storage/                # Message and queue persistence
│   │   ├── queue/                  # Outbound queue worker and retry logic
│   │   ├── spf/                    # SPF evaluation
│   │   ├── dkim/                   # DKIM verification and signing
│   │   ├── dmarc/                  # DMARC policy evaluation
│   │   └── spam/                   # Rspamd checkv2 client
│   ├── pkg/
│   │   ├── config/                 # Environment-based configuration
│   │   └── logger/                 # SMTP logging
│   └── Dockerfile
├── imap/                           # IMAP service
├── pop3/                           # POP3 service
├── webmail/                        # Webmail application
├── admin/                          # Admin API
├── spam/                           # Internal filters and Rspamd configuration
├── storage/                        # Shared storage components
└── configs/                        # Service configuration and key mounts
```

## SMTP service

The SMTP server handles `HELO`/`EHLO`, `MAIL FROM`, `RCPT TO`, and `DATA`.
It can evaluate SPF, verify inbound DKIM signatures, apply DMARC policy, and
optionally submit messages to Rspamd. With PostgreSQL configured, accepted
messages and per-recipient queue entries are stored transactionally. Without
PostgreSQL, the service uses volatile in-memory storage for local testing.

The separate queue worker delivers queued messages using MX records, retries
failed deliveries, and can DKIM-sign outbound mail. Inbound DKIM verification
currently supports RSA-SHA256 with relaxed-header/simple-body canonicalization.

## Authentication and spam filtering

- **SPF** checks the connecting IP against the envelope sender domain's DNS
  policy.
- **DKIM** verifies supported inbound signatures. The queue worker can sign
  outgoing messages when its key, domain, and selector are configured.
- **DMARC** applies policy using exact-domain alignment with the message's
  `From` domain.
- **Rspamd** can score messages through its HTTP `checkv2` endpoint. Internal
  spam filters and Rspamd configuration are located under `spam/`.

## Storage

PostgreSQL stores messages, envelopes, recipients, and per-recipient
`queue_items`, including retry state. The SMTP schema is applied at startup.
S3-compatible attachment storage is part of the broader platform architecture
and is not exercised by the SMTP service setup below.

## Run SMTP locally

Run these commands from `emailserver/smtp/`:

```sh
go test ./...
go build -o smtpd ./cmd/smtpd
./smtpd
```

By default, SMTP listens on `:2525` and stores messages in memory. For persistent
storage, configure PostgreSQL with `SMTP_POSTGRES_DSN`; the queue worker must use
the same DSN and runs separately:

```sh
SMTP_POSTGRES_DSN='postgres://user:password@localhost:5432/email?sslmode=disable' \
  go run ./internal/queue/cmd/queued
```

### Configuration

| Variable | Default | Description |
| --- | --- | --- |
| `SMTP_HOSTNAME` | `localhost` | SMTP greeting and outbound EHLO name |
| `SMTP_LISTEN_ADDRESS` | `:2525` | TCP listen address |
| `SMTP_POSTGRES_DSN` | Unset | PostgreSQL DSN; unset selects volatile in-memory storage |
| `SMTP_MAX_MESSAGE_SIZE` | `20971520` | Maximum SMTP `DATA` size in bytes |
| `SMTP_MAX_CONNECTIONS` | `500` | Maximum number of concurrent sessions |
| `SMTP_READ_TIMEOUT` | `5m` | Command and message read timeout |
| `SMTP_WRITE_TIMEOUT` | `30s` | Response write timeout |
| `SMTP_QUEUE_INTERVAL` | `30s` | Queue worker polling interval |
| `SMTP_ENABLE_SPF` | `true` | Enable SPF evaluation |
| `SMTP_ENABLE_DKIM` | `true` | Enable inbound DKIM verification |
| `SMTP_ENABLE_DMARC` | `true` | Enable DMARC policy evaluation |
| `SMTP_ENABLE_SPAM` | `false` | Enable Rspamd checks |
| `SMTP_RSPAMD_URL` | `http://127.0.0.1:11333/checkv2` | Rspamd endpoint |
| `SMTP_DKIM_KEY_PATH` | Unset | RSA private-key path used by the queue worker |
| `SMTP_DKIM_DOMAIN` | Unset | Outbound DKIM signing domain |
| `SMTP_DKIM_SELECTOR` | Unset | Outbound DKIM selector |

> **Note:** In-memory mode loses messages when the process exits. Use PostgreSQL
> for persistent mail handling. Authentication policy behavior, operational
> controls, and the other platform services need deployment-specific validation
> before production use.

## Run SMTP in Docker

Build from `emailserver/smtp/`:

```sh
docker build -t smtpd-service .
docker run --rm -p 2525:2525 \
  -e SMTP_LISTEN_ADDRESS=:2525 \
  -e SMTP_ENABLE_SPAM=true \
  -e SMTP_RSPAMD_URL=http://rspamd:11333/checkv2 \
  -e SMTP_POSTGRES_DSN='postgres://user:password@postgres:5432/email?sslmode=disable' \
  smtpd-service
```

Configure Rspamd with a hostname reachable from the container; `127.0.0.1`
refers to the SMTP container itself. For outbound DKIM signing, mount the
private key read-only into the queue-worker container (for example,
`/run/secrets/dkim.pem`) and set `SMTP_DKIM_KEY_PATH`,
`SMTP_DKIM_DOMAIN`, and `SMTP_DKIM_SELECTOR`.
