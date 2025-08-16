package goaccount

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type User struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Username        string    `db:"username" json:"username"`
	Password        *string   `db:"password" json:"-"`
	PasswordExpired bool      `db:"password_expired" json:"password_expired"`

	Status StatusType `db:"status" json:"status"`

	Email     string  `db:"email" json:"email"`
	EmailText *string `db:"email_text" json:"email_text"`
	Phone     *string `db:"phone" json:"phone"`

	FirstName         *string `db:"first_name" json:"first_name"`
	LastName          *string `db:"last_name" json:"last_name"`
	Mission           *string `db:"mission" json:"mission"`
	Bio               *string `db:"bio" json:"bio"`
	DescriptionSearch *string `db:"description_search" json:"description_search"`

	City              *string `db:"city" json:"city"`
	Country           *string `db:"country" json:"country"`
	Address           *string `db:"address" json:"address"`
	GeonameId         *int64  `db:"geoname_id" json:"geoname_id"`
	MobileCountryCode *string `db:"mobile_country_code" json:"mobile_country_code"`
	ImpactPoints      *int    `db:"impact_points" json:"impact_points"`

	AvatarID   *uuid.UUID     `db:"avatar_id" json:"avatar_id"`
	Avatar     *Media         `db:"-" json:"avatar"`
	AvatarJson types.JSONText `db:"avatar" json:"-"`

	CoverID   *uuid.UUID     `db:"cover_id" json:"cover_id"`
	Cover     *Media         `db:"-" json:"cover"`
	CoverJson types.JSONText `db:"cover" json:"-"`

	ReferredBy *uuid.UUID `db:"referred_by" json:"referred_by"`

	IdentityVerifiedAt *time.Time `db:"identity_verified_at" json:"identity_verified_at"`
	EmailVerifiedAt    *time.Time `db:"email_verified_at" json:"email_verified_at"`
	PhoneVerifiedAt    *time.Time `db:"phone_verified_at" json:"phone_verified_at"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

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

// Create user from another platform (using email, firstname and lastname)
func (u *User) Create() error {
	data, _ := json.Marshal(u)
	body := map[string]any{}

	json.Unmarshal(data, &body)

	body["client_id"] = config.ID
	body["client_secret"] = config.Secret

	response, err := Request(RequestOptions{
		Endpoint: endpoint("users"),
		Method:   MethodPost,
		Body:     body,
	})
	if err != nil {
		return err
	}

	return json.Unmarshal(response, u)
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
