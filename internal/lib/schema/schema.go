package schema

type UserSignUpSchema struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserSignInSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserJwtPayloadSchema struct {
	FullName        string `json:"fullName"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

