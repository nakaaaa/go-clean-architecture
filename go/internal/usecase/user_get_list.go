package usecase

import "github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"

type UserGetListInputPort struct{}

type UserGetListOutputPort struct {
	// ユーザー一覧
	List []model.User
}

type UserGetListUseCase Usecase[*UserGetListInputPort, *UserGetListOutputPort]
