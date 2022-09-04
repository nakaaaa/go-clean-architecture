package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/api/router"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/config"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/database"
	"github.com/nakaaaa/go-clean-architecture/go/internal/infrastructure/registry"
	"github.com/nakaaaa/go-clean-architecture/go/internal/interface/controller"
	"github.com/nakaaaa/go-clean-architecture/go/internal/interface/gateway/mysql"
	"github.com/nakaaaa/go-clean-architecture/go/internal/usecase/interactor"
)

// サーバ実装
type Server struct {
	Config *config.Config
	Server *gin.Engine
}

func MustNewServer() *Server {
	server, err := NewServer()
	if err != nil {
		logger.Errorf("fail to NewServer(): err=%v", err)
	}

	return server
}

// Server インスタンス返却
func NewServer() (*Server, error) {
	config, err := config.NewConfig()
	if err != nil {
		logger.Errorf("fail to config.NewConfig(): err=%v", err)
		return nil, err
	}

	// Database関連の初期化処理を入れる
	db, err := NewDatabase(config)
	if err != nil {
		logger.Errorf("fail to NewDatabase(): err=%v", err)
		return nil, err
	}

	rp, err := NewRepository()
	if err != nil {
		logger.Errorf("fail to NewRepository(): err=%v", err)
		return nil, err
	}

	uc, err := NewUsecase(config, db, rp)
	if err != nil {
		logger.Errorf("fail to NewUsecase(): err=%v", err)
		return nil, err
	}

	ct, err := NewController(uc)
	if err != nil {
		logger.Errorf("fail to NewController(): err=%v", err)
		return nil, err
	}

	server := gin.Default()
	router.Attach(server.Group("/"), ct)

	return &Server{
		Config: config,
		Server: server,
	}, nil

}

func NewDatabase(config *config.Config) (*registry.Database, error) {
	reader, err := database.Open(config.DB.DSN)
	if err != nil {
		logger.Errorf("fail to database.Open(): err=%v", err)
		return nil, err
	}
	writer, err := database.Open(config.DB.DSN)
	if err != nil {
		logger.Errorf("fail to database.Open(): err=%v", err)
		return nil, err
	}

	return registry.NewDatabase(
		reader,
		writer,
	), nil
}

func NewRepository() (*registry.Repository, error) {
	return registry.NewRepository(
		mysql.NewUserRepository(),
	), nil
}

func NewUsecase(config *config.Config, db *registry.Database, rp *registry.Repository) (*registry.Usecase, error) {
	return registry.NewUsecase(
		interactor.NewUserGetListInteractor(config, db, rp),
	), nil
}

func NewController(uc *registry.Usecase) (*registry.Controller, error) {
	return registry.NewController(
		controller.NewUserGetListController(uc.UserGetList),
	), nil
}
