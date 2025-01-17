package lib

import (
	"simple_auth/internal/lib/schema"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

var sampleJwtSecretKey = []byte("SecretYouShouldHide")

// "github.com/golang-jwt/jwt/v5"
func DecodeJwt(tokenString string) {

}

type JWTToken struct {
	ExpiresAt *time.Time
	Token     *string
	Error     error
}

func GenerateJWT(user *schema.UserJwtPayloadSchema, expiresIn time.Duration) *JWTToken {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expiresAt := time.Now().Add(expiresIn)
	claims["exp"] = expiresAt.Unix()
	claims["authorized"] = true
	claims["user"] = user

	tokenString, err := token.SignedString(sampleJwtSecretKey)
	if err != nil {
		return &JWTToken{
			Token:     nil,
			ExpiresAt: nil,
			Error:     err,
		}
	}
	return &JWTToken{
		Token:     &tokenString,
		ExpiresAt: &expiresAt,
		Error:     err,
	}
}

func ValidateUserInput(schema interface{}) map[string]string {
	var validate = validator.New()
	err := validate.Struct(schema)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		response := make(map[string]string)
		for _, ve := range validationErrors {
			response[ve.Field()] = ve.ActualTag()
		}
		return response
	}
	return nil
}
