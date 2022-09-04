package registry

import "github.com/nakaaaa/go-clean-architecture/go/internal/domain/repository"

type Repository struct {
	Users repository.UserRepository
}

func NewRepository(users repository.UserRepository) *Repository {
	return &Repository{
		Users: users,
	}
}
