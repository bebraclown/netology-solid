package notify

import "fmt"

type EmailSender struct{}

func (EmailSender) Send(customer string) {
	fmt.Printf("Уведомление отправлено клиенту %s\n", customer)
}

type SMSSender struct{}

func (SMSSender) Send(customer string) {
	fmt.Printf("SMS-уведомление отправлено клиенту %s\n", customer)
}
