package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/codepnw/auth-redis-postgres/handlers"
	"github.com/codepnw/auth-redis-postgres/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - no session"})
			return
		}

		sessionDataJSON, err := cfg.RedisClient.Get(c, sessionID).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			return
		}

		var sessionData handlers.SessionData
		err = json.Unmarshal([]byte(sessionDataJSON), &sessionData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode session data"})
			return
		}

		token, err := jwt.ParseWithClaims(sessionData.Token, &handlers.Claims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userID", sessionData.UserID)
		c.Next()
	}
}