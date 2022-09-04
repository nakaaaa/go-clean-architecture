package interactor

import (
	"context"

	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/config"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/registry"
	"github.com/nakaaaa/go-clean-architecture/go/internal/usecase"
	"gorm.io/gorm"
)

type UserGetListInteractor struct {
	Interactor
}

func NewUserGetListInteractor(config *config.Config, db *registry.Database, rp *registry.Repository) *UserGetListInteractor {
	return &UserGetListInteractor{
		Interactor: NewInteractor(config, db, rp),
	}
}

func (uc *UserGetListInteractor) Execute(ctx context.Context, input *usecase.UserGetListInputPort) (*usecase.UserGetListOutputPort, error) {
	var output *usecase.UserGetListOutputPort

	err := uc.ReaderTransaction(ctx, func(tx *gorm.DB) error {
		out, err := uc.rp.Users.GetList(tx)
		if err != nil {
			logger.Errorf("fail to uc.rp.Users.GetList(): err=%", err)
			return err
		}

		output = &usecase.UserGetListOutputPort{
			List: out,
		}

		return nil
	})

	return output, err
}
