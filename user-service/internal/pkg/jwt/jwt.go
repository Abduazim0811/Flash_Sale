package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)
var (
	secretKey = []byte("abduazim11")
)

func GenerateJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":      email,
			"created_at": time.Now().Unix(),
			// "exp":        time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}