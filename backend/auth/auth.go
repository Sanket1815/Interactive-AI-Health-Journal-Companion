package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey []byte

type Claims struct {
	UserID int `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func InitializeAuth(secret string) {
	JwtKey = []byte(secret)
}

func GenerateToken(userID int, email string) (string, error) {
	if len(JwtKey) == 0 {
		return "", errors.New("JWT key not initialized")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	if len(JwtKey) == 0 {
		return nil, errors.New("JWT key not initialized")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}