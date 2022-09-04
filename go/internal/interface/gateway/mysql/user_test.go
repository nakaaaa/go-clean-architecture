package mysql

import (
	"reflect"
	"testing"

	"github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	tests := []struct {
		name string
		want *UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetList(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		rp      *UserRepository
		args    args
		want    []model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &UserRepository{}
			got, err := rp.GetList(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetByUserID(t *testing.T) {
	type args struct {
		db     *gorm.DB
		userID int
	}
	tests := []struct {
		name    string
		rp      *UserRepository
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &UserRepository{}
			got, err := rp.GetByUserID(tt.args.db, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_AddOrUpdate(t *testing.T) {
	type args struct {
		db   *gorm.DB
		user *model.User
	}
	tests := []struct {
		name    string
		rp      *UserRepository
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &UserRepository{}
			got, err := rp.AddOrUpdate(tt.args.db, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.AddOrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
