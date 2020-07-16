package message

import (
	"fmt"
	"go-msgcommitter/msgformat"
	"testing"

	"github.com/franela/goblin"
)

func TestBuild(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Tests for Build function", func() {
		msgSubType := "refactor"
		msgSubScope := "Refactor"
		msgSubText := "this is a code refactoring"
		msgBody := "This is a message body."
		msgFooter := "This is a message footer."

		g.It("should return a valid generated string", func() {
			message := msgformat.Message{
				Subject: msgformat.Subject{
					Type:  msgSubType,
					Scope: msgSubScope,
					Text:  msgSubText,
				},
				Body:   msgBody,
				Footer: msgFooter,
			}

			builtMessage := Build(message)
			expectedMessage := fmt.Sprintf(
				"%s(%s): %s\n\n%s\n\n%s",
				msgSubType,
				msgSubScope,
				msgSubText,
				msgBody,
				msgFooter,
			)

			g.Assert(builtMessage).Equal(expectedMessage)
		})

		g.It("should return a valid generated string without parenthesis in case of an empty scope", func() {
			message := msgformat.Message{
				Subject: msgformat.Subject{
					Type:  msgSubType,
					Scope: "",
					Text:  msgSubText,
				},
				Body:   msgBody,
				Footer: msgFooter,
			}
			builtMessage := Build(message)
			expectedMessage := fmt.Sprintf(
				"%s: %s\n\n%s\n\n%s",
				msgSubType,
				msgSubText,
				msgBody,
				msgFooter,
			)

			g.Assert(builtMessage).Equal(expectedMessage)
		})
	})
}
