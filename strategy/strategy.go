package main

import (
	"fmt"
)

type Messages interface {
	Send()
}

type Email struct{}

type SMS struct{}

func (e *Email) Send(msg string) {
	fmt.Printf("sending email...: %s", msg)
}

func (s *SMS) Send(msg string) {
	fmt.Printf("sending text...: %s", msg)
}

func main() {
	var email *Email
	var sms *SMS

	email.Send("Hello World with email!")
	sms.Send("Hellow Worldw with iMessage!")
}
