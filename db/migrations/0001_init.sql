-- ============================================================
-- MIGRATION 0001 — CORE IDENTITY + EDUCATION SYSTEM
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
-- END OF MIGRATION 0001
-- ============================================================

