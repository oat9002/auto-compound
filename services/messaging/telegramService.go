package messaging

import (
	"fmt"
	"io"
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
	req, err := http.NewRequest("GET", baseUrl, strings.NewReader(url.Values{"text": {message}, "chat_id": {t.chatId}, "parse_mode": {"HTML"}}.Encode()))
	if err != nil {
		return fmt.Errorf("create telegram request failed, %s", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot send telegram message, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("cannot send telegram message, %s", string(bodyBytes))
	}

	return nil
}
