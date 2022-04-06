package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TelegramService struct {
	httpClient *http.Client
	botToken   string
	chatId     string
}

func NewTelegramService(httpClient *http.Client, botToken string, chatId string) *TelegramService {
	return &TelegramService{
		httpClient: httpClient,
		botToken:   botToken,
		chatId:     chatId,
	}
}

func (t *TelegramService) SendMessage(message string) error {
	if t.botToken == "" || t.chatId == "" {
		log.Println("no bot token or chat id, not sending message")

		return nil
	}

	baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)
	req, err := http.NewRequest("GET", baseUrl, strings.NewReader(url.Values{"message": {message}, "chatId": {t.chatId}}.Encode()))
	if err != nil {
		return fmt.Errorf("cannot create send telegram message, %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot create send telegram message, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot create send telegram message, %s", resp.Status)

	}

	return nil
}
