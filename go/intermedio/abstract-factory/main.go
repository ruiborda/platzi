package main

import "fmt"

const (
	SMS uint8 = iota
	EMAIL
)

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Printf("Sending SMS notification\n")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "AWS SNS"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Printf("Sending Email notification\n")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Mail Chimp"
}

func GetNotificationFactory(notification uint8) (INotificationFactory, error) {
	switch notification {
	case SMS:
		return SMSNotification{}, nil
	case EMAIL:
		return EmailNotification{}, nil
	}
	return nil, fmt.Errorf("invalid notification type")
}

func SendNotification(notification INotificationFactory) {
	notification.SendNotification()
}

func PrintSenderMethod(notification INotificationFactory) {
	fmt.Printf("Sender method: %s\n", notification.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := GetNotificationFactory(SMS)
	emailFactory, _ := GetNotificationFactory(EMAIL)

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	PrintSenderMethod(smsFactory)
	PrintSenderMethod(emailFactory)
}
