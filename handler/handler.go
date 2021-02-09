package handler

import "github.com/gumpen/gs1/domain/repository"

type Handler struct {
	UserRepository repository.UserRepository
}

func NewHandler(user repository.UserRepository) *Handler {
	return &Handler{
		UserRepository: user,
	}
}
