package infrastructure

import "github.com/gumpen/gs1/domain"

// 実装
type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r UserRepositoryImpl) Get(id string) (*domain.User, error) {
	return &domain.User{
		ID:   1,
		Name: "hoge",
	}, nil
}

func (r UserRepositoryImpl) GetAll() ([]*domain.User, error) {
	return []*domain.User{
		{
			ID:   1,
			Name: "hoge",
		},
		{
			ID:   2,
			Name: "fuga",
		},
	}, nil
}
