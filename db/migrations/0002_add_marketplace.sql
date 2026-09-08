-- marketplace migration
-- ============================================================
-- MIGRATION 0002 — MARKETPLACE SYSTEM
-- ============================================================

-- ============================================================
-- PRODUCT TYPES & STATUS ENUMS
-- ============================================================

CREATE TYPE product_type AS ENUM ('ebook', 'audiobook', 'coffee_audio');
CREATE TYPE product_status AS ENUM ('active', 'inactive');

CREATE TYPE file_visibility AS ENUM ('private', 'course_only', 'purchased_only', 'public');

CREATE TYPE order_status AS ENUM ('pending', 'paid', 'cancelled', 'refunded');

-- ============================================================
-- FILE STORAGE TABLE
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
-- PRODUCTS TABLE
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

-- ============================================================
-- PRODUCT FILES (EBOOKS, AUDIOBOOKS, COFFEE AUDIOS)
-- ============================================================

CREATE TABLE product_files (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    file_id INT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
    format TEXT,
    quality TEXT
);

-- ============================================================
-- ORDERS TABLE
-- ============================================================

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total_price_cents INT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'ZAR',
    status order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================
-- ORDER ITEMS TABLE
-- ============================================================

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(id),
    unit_price_cents INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1
);

-- ============================================================
-- END OF MIGRATION 0002
-- ============================================================
