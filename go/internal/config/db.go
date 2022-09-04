package config

import (
	"fmt"
	"os"
)

const keyDSN = "DSN"

type DBConfig struct {
	DSN string
}

func NewDBConfig() (*DBConfig, error) {
	dsn, ok := os.LookupEnv(keyDSN)
	if !ok {
		logger.Errorf("env Not Found :%s", keyDSN)
		return nil, fmt.Errorf("env not found: %s", keyDSN)
	}

	return &DBConfig{
		DSN: dsn,
	}, nil
}
