package main

import (
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/api"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/log"
)

var logger = log.GetLogger()

func main() {
	err := NewServer()
	if err != nil {
		logger.Errorf("fail to NewServer(): err=%v", err)
		panic(err.Error())
	}
}

func NewServer() error {
	// server, err := config.NewServerConfig()
	// if err != nil {
	// 	logger.Errorf("fail to config.NewServerConfig(): err=%v", err)
	// 	return err
	// }

	// port := fmt.Sprintf(":%d", server.Port)
	// r := config.NewRouter(port)
	// r.Run()

	server := api.MustNewServer()
	server.Server.Run()

	return nil
}
