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
