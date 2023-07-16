package test

import (
	"fmt"

	"gorm.io/gorm"
)

func DeleteAll(db *gorm.DB) error {
	tables := []string{
		"users",
	}

	for _, table := range tables {
		err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE 1=1", table)).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func InsertAll(db *gorm.DB, models []interface{}) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		for _, model := range models {
			err := tx.Create(model).Error
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
