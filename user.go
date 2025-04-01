package goaccount

import (
	"encoding/json"
)

// Get User profile base on access token given
func (t *SessionToken) GetUserProfile(user interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users"),
		Method:   MethodGet,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, user); err != nil {
		return err
	}
	return nil
}

// Get User profile base on access token given
func (t *SessionToken) UpdateUserProfile(user interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users"),
		Method:   MethodPut,
		Body:     user,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, user); err != nil {
		return err
	}
	return nil
}

func VerifyUser(user interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users/verify"),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, user); err != nil {
		return err
	}
	return nil
}

func ChangeUserStatus(status string, user interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users/verify"),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"status":        status,
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, user); err != nil {
		return err
	}
	return nil
}
