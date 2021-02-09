package repository

import (
	"github.com/gumpen/gs1/domain"
)

type UserRepository interface {
	Get(id string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
}
