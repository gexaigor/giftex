package model

import "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	// UserID ...
	UserID int64
	jwt.StandardClaims
}
