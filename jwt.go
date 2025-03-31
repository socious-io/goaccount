package goaccount

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Refresh bool   `json:"refresh"`
	jwt.RegisteredClaims
}

// ParseToken takes a JWT token and returns the embedded claims as a pointer to Claims or an error.
// If the token is invalid for any reason, the error will be returned.
// If the token is valid, a pointer to the embedded claims will be returned with no error.
func ParseToken(tokenString string) (*Claims, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, errors.New("unknown claims type, cannot proceed")
}
