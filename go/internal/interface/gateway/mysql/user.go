package mysql

import (
	"github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct{}

// NewUserRepositoryはUserレポジトリの新規インスタンスを返却する
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (rp *UserRepository) GetList(db *gorm.DB) ([]model.User, error) {
	var list []model.User

	err := db.Find(&list).Error
	if err != nil {
		logger.Errorf("fail to db.Find(): err=%v", err)
		return nil, err
	}

	return list, nil
}

func (rp *UserRepository) GetByUserID(db *gorm.DB, userID int) (*model.User, error) {
	var model model.User

	err := db.Where("`uid`=?", userID).First(model).Error
	if err != nil {
		logger.Errorf("fail to db.First(): err=%v", err)
		return nil, err
	}

	return &model, nil
}

func (rp *UserRepository) AddOrUpdate(db *gorm.DB, user *model.User) (*model.User, error) {
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user).Error
	if err != nil {
		logger.Errorf("fail to db.Clauses(); err=%v", err)
		return nil, err
	}

	return user, nil
}
