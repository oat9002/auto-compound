package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	Ok                  int = 200
	BadRequest              = 400
	InternalServerError     = 500
)

type LineService struct {
	ApiKey string
}

func GetLineService(apikey string) *LineService {
	lineService := &LineService{ApiKey: apikey}

	return lineService
}

func (l *LineService) Send(message string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(url.Values{"message": {message}}.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+l.ApiKey)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != Ok {
		return fmt.Errorf("Sending message to line fail")
	}

	return nil
}
