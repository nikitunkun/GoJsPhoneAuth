package sqldb

import (
	"backend/internal/app/db"
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (database *Database) User() db.UserRepository {
	if database.userRepository == nil {
		database.userRepository = &UserRepository{
			db: database,
		}
	}

	return database.userRepository
}
