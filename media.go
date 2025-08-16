package goaccount

import (
	"encoding/json"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID         uuid.UUID `db:"id" json:"id"`
	IdentityID uuid.UUID `db:"identity_id" json:"identity_id"`
	URL        string    `db:"url" json:"url"`
	Filename   string    `db:"filename" json:"filename"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func UploadMedia(file multipart.File, media interface{}) error {
	response, err := RequestMultipart(RequestOptions{
		Endpoint: endpoint("media"),
		Method:   MethodPost,
		Body: map[string]any{
			"file": file,
		},
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, media); err != nil {
		return err
	}
	return nil
}

func GetMedia(mediaId string, media interface{}) error {
	response, err := Request(RequestOptions{
		Endpoint: endpoint("media/%s", mediaId),
		Method:   MethodGet,
	})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, media); err != nil {
		return err
	}
	return nil
}
