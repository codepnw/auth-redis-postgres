package router

import (
	"github.com/codepnw/auth-redis-postgres/handlers"
	"github.com/codepnw/auth-redis-postgres/internal/config"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func RegisterRoutes(apiConfig *config.ApiConfig) *gin.Engine {
	r = gin.Default()

	hdl := &handlers.LocalApiConfig{
		ApiConfig: apiConfig,
	}

	r.GET("/health-check", hdl.HealthCheck)

	r.POST("/login", hdl.LoginHandler)
	r.POST("/logout", hdl.LogoutHandler)

	return r
}

func Start(port string) *gin.Engine {
	r.Run(":" + port)
	return r
}
