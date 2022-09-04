package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routing struct {
	Gin  *gin.Engine
	Port string
}

func NewRouter(port string) *Routing {
	r := &Routing{
		Gin:  gin.Default(),
		Port: port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	r.Gin.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
