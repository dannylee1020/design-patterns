/*
	Singleton Pattern ensures a class has only one instance, and provides
	a global point of access to it.
*/

package singleton

import (
	"fmt"
)

type Message struct {
	msg string
}

func (m *Message) New(msg string) *Message {
	return &Message{msg: msg}
}

func (m *Message) Send() {
	fmt.Printf("Sending message...: %s", m.msg)
}

func main() {
	var message *Message

	newMsg := message.New("Hello World!")
	newMsg.Send()
}
