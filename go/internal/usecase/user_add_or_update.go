package usecase

import "github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"

type UserAddOrUpdateInputPort struct {
	// ユーザー
	Model *model.User
}

type UserAddOrUpdateOutputPort struct {
	// ユーザー
	Model *model.User
}

type UserAddOrUpdateUseCase Usecase[*UserAddOrUpdateInputPort, *UserAddOrUpdateOutputPort]
