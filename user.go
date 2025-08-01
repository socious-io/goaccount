package goaccount

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ImpactPoint struct {
	UserID              uuid.UUID      `json:"user_id" form:"user_id" validate:"required"`
	TotalPoints         int            `json:"total_points" form:"total_points"`
	SocialCause         string         `json:"social_cause" form:"social_cause"`
	SocialCauseCategory string         `json:"social_cause_category" form:"social_cause_category"`
	Type                string         `json:"type" form:"type" validate:"required,oneof=WORKSUBMIT DONATION VOLUNTEER OTHER"`
	AccessID            *uuid.UUID     `json:"access_id" form:"access_id"`
	Meta                map[string]any `json:"meta" form:"meta"`
	UniqueTag           string         `json:"unique_tag" form:"unique_tag"`
	Value               float64        `json:"value" form:"value"`
}

// Get User profile base on access token given
func (t *SessionToken) GetUserProfile() (*User, error) {
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

func (ip ImpactPoint) AddImpactPoint() error {
	data, _ := json.Marshal(ip)
	body := map[string]any{}

	json.Unmarshal(data, &body)

	body["client_id"] = config.ID
	body["client_secret"] = config.Secret

	_, err := Request(RequestOptions{
		Endpoint: endpoint("impact-points"),
		Method:   MethodPost,
		Body:     body,
	})
	if err != nil {
		return err
	}
	return nil
}
