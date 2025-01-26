package auth

import "github.com/dgrijalva/jwt-go"

// JwtKey is the secret key used to sign JWT tokens.
var JwtKey = []byte("your_secret_key")

// Claims defines the structure of the JWT claims.
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
