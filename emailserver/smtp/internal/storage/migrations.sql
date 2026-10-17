CREATE TABLE IF NOT EXISTS messages (
    id          UUID PRIMARY KEY,
    sender      TEXT NOT NULL,
    body        BYTEA NOT NULL,
    created_at  TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS envelopes (
    id          UUID PRIMARY KEY,
    message_id  UUID REFERENCES messages(id) ON DELETE CASCADE,
    helo        TEXT,
    mail_from   TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS recipients (
    id           UUID PRIMARY KEY,
    envelope_id  UUID REFERENCES envelopes(id) ON DELETE CASCADE,
    address      TEXT NOT NULL,
    created_at   TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_messages_created ON messages(created_at);
CREATE INDEX IF NOT EXISTS idx_recipients_address ON recipients(address);

CREATE TABLE IF NOT EXISTS queue_items (
    id           UUID PRIMARY KEY,
    message_id   UUID NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    recipient    TEXT NOT NULL,
    attempts     INTEGER NOT NULL DEFAULT 0,
    next_attempt TIMESTAMP NOT NULL,
    last_error   TEXT NOT NULL DEFAULT '',
    created_at   TIMESTAMP NOT NULL
);

ALTER TABLE queue_items ADD COLUMN IF NOT EXISTS last_error TEXT NOT NULL DEFAULT '';
ALTER TABLE queue_items ADD COLUMN IF NOT EXISTS lease_until TIMESTAMP;

CREATE INDEX IF NOT EXISTS idx_queue_items_next_attempt ON queue_items(next_attempt);
