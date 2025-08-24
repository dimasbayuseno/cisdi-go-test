package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type JWTConfig struct {
	Secret       string
	ExcludePaths []string
}

type JWTOptFunc func(*JWTConfig)

func WithJWTExcludePaths(paths ...string) JWTOptFunc {
	return func(c *JWTConfig) {
		c.ExcludePaths = append(c.ExcludePaths, paths...)
	}
}

func JWTMiddleware(secret string, opts ...JWTOptFunc) fiber.Handler {
	config := JWTConfig{Secret: secret}
	for _, opt := range opts {
		opt(&config)
	}

	return func(c *fiber.Ctx) error {

		for _, path := range config.ExcludePaths {
			if c.Path() == path {
				return c.Next()
			}
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		id, okId := claims["id"].(string)
		username, okUsername := claims["username"].(string)
		email, okEmail := claims["email"].(string)
		role, okRole := claims["role"].(string)
		if !okId || !okUsername || !okEmail || !okRole {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims data"})
		}

		userData := UserData{
			ID:       id,
			Username: username,
			Email:    email,
			Role:     role,
		}

		c.Locals("user", userData)

		return c.Next()
	}
}
