package schema

import (
	"simple_auth/internal/lib/db"
	"time"
)

type UserSignUpSchema struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"omitempty"`
}

type SessionSchema struct {
	User         db.User `json:"userId" validate:"required"`
	RefreshToken string  `json:"refreshToken" validate:"required"`
}

type VerificationTokenSchema struct {
	User      db.User   `json:"userId" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	ExpiresAt time.Time `json:"expiresAt" validate:"required"`
}

type UserSignInSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserJwtPayloadSchema struct {
	FullName        string `json:"fullName"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}
