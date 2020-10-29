package mw

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

func SetInterceptors(g *gin.RouterGroup) {
	g.Use(RequestID())
	g.Use(Logger())
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		requestId := c.Writer.Header().Get("Request-Id")
		path := c.Request.URL.Path
		method := c.Request.Method

		log.Info().Str("RequestId", requestId).Str("RealIP", c.Request.RemoteAddr).Str("method", method).Str("path", path).Int("status", c.Writer.Status()).Str("latency", latency.String()).Msg("request")
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
