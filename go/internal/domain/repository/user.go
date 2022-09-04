package repository

import (
	"github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetList(db *gorm.DB) ([]model.User, error)

	GetByUserID(db *gorm.DB, userID int) (*model.User, error)

	AddOrUpdate(db *gorm.DB, user *model.User) (*model.User, error)
}
