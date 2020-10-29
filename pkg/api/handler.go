package api

import (
	"github.com/bilalkocoglu/file-service/pkg/mw"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(g *gin.RouterGroup) {
	g.POST("/login", Login)

	g.GET("/app-user/all", mw.Auth(), AllUsers)
	g.POST("/app-user/", mw.Auth(), SaveUser)
	g.GET("/app-user/by-id/:id", mw.Auth(), FindUserById)
}
