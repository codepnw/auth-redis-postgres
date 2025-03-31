package models

import (
	"time"

	"github.com/google/uuid"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserReq struct {
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
}
