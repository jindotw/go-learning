package main

import "fmt"

type MessageSender interface {
	Send(content string)
}

type EmailSender struct {
	address string
}

type SmsSender struct {
	mobile string
}

func (sender *EmailSender) Send(content string) {
	fmt.Println("Sending", content, "to", sender.address)
}

func (sender *SmsSender) Send(content string) {
	fmt.Println("Sending", content, "to", sender.mobile)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var sender MessageSender
	describe(sender)
	sender = &EmailSender{address: "testuser@gmail.com"}
	describe(sender)
	sender.Send("Hello, email")
	sender = &SmsSender{mobile: "0911-234-567"}
	describe(sender)
	sender.Send("Hello, mobile")

	var publisher Publisher
	publisher = &EmailSender{"testuser@gmail.com"}
	publisher.Connect()
	publisher.Log("hello")
	publisher.Send("world")
}
