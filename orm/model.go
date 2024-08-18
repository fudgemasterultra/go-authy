package orm

import (
	"time"

	"gorm.io/gorm"
)

type Response[T any] struct {
	Success bool
	Message string
	Data    *T
}

type ConnectedDb struct {
	Url string
}

type ConnectionData struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Salt     string
	Sessions []SessionToken
}

type SessionToken struct {
	gorm.Model
	Token                 string
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
	RefreshUsed           bool
	ExpiresAt             time.Time
	UserID                int
	User                  User
}
