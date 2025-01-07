package lib

import (
	"simple_auth/internal/lib/schema"
	"time"

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
