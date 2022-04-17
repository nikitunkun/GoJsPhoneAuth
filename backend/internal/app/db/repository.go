package db

import "backend/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Delete(string) (*model.User, error)
	Get(string) (*model.User, error)
}
