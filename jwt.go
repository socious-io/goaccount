package goaccount

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("invalid token")
	} else if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, errors.New("unknown claims type, cannot proceed")
}

func ClaimsFromBearerToken(token string) (*Claims, error) {
	splited := strings.Split(token, " ")
	if len(splited) > 1 {
		token = splited[1]
	} else {
		token = splited[0]
	}
	if token == "" {
		return nil, errors.New("authorization header missing")
	}

	claims, err := VerifyToken(token)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")

		}
		return nil, errors.New("invalid token signature")
	}
	return claims, nil
}

func GenerateToken(id string, refresh bool) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.Jwt.Duration) * time.Hour)
	claims := &Claims{
		ID:      id,
		Refresh: refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Jwt.Secret))
}

func GenerateFullTokens(id string) (map[string]any, error) {
	accessToken, err := GenerateToken(id, false)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateToken(id, true)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
	}, nil
}
