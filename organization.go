package goaccount

import (
	"encoding/json"
)

func GetAllOrganizations(organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations"),
		Method:   MethodGet,
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) GetMyOrganizations(organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/membered"),
		Method:   MethodGet,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func GetOrganization(organizationId string, organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", organizationId),
		Method:   MethodGet,
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) CreateOrganization(organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations"),
		Method:   MethodPost,
		Body:     organization,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) UpdateOrganization(organizationId string, organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", organizationId),
		Method:   MethodPut,
		Body:     organization,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) DeleteOrganization(organizationId string, organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s", organizationId),
		Method:   MethodDelete,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) AddMemberToOrganization(organizationId string, userId string, organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/members/%s", organizationId, userId),
		Method:   MethodPost,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}

func (t *SessionToken) RemoveMemberFromOrganization(organizationId string, userId string, organization interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("organizations/%s/members/%s", organizationId, userId),
		Method:   MethodDelete,
		Headers: map[string]string{
			"Authorization": bearer(t.AccessToken),
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, organization); err != nil {
		return err
	}
	return nil
}
