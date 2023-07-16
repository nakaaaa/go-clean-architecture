package mysql

import (
	"reflect"
	"testing"

	"github.com/ahmetalpbalkan/go-linq"
	"github.com/go-playground/assert/v2"
	"github.com/nakaaaa/go-clean-architecture/go/internal/domain/model"
	"github.com/nakaaaa/go-clean-architecture/go/internal/test"
	"github.com/nakaaaa/go-clean-architecture/go/internal/test/fixture"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	tests := []struct {
		name string
		want *UserRepository
	}{
		{
			name: "新規インスタンスが返却される",
			want: &UserRepository{},
		},
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
	type testArgs struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    testArgs
		up      func(db *gorm.DB) error
		want    []model.User
		wantErr bool
	}{
		{
			name: "users からレコードを全件取得する",
			args: testArgs{
				db: testdb,
			},
			up: func(db *gorm.DB) error {
				return test.InsertAll(db, []interface{}{
					fixture.NewUser(&model.User{UserID: 1}),
					fixture.NewUser(&model.User{UserID: 2}),
					fixture.NewUser(&model.User{UserID: 3}),
				})
			},
			want: []model.User{
				{
					UserID: 1,
					Name:   "TestUser1",
					Age:    20,
					Gender: model.UserGenderMan,
				},
				{
					UserID: 2,
					Name:   "TestUser2",
					Age:    20,
					Gender: model.UserGenderMan,
				},
				{
					UserID: 3,
					Name:   "TestUser3",
					Age:    20,
					Gender: model.UserGenderMan,
				},
			},
		},
		{
			name: "users にデータが存在しないとき、 空配列が返却される",
			args: testArgs{
				db: testdb,
			},
			up: func(db *gorm.DB) error {
				return nil
			},
			want: []model.User{},
		},
	}
	for _, tt := range tests {
		err := test.DeleteAll(tt.args.db)
		require.NoError(t, err, "failed to test.DeleteAll()")

		err = tt.up(tt.args.db)
		require.NoError(t, err, "failed to tt.up()")

		t.Run(tt.name, func(t *testing.T) {
			rp := NewUserRepository()
			got, err := rp.GetList(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserRepository_GetByUserID(t *testing.T) {
	type testArgs struct {
		db     *gorm.DB
		userID int
	}
	tests := []struct {
		name    string
		args    testArgs
		up      func(db *gorm.DB) error
		want    *model.User
		wantErr bool
	}{
		{
			name: "指定された条件にマッチするレコードを 1 件取得する",
			args: testArgs{
				db:     testdb,
				userID: 2,
			},
			up: func(db *gorm.DB) error {
				return test.InsertAll(db, []interface{}{
					fixture.NewUser(&model.User{UserID: 1}),
					fixture.NewUser(&model.User{UserID: 2}),
					fixture.NewUser(&model.User{UserID: 3}),
				})
			},
			want: fixture.NewUser(&model.User{UserID: 2}),
		},
		{
			name: "指定された条件にマッチするレコードが存在しないとき、エラーにならない",
			args: testArgs{
				db:     testdb,
				userID: 4,
			},
			up: func(db *gorm.DB) error {
				return test.InsertAll(db, []interface{}{
					fixture.NewUser(&model.User{UserID: 1}),
					fixture.NewUser(&model.User{UserID: 2}),
					fixture.NewUser(&model.User{UserID: 3}),
				})
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		err := test.DeleteAll(tt.args.db)
		require.NoError(t, err, "failed to test.DeleteAll()")

		err = tt.up(tt.args.db)
		require.NoError(t, err, "failed to tt.up()")

		t.Run(tt.name, func(t *testing.T) {
			rp := &UserRepository{}
			got, err := rp.GetByUserID(tt.args.db, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserRepository_AddOrUpdate(t *testing.T) {
	type testArgs struct {
		db   *gorm.DB
		user *model.User
	}
	tests := []struct {
		name    string
		args    testArgs
		up      func(db *gorm.DB) error
		want    *model.User
		wantErr bool
		more    func(rp *UserRepository, db *gorm.DB)
	}{
		{
			name: "レコード 1 件追加に成功する",
			args: testArgs{
				db: testdb,
				user: &model.User{
					Name:   "TestInsert",
					Age:    25,
					Gender: model.UserGenderWoman,
				},
			},
			up: func(db *gorm.DB) error {
				return test.InsertAll(db, []interface{}{
					fixture.NewUser(&model.User{UserID: 1}),
					fixture.NewUser(&model.User{UserID: 2}),
					fixture.NewUser(&model.User{UserID: 3}),
				})
			},
			want: &model.User{
				UserID: 4,
				Name:   "TestInsert",
				Age:    25,
				Gender: model.UserGenderWoman,
			},
			more: func(rp *UserRepository, db *gorm.DB) {
				users, err := rp.GetList(db)
				require.NoError(t, err, "failed to rp.GetList()")

				assert.Equal(t, 4, len(users))
				has := linq.From(users).AnyWith(
					func(i interface{}) bool {
						return i.(model.User).Name == "TestInsert"
					},
				)
				assert.Equal(t, true, has)
			},
		},
		{
			name: "指定した条件にマッチするレコードを更新する",
			args: testArgs{
				db: testdb,
				user: &model.User{
					UserID: 2,
					Name:   "TestUpdate",
					Age:    22,
					Gender: model.UserGenderWoman,
				},
			},
			up: func(db *gorm.DB) error {
				return test.InsertAll(db, []interface{}{
					fixture.NewUser(&model.User{UserID: 1}),
					fixture.NewUser(&model.User{UserID: 2}),
					fixture.NewUser(&model.User{UserID: 3}),
				})
			},
			want: &model.User{
				UserID: 2,
				Name:   "TestUpdate",
				Age:    22,
				Gender: model.UserGenderWoman,
			},
			more: func(rp *UserRepository, db *gorm.DB) {
				users, err := rp.GetList(db)
				require.NoError(t, err, "failed to rp.GetList()")

				assert.Equal(t, 3, len(users))

				has := linq.From(users).AnyWith(
					func(i interface{}) bool {
						return i.(model.User).Name == "TestUpdate"
					},
				)
				assert.Equal(t, true, has)
			},
		},
	}
	for _, tt := range tests {
		err := test.DeleteAll(tt.args.db)
		require.NoError(t, err, "failed to test.DeleteAll()")

		err = tt.up(tt.args.db)
		require.NoError(t, err, "failed to tt.up()")

		t.Run(tt.name, func(t *testing.T) {
			rp := NewUserRepository()
			got, err := rp.AddOrUpdate(tt.args.db, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.AddOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)

			if tt.more != nil {
				tt.more(rp, tt.args.db)
			}
		})
	}
}
