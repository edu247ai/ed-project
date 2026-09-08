-- ============================================================
-- MIGRATION 0003 — CHAT SYSTEM
-- ============================================================

-- ============================================================
-- ENUMS
-- ============================================================

CREATE TYPE chat_room_type AS ENUM ('course_room', 'support', 'private');

-- ============================================================
-- CHAT ROOMS
-- ============================================================

CREATE TABLE chat_rooms (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type chat_room_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- CHAT ROOM MEMBERS
-- ============================================================

CREATE TABLE chat_room_members (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- CHAT MESSAGES
-- ============================================================

CREATE TABLE chat_messages (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    sender_user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- END OF MIGRATION 0003
-- ============================================================
