package registry

import "gorm.io/gorm"

type Database struct {
	Reader *gorm.DB
	Writer *gorm.DB
}

func NewDatabase(
	reader *gorm.DB,
	writer *gorm.DB,
) *Database {
	return &Database{
		Reader: reader,
		Writer: writer,
	}
}
