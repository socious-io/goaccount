package goaccount

import (
	"encoding/json"
)

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
