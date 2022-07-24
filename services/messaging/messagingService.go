package messaging

type MessagingService interface {
	SendMessage(message string) error
}
