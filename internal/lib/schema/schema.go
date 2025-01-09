package schema

type UserSignUpSchema struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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
