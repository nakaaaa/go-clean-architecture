package config

import (
	"fmt"
	"os"
	"strconv"
)

const keyServerPort = "SERVER_PORT"

type ServerConfig struct {
	Port int
}

func NewServerConfig() (*ServerConfig, error) {
	port, ok := os.LookupEnv(keyServerPort)
	if !ok {
		logger.Errorf("env not found: %s", keyServerPort)
		return nil, fmt.Errorf("env not found: %s", keyServerPort)
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		logger.Errorf("fail to strconv.Atoi(): err=%v", err)
		return nil, err
	}

	return &ServerConfig{
		Port: p,
	}, nil
}
