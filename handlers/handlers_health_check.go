package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *LocalApiConfig) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}