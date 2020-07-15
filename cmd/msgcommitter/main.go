package main

import (
	"fmt"
	"go-msgcommitter/message"
	"go-msgcommitter/msgformat"
)

func main() {
	fmt.Println("Hello from Message Committer!!!")
	msgSubType := "refactor"
	msgSubScope := "Refactor"
	msgSubText := "this is a code refactoring"
	msgSubject := msgformat.Subject{
		Type:  msgSubType,
		Scope: msgSubScope,
		Text:  msgSubText,
	}
	msgBody := "This is a message body."
	msgFooter := "This is a message footer."
	msg := msgformat.Message{
		Subject: msgSubject,
		Body:    msgBody,
		Footer:  msgFooter,
	}
	message.Build(msg)
}
