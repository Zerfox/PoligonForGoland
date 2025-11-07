package main

import "fmt"

type communicationСhannel interface {
	sendingMessage()
}

func generalMessage(channel communicationСhannel) {
	channel.sendingMessage()
}

// sms email push-уведомления

type sms struct{}

func newSms() *sms {
	return &sms{}
}
func (s *sms) sendingMessage() {
	fmt.Println("Отправлено sms сообщение")
}

type email struct{}

func newEmail() *email {
	return &email{}
}
func (s *email) sendingMessage() {
	fmt.Println("Осуществлена email рассылка")

}

type pushMessage struct{}

func newPushMessage() *pushMessage {
	return &pushMessage{}
}
func (s *pushMessage) sendingMessage() {
	fmt.Println("Отправлено push-уведомление")

}

func main() {
	sms := newSms()
	email := newEmail()
	pushMessage := newPushMessage()

	generalMessage(sms)
	generalMessage(email)
	generalMessage(pushMessage)

}
