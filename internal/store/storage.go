package store

import "database/sql"

type Storage struct {
	Auth interface {
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Auth: &authStorage{db: db},
	}
}
