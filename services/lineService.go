package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/log"
)

const (
	Ok                  int = 200
	BadRequest          int = 400
	InternalServerError int = 500
)

type LineService struct {
	httpClient *http.Client
	ApiKey     string
}

func NewLineService(httpClient *http.Client, apiKey string) *LineService {
	lineService := &LineService{httpClient: httpClient, ApiKey: apiKey}

	return lineService
}

func (l *LineService) Send(message string) error {
	if l.ApiKey == "" {
		log.Info("No api key for line service. Not sending message.")

		return nil
	}

	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(url.Values{"message": {message}}.Encode()))

	if err != nil {
		return fmt.Errorf("cannot create http request, %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+l.ApiKey)

	resp, err := l.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("http call failed, %w", err)
	}

	if resp.StatusCode != Ok {
		return fmt.Errorf("sending message to line fail")
	}

	return nil
}
