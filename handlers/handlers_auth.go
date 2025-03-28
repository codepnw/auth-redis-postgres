package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/codepnw/auth-redis-postgres/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTRes struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type SessionData struct {
	Token  string    `json:"token"`
	UserID uuid.UUID `json:"user_id"`
}

func (l *LocalApiConfig) LoginHandler(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := l.DB.FindUserByEmail(c, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	expiresTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Email: foundUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sessionID := uuid.New().String()

	sessionData := map[string]any{
		"token": tokenString,
		"user_id": foundUser.ID,
	}

	sessionDataJson, err := json.Marshal(sessionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed encode the session data"})
		return
	}

	err = l.RedisClient.Set(c, sessionID, sessionDataJson, time.Until(expiresTime)).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session data to the redis"})
		return
	}

	maxAge := int(time.Until(expiresTime).Seconds())
	c.SetCookie("session_id", sessionID, maxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login success", 
		"expires": expiresTime,
	})
}

func (l *LocalApiConfig) LogoutHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := l.RedisClient.Del(c, sessionID).Err(); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed delete session"})
		return
	}

	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}