package config

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
}

func NewConfig() (*Config, error) {
	server, err := NewServerConfig()
	if err != nil {
		logger.Errorf("fail to NewServerConfig(): err=%v", err)
		return nil, err
	}

	db, err := NewDBConfig()
	if err != nil {
		logger.Errorf("fail to NewDBConfig(): err=%v", err)
		return nil, err
	}

	return &Config{
		Server: server,
		DB:     db,
	}, nil
}
