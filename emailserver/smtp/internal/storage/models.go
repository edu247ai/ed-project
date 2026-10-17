package storage

import "time"

type Message struct {
	ID        string // UUID
	Sender    string
	Body      []byte
	CreatedAt time.Time
}

type Envelope struct {
	ID        string // UUID
	MessageID string
	Helo      string
	MailFrom  string
	CreatedAt time.Time
}

type Recipient struct {
	ID         string // UUID
	EnvelopeID string
	Address    string
	CreatedAt  time.Time
}
