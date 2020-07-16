package main

import (
	"fmt"
	"go-msgcommitter/committer"
	"go-msgcommitter/message"
)

func main() {
	inputMessage := committer.AskForInput()
	message.Build(*inputMessage)
	fmt.Printf("Message: %v\n", inputMessage)
}
