package db

import "database/sql"

type IDatabase interface {
	GetDB() (*sql.DB, error)
}
