package goaccount

import (
	"encoding/json"
	"fmt"
)

// Starts an auth session
func StartSession(redirectURL string) (*Session, string, error) {
	response, err := Request(RequestOptions{
		Endpoint: fmt.Sprintf("%s/auth/session", config.Host),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"redirect_url":  redirectURL, //NOTE: if needs redirection within backend fmt.Sprintf("%s/auth/login/callback", config.Config.Host),
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

// Get User profile base on access token given
func GetUserProfile(token SessionToken, user interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: fmt.Sprintf("%s/users", config.Host),
		Method:   MethodGet,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &user); err != nil {
		return err
	}
	return nil
}
