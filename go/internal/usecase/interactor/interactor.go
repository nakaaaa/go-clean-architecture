package interactor

import (
	"context"

	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/config"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/registry"
	"gorm.io/gorm"
)

type Interactor struct {
	config *config.Config
	db     *registry.Database
	rp     *registry.Repository
}

func NewInteractor(config *config.Config, db *registry.Database, rp *registry.Repository) Interactor {
	return Interactor{
		config: config,
		db:     db,
		rp:     rp,
	}
}

func (uc *Interactor) ReaderTransaction(ctx context.Context, fc func(*gorm.DB) error) error {
	err := uc.db.Reader.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
	if err != nil {
		logger.Errorf("fail to uc.db.Reader.Transaction(): err=%v", err)
	}

	return err
}

func (uc *Interactor) WriterTransaction(ctx context.Context, fc func(*gorm.DB) error) error {
	err := uc.db.Writer.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	})
	if err != nil {
		logger.Errorf("fail to uc.db.Writer.Transaction(): err=%v", err)
	}

	return err
}
