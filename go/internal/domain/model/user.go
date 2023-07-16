package model

type User struct {
	// ユーザーID
	UserID int `gorm:"primaryKey;column:uid"`

	// ユーザー名
	Name string

	// 年齢
	Age uint32

	// 性別　0:男 1:女
	Gender uint32
}

const (
	UserGenderMan   = 0
	UserGenderWoman = 1
)

func (u User) TableName() string {
	return "users"
}
