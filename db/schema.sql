-- ============================================================
-- EDU247 — FULL DATABASE SCHEMA
-- ============================================================

-- ============================================================
-- ENUMS
-- ============================================================

CREATE TYPE user_role AS ENUM ('student', 'teacher', 'admin');
CREATE TYPE user_status AS ENUM ('active', 'suspended', 'pending_verification');

CREATE TYPE teacher_verification AS ENUM ('pending', 'verified', 'rejected');

CREATE TYPE course_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE lesson_type AS ENUM ('cartoon_site', 'video', 'pdf', 'quiz');

CREATE TYPE progress_status AS ENUM ('not_started', 'in_progress', 'completed');

CREATE TYPE product_type AS ENUM ('ebook', 'audiobook', 'coffee_audio');
CREATE TYPE product_status AS ENUM ('active', 'inactive');

CREATE TYPE order_status AS ENUM ('pending', 'paid', 'cancelled', 'refunded');

CREATE TYPE file_visibility AS ENUM ('private', 'course_only', 'purchased_only', 'public');

CREATE TYPE chat_room_type AS ENUM ('course_room', 'support', 'private');

CREATE TYPE event_type AS ENUM (
    'view_course',
    'view_product',
    'purchase',
    'chat_message',
    'lesson_complete'
);

-- ============================================================
-- USERS
-- ============================================================

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role user_role NOT NULL,
    status user_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- TEACHER PROFILES
-- ============================================================

CREATE TABLE teacher_profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    full_name TEXT NOT NULL,
    bio TEXT,
    subjects JSONB,
    verification_status teacher_verification NOT NULL DEFAULT 'pending',
    verification_notes TEXT
);

-- ============================================================
-- STUDENT PROFILES
-- ============================================================

CREATE TABLE student_profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    grade_level TEXT,
    parent_contact TEXT
);

-- ============================================================
-- COURSES
-- ============================================================

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL REFERENCES teacher_profiles(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    age_range TEXT,
    status course_status NOT NULL DEFAULT 'draft',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- LESSONS
-- ============================================================

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content_type lesson_type NOT NULL,
    content_url TEXT NOT NULL,
    order_index INT NOT NULL DEFAULT 0
);

-- ============================================================
-- STUDENT PROGRESS
-- ============================================================

CREATE TABLE student_progress (
    id SERIAL PRIMARY KEY,
    student_user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    status progress_status NOT NULL DEFAULT 'not_started',
    last_accessed_at TIMESTAMPTZ
);

-- ============================================================
-- FILE STORAGE
-- ============================================================

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    owner_user_id INT REFERENCES users(id) ON DELETE SET NULL,
    original_name TEXT NOT NULL,
    storage_path TEXT NOT NULL,
    mime_type TEXT,
    size_bytes BIGINT,
    visibility file_visibility NOT NULL DEFAULT 'private',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- MARKETPLACE PRODUCTS
-- ============================================================

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    type product_type NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    price_cents INT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'ZAR',
    status product_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE product_files (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    file_id INT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
    format TEXT,
    quality TEXT
);

-- ============================================================
-- ORDERS
-- ============================================================

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total_price_cents INT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'ZAR',
    status order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(id),
    unit_price_cents INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1
);

-- ============================================================
-- CHAT SYSTEM
-- ============================================================

CREATE TABLE chat_rooms (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type chat_room_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE chat_room_members (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE chat_messages (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL REFERENCES chat_rooms(id) ON DELETE CASCADE,
    sender_user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- ML TELEMETRY
-- ============================================================

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    event_type event_type NOT NULL,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE recommendation_logs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recommended_items JSONB NOT NULL,
    context TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- SECURITY / AUDIT
-- ============================================================

CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    actor_user_id INT REFERENCES users(id) ON DELETE SET NULL,
    action TEXT NOT NULL,
    target TEXT,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

-- ============================================================
-- END OF SCHEMA
-- ============================================================

