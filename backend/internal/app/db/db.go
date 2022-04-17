package db

type Database interface {
	User() UserRepository
}
