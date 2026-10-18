package db

import (
	"database/sql"
)

type DB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *DB {
	return &DB{conn: conn}
}

func (d *DB) SaveMetadata(cid string, size int) error {
	_, err := d.conn.Exec(
		"INSERT INTO metadata (cid, size) VALUES ($1, $2) ON CONFLICT (cid) DO NOTHING",
		cid, size,
	)
	return err
}

func (d *DB) GetMetadata(cid string) (int, error) {
	var size int
	err := d.conn.QueryRow("SELECT size FROM metadata WHERE cid=$1", cid).Scan(&size)
	return size, err
}
