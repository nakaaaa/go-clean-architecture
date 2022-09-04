package registry

import "github.com/nakaaaa/go-clean-architecture/go/internal/usecase"

type Usecase struct {
	UserGetList usecase.UserGetListUseCase
}

func NewUsecase(
	userGetList usecase.UserGetListUseCase,
) *Usecase {
	return &Usecase{
		UserGetList: userGetList,
	}
}
