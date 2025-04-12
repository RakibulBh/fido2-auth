package store

import "database/sql"

type authStorage struct {
	db *sql.DB
}
