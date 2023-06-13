package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB представляет подключение к базе данных.
type DB struct {
	conn *sql.DB
}

// NewDB создает новое подключение к базе данных MySQL.
func NewDB() (*DB, error) {
	// Замените значения на свои настройки базы данных MySQL.
	db, err := sql.Open("mysql", "root:password@book(localhost:3306)/book_db")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &DB{
		conn: db,
	}, nil
}

// Close закрывает подключение к базе данных.
func (db *DB) Close() error {
	return db.conn.Close()
}

// Conn возвращает активное подключение к базе данных.
func (db *DB) Conn() *sql.DB {
	return db.conn
}
