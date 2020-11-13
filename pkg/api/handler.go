package api

import (
	"github.com/gin-gonic/gin"
	"github.com/imminoglobulin/file-service/pkg/mw"
)

func RegisterHandlers(g *gin.RouterGroup) {
	g.POST("/login", Login)

	g.GET("/app-user/all", mw.Auth(), AllUsers)
	g.POST("/app-user/", mw.Auth(), SaveUser)
	g.GET("/app-user/by-id/:id", mw.Auth(), FindUserById)
	g.DELETE("/app-user/:id", mw.Auth(), DeleteUser)

	g.POST("/file/upload", mw.Auth(), UploadFile)
}
