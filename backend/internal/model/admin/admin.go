package admin

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int            `json:"id"`
	Username     string         `json:"username"`
	Password     string         `json:"-"`
	PasswordHash string         `json:"-"`
	Role         string         `json:"role"`
	Category     sql.NullString `json:"category,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
}
