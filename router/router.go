package router

import (
	"github.com/codepnw/auth-redis-postgres/handlers"
	"github.com/codepnw/auth-redis-postgres/internal/config"
	"github.com/codepnw/auth-redis-postgres/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func RegisterRoutes(apiConfig *config.ApiConfig) *gin.Engine {
	r = gin.Default()
	r.Use(cors.Default())

	hdl := &handlers.LocalApiConfig{
		ApiConfig: apiConfig,
	}

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware(apiConfig))
	{
		authorized.GET("/health-check", hdl.HealthCheck)
	}

	r.POST("/login", hdl.LoginHandler)
	r.POST("/logout", hdl.LogoutHandler)
	r.POST("/register", hdl.CreateUserHandler)

	return r
}

func Start(port string) *gin.Engine {
	r.Run(":" + port)
	return r
}
