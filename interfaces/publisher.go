package main

import (
	"fmt"
)

type Logger interface {
	Log(info string)
}

type Publisher interface {
	MessageSender
	Logger
	Connect()
}

func (sender *EmailSender) Log(content string) {
	fmt.Println("email sender is logging", content)
}

func (sender *EmailSender) Connect() {
	fmt.Println("email sender is trying to connect")
}
