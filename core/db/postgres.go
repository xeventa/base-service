package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQL creates a sql.DB using the provided MySQL/MariaDB DSN.
// Example DSN: user:pass@tcp(host:port)/dbname?parseTime=true&charset=utf8mb4&loc=Local
// If dsn is empty, returns nil without error to allow services that don't need DB.
func NewMySQL(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, nil
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Reasonable defaults; callers can adjust if needed
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	// Ping to validate DSN
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
