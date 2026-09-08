# model/db/migrate.py

import psycopg2
import os
from pathlib import Path

MIGRATIONS_DIR = Path(__file__).parent / "migrations"

def get_connection():
    url = os.getenv("DATABASE_URL")
    if not url:
        raise Exception("DATABASE_URL not set")
    return psycopg2.connect(url)

def ensure_version_table(conn):
    with conn.cursor() as cur:
        cur.execute("""
            CREATE TABLE IF NOT EXISTS schema_version (
                version INTEGER NOT NULL
            );
        """)
        conn.commit()

        cur.execute("SELECT COUNT(*) FROM schema_version;")
        if cur.fetchone()[0] == 0:
            cur.execute("INSERT INTO schema_version (version) VALUES (0);")
            conn.commit()

def get_current_version(conn):
    with conn.cursor() as cur:
        cur.execute("SELECT version FROM schema_version;")
        return cur.fetchone()[0]

def apply_migration(conn, version, file_path):
    print(f"Applying migration {version}: {file_path.name}")
    sql = file_path.read_text()

    with conn.cursor() as cur:
        cur.execute(sql)
        cur.execute("UPDATE schema_version SET version = %s;", (version,))
        conn.commit()

def run_migrations():
    conn = get_connection()
    ensure_version_table(conn)

    current_version = get_current_version(conn)

    migration_files = sorted(MIGRATIONS_DIR.glob("*.sql"))

    for idx, file_path in enumerate(migration_files, start=1):
        if idx > current_version:
            apply_migration(conn, idx, file_path)

    print("All migrations applied.")
    conn.close()

if __name__ == "__main__":
    run_migrations()
