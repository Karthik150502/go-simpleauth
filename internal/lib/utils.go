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

func GenerateJWT(user *schema.UserJwtPayloadSchema, expiresIn time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expiresIn).Unix()
	claims["authorized"] = true
	claims["user"] = user

	tokenString, err := token.SignedString(sampleJwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
