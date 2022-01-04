package todo

import "database/sql"

type Repository interface {
}

type repository struct {
	connDB *sql.DB
}

func NewRepository(connDB *sql.DB) Repository {
	return &repository{
		connDB: connDB,
	}
}
