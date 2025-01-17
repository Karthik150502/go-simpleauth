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
	FullName          string         `gorm:"not null" json:"fullName"`
	Email             string         `gorm:"unique;not null" json:"email"`
	Password          *string        `json:"password"`
	EmailVerified     bool           `gorm:"default:false" json:"emailVerified"`
	OAuthProvider     *string        `json:"oauthProvider"`
	OAuthProviderId   *string        `gorm:"unique" json:"oauthProviderId"`
	Role              string         `gorm:"default:User"`
	Session           *string        `json:"session"`
	OauthToken        *string        `json:"oauthToken"`
	VerificationToken *string        `json:"verificationToken"`
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}

type Session struct {
	UserID       uint           `json:"userId"`
	User         User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RefreshToken string         `gorm:"not null" json:"refreshToken"`
	ExpiresAt    time.Time      `gorm:"not null" json:"expiresAt"`
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

type OauthToken struct {
	UserID       uint           `json:"userId"`
	User         User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	AccessToken  string         `gorm:"not null" json:"accessToken"`
	RefreshToken string         `gorm:"not null" json:"refreshToken"`
	ExpiresAt    time.Time      `gorm:"not null" json:"expiresAt"`
	Provider     string         `gorm:"not null" json:"provider"`
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

type VerificationToken struct {
	UserID    uint      `json:"userId"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Token     string    `gorm:"not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires"`
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
