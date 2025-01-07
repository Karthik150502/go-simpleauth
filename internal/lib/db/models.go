package db

import (
	"time"

	"gorm.io/gorm"
)

const (
	ADMIN = "Admin"
	USER  = "User"
)

type User struct {
	FullName      string `gorm:"not null"`
	Email         string `gorm:"unique;not null" binding:"email"`
	Password      string `gorm:"null"`
	EmailVerified bool   `gorm:"default:false"`
	Role          string `gorm:"default:'User'"`
	gorm.Model
}

type VerificationToken struct {
	UserId  string    `gorm:"not null"`
	Token   string    `gorm:"not null"`
	Expires time.Time `gorm:"not null"`
	gorm.Model
}
