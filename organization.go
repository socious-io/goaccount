package goaccount

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type Organization struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Shortname   string    `db:"shortname" json:"shortname"`
	Name        *string   `db:"name" json:"name"`
	Bio         *string   `db:"bio" json:"bio"`
	Description *string   `db:"description" json:"description"`
	Email       *string   `db:"email" json:"email"`
	Phone       *string   `db:"phone" json:"phone"`

	City    *string `db:"city" json:"city"`
	Country *string `db:"country" json:"country"`
	Address *string `db:"address" json:"address"`
	Website *string `db:"website" json:"website"`

	Mission *string `db:"mission" json:"mission"`
	Culture *string `db:"culture" json:"culture"`

	LogoID   *uuid.UUID     `db:"logo_id" json:"logo_id"`
	Logo     *Media         `db:"-" json:"logo"`
	LogoJson types.JSONText `db:"logo" json:"-"`

	CoverID   *uuid.UUID     `db:"cover_id" json:"cover_id"`
	Cover     *Media         `db:"-" json:"cover"`
	CoverJson types.JSONText `db:"cover" json:"-"`

	Status OrganizationStatusType `db:"status" json:"status"`

	VerifiedImpact bool `db:"verified_impact" json:"verified_impact"`
	Verified       bool `db:"verified" json:"verified"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func GetAllOrganizations() ([]Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations"),
		Method:   MethodGet,
	})
	if err != nil {
		return nil, err
	}

	result := []Organization{}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) GetMyOrganizations() ([]Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/membered"),
		Method:   MethodGet,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := []Organization{}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetOrganization(id string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", id),
		Method:   MethodGet,
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) CreateOrganization(organization Organization) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations"),
		Method:   MethodPost,
		Body:     organization,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) UpdateOrganization(id string, organization Organization) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", id),
		Method:   MethodPut,
		Body:     organization,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) DeleteOrganization(id string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", id),
		Method:   MethodDelete,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) AddMemberToOrganization(organizationId string, userId string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/members/%s", organizationId, userId),
		Method:   MethodPost,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *SessionToken) RemoveMemberFromOrganization(organizationId string, userId string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/members/%s", organizationId, userId),
		Method:   MethodDelete,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, nil
}

func VerifyOrganization(organizationId string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/verify", organizationId),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, err
}

func ChangeOrganizationStatus(organizationId string, status string) (*Organization, error) {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/status", organizationId),
		Method:   MethodPost,
		Body: map[string]any{
			"client_id":     config.ID,
			"client_secret": config.Secret,
			"status":        status,
		},
	})
	if err != nil {
		return nil, err
	}

	result := new(Organization)
	if err := json.Unmarshal(response, result); err != nil {
		return nil, err
	}
	return result, err
}
