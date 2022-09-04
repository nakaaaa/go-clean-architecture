package model

type User struct {
	// ユーザーID
	UserID int `gorm:"primaryKey;column:uid"`

	// ユーザー名
	Name string

	// 年齢
	Age uint32

	// 性別　0:男 1:女
	Gendar uint32
}

func (u User) TableName() string {
	return "users"
}
