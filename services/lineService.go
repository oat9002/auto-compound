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

type lineService struct {
	httpClient *http.Client
	ApiKey     string
}

func NewLineService(httpClient *http.Client, apiKey string) *lineService {
	lineService := &lineService{httpClient: httpClient, ApiKey: apiKey}

	return lineService
}

func (l *lineService) Send(message string) error {
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(url.Values{"message": {message}}.Encode()))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+l.ApiKey)

	resp, err := l.httpClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != Ok {
		return fmt.Errorf("Sending message to line fail")
	}

	return nil
}
