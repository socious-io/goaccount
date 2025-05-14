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
	IdentityID uuid.UUID `db:"identity_id" json:"-"`
	URL        string    `db:"url" json:"url"`
	Filename   string    `db:"filename" json:"filename"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

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

	AvatarID   *uuid.UUID     `db:"avatar_id" json:"avatar_id"`
	Avatar     *Media         `db:"-" json:"avatar"`
	AvatarJson types.JSONText `db:"avatar" json:"-"`

	CoverID   *uuid.UUID     `db:"cover_id" json:"cover_id"`
	Cover     *Media         `db:"-" json:"cover"`
	CoverJson types.JSONText `db:"cover" json:"-"`

	IdentityVerifiedAt *time.Time `db:"identity_verified_at" json:"identity_verified_at"`
	EmailVerifiedAt    *time.Time `db:"email_verified_at" json:"email_verified_at"`
	PhoneVerifiedAt    *time.Time `db:"phone_verified_at" json:"phone_verified_at"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
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
