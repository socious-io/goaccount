package goaccount

import (
	"encoding/json"
)

// Get User profile base on access token given
func (t *SessionToken) GetUserProfile(user interface{}) (*User, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users"),
		Method:   MethodGet,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(User)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Get User profile base on access token given
func (t *SessionToken) UpdateUserProfile(user interface{}) (*User, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("users"),
		Method:   MethodPut,
		Body:     user,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(User)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}
