package messaging

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type LineService struct {
	httpClient *http.Client
	apiKey     string
}

func NewLineService(httpClient *http.Client, apiKey string) *LineService {
	lineService := &LineService{httpClient: httpClient, apiKey: apiKey}

	return lineService
}

func (l *LineService) SendMessage(message string) error {
	if l.apiKey == "" {
		log.Println("no api key for line service. not sending message.")

		return nil
	}

	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(url.Values{"message": {message}}.Encode()))

	if err != nil {
		return fmt.Errorf("cannot send line message, %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+l.apiKey)

	resp, err := l.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("cannot send line message, %s", err)

	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("cannot send line message, %s", string(bodyBytes))
	}

	return nil
}
