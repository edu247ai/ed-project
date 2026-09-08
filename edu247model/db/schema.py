# model/db/schema.py

SCHEMA = {
    "users": """
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL UNIQUE,
            password_hash TEXT NOT NULL,
            role TEXT NOT NULL DEFAULT 'student',
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """,

    "courses": """
        CREATE TABLE IF NOT EXISTS courses (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            description TEXT,
            teacher_id INTEGER REFERENCES users(id),
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """,

    "lessons": """
        CREATE TABLE IF NOT EXISTS lessons (
            id SERIAL PRIMARY KEY,
            course_id INTEGER REFERENCES courses(id),
            title TEXT NOT NULL,
            content TEXT,
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """,

    "files": """
        CREATE TABLE IF NOT EXISTS files (
            id SERIAL PRIMARY KEY,
            path TEXT NOT NULL,
            owner_id INTEGER REFERENCES users(id),
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """,

    "marketplace_products": """
        CREATE TABLE IF NOT EXISTS marketplace_products (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            description TEXT,
            price NUMERIC NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """,

    "marketplace_orders": """
        CREATE TABLE IF NOT EXISTS marketplace_orders (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id),
            product_id INTEGER REFERENCES marketplace_products(id),
            created_at TIMESTAMP NOT NULL DEFAULT NOW()
        );
    """
}
