package services

type MessagingService interface {
	SendMessage(message string) error
}
