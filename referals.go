package goaccount

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type ReferralAchievement struct {
	ID              uuid.UUID       `db:"id" json:"id"`
	ReferrerID      *uuid.UUID      `db:"referrer_id" json:"referrer_id"`
	RefereeID       uuid.UUID       `db:"referee_id" json:"referee_id"`
	AchievementType string          `db:"achievement_type" json:"achievement_type"`
	RewardAmount    float32         `db:"reward_amount" json:"reward_amount"`
	RewardClaimedAt *time.Time      `db:"reward_claimed_at" json:"reward_claimed_at"`
	CreatedAt       time.Time       `db:"created_at" json:"created_at"`
	Meta            map[string]any  `db:"-" json:"meta"`
	MetaJson        *types.JSONText `db:"meta" json:"-"`
}

func (ra ReferralAchievement) AddReferralAchievement() error {
	data, _ := json.Marshal(ra)
	body := map[string]any{}

	json.Unmarshal(data, &body)

	body["client_id"] = config.ID
	body["client_secret"] = config.Secret

	_, err := Request(RequestOptions{
		Endpoint: endpoint("referrals/achievements"),
		Method:   MethodPost,
		Body:     body,
	})
	if err != nil {
		return err
	}
	return nil
}
