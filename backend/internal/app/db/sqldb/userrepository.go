package sqldb

import (
	"backend/internal/app/db"
	"backend/internal/app/model"
	"database/sql"
)

type UserRepository struct {
	db *Database
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.db.QueryRow(
		"INSERT INTO users (name, phone) VALUES ($1, $2) RETURNING id",
		user.Name,
		user.Phone,
	).Scan(&user.ID)
}

func (r *UserRepository) Delete(phone string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.db.QueryRow(
		"DELETE FROM users WHERE phone = $1 RETURNING id, name, phone",
		phone,
	).Scan(&user.ID, &user.Name, &user.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Get(phone string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.db.QueryRow(
		"SELECT id, name, phone FROM users WHERE phone = $1",
		phone,
	).Scan(&user.ID, &user.Name, &user.Phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}
