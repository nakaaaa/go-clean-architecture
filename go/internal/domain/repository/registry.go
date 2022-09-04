package repository

type Registry struct {
	User UserRepository
}

func NewRegistry(
	user UserRepository,
) *Registry {
	return &Registry{
		User: user,
	}
}
