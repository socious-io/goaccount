package goaccount

import (
	"encoding/json"
	"fmt"
	"time"
)

type Session struct {
	ID          string  `json:"id" form:"id"`
	RedirectURL string  `json:"redirect_url" form:"redirect_url"`
	AccessID    string  `json:"access_id" form:"access_id"`
	Access      *string `json:"access" form:"access"`
	ExpireAt    string  `json:"expire_at" form:"expire_at"`
	VerifiedAt  *string `json:"verified_at" form:"verified_at"`
	UpdatedAt   string  `json:"updated_at" form:"updated_at"`
	CreatedAt   string  `json:"created_at" form:"created_at"`
}

type SessionToken struct {
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	TokenType    string `json:"token_type" form:"token_type"`
	Renewed      bool   `json:"renewed" form:"renewed"`
}

type AuthSessionResponse struct {
	AuthSession Session `json:"auth_session" form:"auth_session" validate:"required,min=8"`
}

// Starts an auth session
func StartSession(redirectURL string) (*Session, string, error) {
	response, err := Request(RequestOptions{
		Endpoint: fmt.Sprintf("%s/auth/session", config.Host),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"redirect_url":  redirectURL,
		},
	})
	if err != nil {
		return nil, "", err
	}
	sessionToken := new(AuthSessionResponse)
	if err := json.Unmarshal(response, &sessionToken); err != nil {
		return nil, "", err
	}
	return &sessionToken.AuthSession, fmt.Sprintf(
		"%s/auth/session/%s?auth_mode=login",
		config.Host,
		sessionToken.AuthSession.ID,
	), nil
}

// Verify digital code and fetch sso token
func GetSessionToken(code string) (*SessionToken, error) {
	response, err := Request(RequestOptions{
		Endpoint: fmt.Sprintf("%s/auth/session/token", config.Host),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"code":          code,
		},
	})
	if err != nil {
		return nil, err
	}
	sessionToken := new(SessionToken)
	if err := json.Unmarshal(response, &sessionToken); err != nil {
		return nil, err
	}
	return sessionToken, nil
}

func NewSessionToken(accessToken, refreshToken string) (*SessionToken, error) {
	sessionToken := &SessionToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Renewed:      false,
	}
	if err := sessionToken.Refresh(); err != nil {
		return nil, err
	}

	return sessionToken, nil
}

func (t *SessionToken) Refresh() error {
	claims, err := ParseToken(t.AccessToken)
	if err != nil {
		return err
	}

	expireAt, err := claims.GetExpirationTime()
	if err != nil {
		return err
	}
	if expireAt.Before(time.Now()) {
		return nil
	}

	response, err := Request(RequestOptions{
		Endpoint: fmt.Sprintf("%s/auth/refresh", config.Host),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"refresh_token": t.RefreshToken,
		},
	})
	if err != nil {
		return nil
	}
	if err := json.Unmarshal(response, t); err != nil {
		return nil
	}
	t.Renewed = true
	return nil
}
