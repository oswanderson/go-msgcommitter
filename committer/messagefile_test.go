package committer

import (
	"bufio"
	"fmt"
	"go-msgcommitter/message"
	"go-msgcommitter/msgformat"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

func TestMessageFile(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Tests for WriteMessageToFile function", func() {
		msgSubType := "refactor"
		msgSubScope := "Refactor"
		msgSubText := "this is a code refactoring"
		msgBody := "This is a message body."
		msgFooter := "This is a message footer."

		g.It("should successfully write a built message to the passed file", func() {
			commitMessage := msgformat.Message{
				Subject: msgformat.Subject{
					Type:  msgSubType,
					Scope: msgSubScope,
					Text:  msgSubText,
				},
				Body:   msgBody,
				Footer: msgFooter,
			}
			builtMessage := message.Build(commitMessage)
			builtMessage = fmt.Sprintf("%s\n", builtMessage) // This is made just to use \n as delimiter when reading the file
			filePath := "tempfile"

			createFile(filePath)
			WriteMessageToFile(filePath, builtMessage)

			fileContent := readFileContent(filePath)

			deleteFile(filePath)

			g.Assert(builtMessage).Equal(fileContent)
		})
	})
}

func createFile(filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}

func readFileContent(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	var strbuilder strings.Builder

	for content, err := reader.ReadString('\n'); err == nil; content, err = reader.ReadString('\n') {
		strbuilder.WriteString(content)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	return strbuilder.String()
}

func deleteFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Fatal(err)
	}
}
