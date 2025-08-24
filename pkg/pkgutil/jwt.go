package pkgutil

import (
	"fmt"
	"github.com/dimasbayuseno/cisdi-go-test/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWT(id, username, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := config.Get().JwtSecret
	if secretKey == "" {
		return "", fmt.Errorf("failed to create token")
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
