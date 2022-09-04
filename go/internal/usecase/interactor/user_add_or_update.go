package interactor

import (
	"context"

	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/config"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/registry"
	"github.com/nakaaaa/go-clean-architecture/go/internal/usecase"
	"gorm.io/gorm"
)

type UserAddOrUpdateInteractor struct {
	Interactor
}

func NewUserAddOrUpdateInteractor(config *config.Config, db *registry.Database, rp *registry.Repository) *UserAddOrUpdateInteractor {
	return &UserAddOrUpdateInteractor{
		Interactor: NewInteractor(config, db, rp),
	}
}

func (uc *UserAddOrUpdateInteractor) Execute(ctx context.Context, input *usecase.UserAddOrUpdateInputPort) (*usecase.UserAddOrUpdateOutputPort, error) {
	var output *usecase.UserAddOrUpdateOutputPort

	err := uc.WriterTransaction(ctx, func(tx *gorm.DB) error {
		out, err := uc.rp.Users.AddOrUpdate(tx, input.Model)
		if err != nil {
			logger.Errorf("fail to uc.rp.Users.AddOrUpdate(): err=%v", err)
			return err
		}

		output = &usecase.UserAddOrUpdateOutputPort{
			Model: out,
		}

		return nil
	})

	return output, err
}
