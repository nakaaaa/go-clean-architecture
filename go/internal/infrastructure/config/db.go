package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const keyDSN = "DSN"

type DBConfig struct {
	DSN string
}

func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}))
	if err != nil {
		logger.Errorf("fail to gorm.Open(): err=%v", err)
		return nil, err
	}

	return db, nil
}

func NewDBConfig() (*DBConfig, error) {
	dsn, ok := os.LookupEnv(keyDSN)
	if !ok {
		logger.Errorf("Env Not Found :%s", keyDSN)
		return nil, fmt.Errorf("%s not found", keyDSN)
	}

	return &DBConfig{
		DSN: dsn,
	}, nil
}
