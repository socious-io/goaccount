package goaccount

// Responses
type AuthSessionResponse struct {
	AuthSession Session `json:"auth_session" form:"auth_session" validate:"required,min=8"`
}

// Data Types
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
}
