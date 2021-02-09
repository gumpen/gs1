package infrastructure

import (
	"github.com/gumpen/gs1/domain/repository"
)

type Repositories struct {
	UserRepository repository.UserRepository
}

// 実装の集合
func NewRepositories() *Repositories {
	return &Repositories{
		UserRepository: NewUserRepositoryImpl(),
	}
}
