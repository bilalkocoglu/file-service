package config

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/file-service/pkg/api"
	"github.com/imminoglobulin/file-service/pkg/mw"
)

type Server struct {
	cfg *Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func PrepareServer(config *Config) *gin.Engine {
	router := gin.Default()
	router.Use(mw.CORSMiddleware())
	router.MaxMultipartMemory = 100 << 20 // 100MiB

	g := router.Group("/v1")

	mw.SetInterceptors(g)
	api.RegisterHandlers(g)

	return router
}
