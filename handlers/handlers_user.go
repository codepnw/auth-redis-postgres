package handlers

import (
	"net/http"
	"time"

	"github.com/codepnw/auth-redis-postgres/internal/config"
	"github.com/codepnw/auth-redis-postgres/internal/database"
	"github.com/codepnw/auth-redis-postgres/internal/models"
	"github.com/codepnw/auth-redis-postgres/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LocalApiConfig struct {
	*config.ApiConfig
}

func (l *LocalApiConfig) CreateUserHandler(c *gin.Context) {
	req := models.UserReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser, err := l.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		Name:      req.Name,
		Username:  req.Username,
		Email:     req.Email,
		Password:  password,
		CreatedAt: time.Now().UTC().Local(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
