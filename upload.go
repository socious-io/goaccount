package goaccount

import (
	"encoding/json"
	"mime/multipart"
)

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
