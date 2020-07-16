package main

import (
	"go-msgcommitter/committer"
	"go-msgcommitter/message"
	"os"
)

func main() {
	inputMessage := committer.AskForInput()
	builtMessage := message.Build(*inputMessage)
	filePath := os.Args[1]
	committer.WriteMessageToFile(filePath, builtMessage)
}
