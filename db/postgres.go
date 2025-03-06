//go:build postgres

package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	return sql.Open("postgres", "postgresql://root:password@localhost:5433/gobackend?sslmode=disable")
}
