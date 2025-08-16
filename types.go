package goaccount

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

// -----------------Sessions-----------------
type Session struct {
	ID          string  `json:"id" form:"id"`
	RedirectURL string  `json:"redirect_url" form:"redirect_url"`
	AccessID    string  `json:"access_id" form:"access_id"`
	Access      *string `json:"access" form:"access"`
	ExpireAt    string  `json:"expire_at" form:"expire_at"`
	VerifiedAt  *string `json:"verified_at" form:"verified_at"`
	UpdatedAt   string  `json:"updated_at" form:"updated_at"`
	CreatedAt   string  `json:"created_at" form:"created_at"`
}

type SessionToken struct {
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	TokenType    string `json:"token_type" form:"token_type"`
	Renewed      bool   `json:"renewed" form:"renewed"`
}

type AuthSessionResponse struct {
	AuthSession Session `json:"auth_session" form:"auth_session" validate:"required,min=8"`
}

// -----------------Entities-----------------

type Media struct {
	ID         uuid.UUID `db:"id" json:"id"`
	IdentityID uuid.UUID `db:"identity_id" json:"identity_id"`
	URL        string    `db:"url" json:"url"`
	Filename   string    `db:"filename" json:"filename"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

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
