package messaging

import "github.com/oat9002/auto-compound/config"

type MessagingService interface {
	SendMessage(message string) error
}

func NewMessageService(conf config.Config, lineService *LineService, telegramService *TelegramService) MessagingService {
	var messagingService MessagingService

	switch conf.MessagingProvider {
	case config.Line:
		messagingService = lineService
	case config.Telegram:
		messagingService = telegramService
	default:
		messagingService = telegramService
	}

	return messagingService
}
