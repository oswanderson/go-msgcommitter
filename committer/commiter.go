package committer

import (
	"bufio"
	"fmt"
	"go-msgcommitter/message"
	"go-msgcommitter/msgformat"
	"os"
	"strings"
)

const (
	lineStarter = '>'
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Trim(input, " ")

	return input
}

func printLineCursor() {
	fmt.Printf("%c ", lineStarter)
}

func showSubjectTypes(subjectTypes []message.SubjectType) {
	var builder strings.Builder

	for _, subjectType := range subjectTypes {
		builder.WriteString(fmt.Sprintf("  * %s (%s)\n", subjectType.Name, subjectType.Desc))
	}

	builder.WriteRune('\n')

	fmt.Print(builder.String())
}

func askForSubjectType() string {
	subjectTypes := message.GetSubjectTypes()

	fmt.Printf("What's the subject type? Must be one of:\n\n")
	showSubjectTypes(subjectTypes)
	printLineCursor()

	input := getInput()
	input = strings.ToLower(input)

	if !message.IsValidSubjectType(input) {
		fmt.Printf("\nInvalid subject type: %s\n\n", input)
		return askForSubjectType()
	}

	return input
}

func askForSubjectScope() string {
	fmt.Println("\nWhat's the scope?")
	fmt.Println("Examples: UserController, AccountRepository, TransferGateway, etc.")
	fmt.Printf("It can be empty if it can't be defined or the context is global.\n\n")
	printLineCursor()

	input := getInput()

	return input
}

func askForSubjectText() string {
	fmt.Println("\nWrite the subject text")
	fmt.Printf("It should be a short description of the changes.\n\n")
	printLineCursor()

	input := getInput()

	return input
}

func askForSubject() *msgformat.Subject {
	subjectType := askForSubjectType()
	subjectScope := askForSubjectScope()
	subjectText := askForSubjectText()

	return &msgformat.Subject{
		Type:  subjectType,
		Scope: subjectScope,
		Text:  subjectText,
	}
}

func askForBody() string {
	fmt.Println("\nWrite the message body (it can be empty)")
	fmt.Println("  * Make use of imperative and present tense: \"change\" not \"changed\" nor \"changes\"")
	fmt.Printf("  * include motivation for the change and contrasts with previous behavior\n\n")
	printLineCursor()

	input := getInput()

	return input
}

func askForFooter() string {
	fmt.Println("\nWrite the message footer (it can bem empty)")
	fmt.Println("  When referencing issues")
	fmt.Println("    * Closed issues should be listed on a separate line in")
	fmt.Println("      the footer prefixed with \"Closes\" keyword like this:")
	fmt.Println("        Closes #234")
	fmt.Println("    * or in the case of multiple issues:")
	fmt.Printf("        Closes #123, #245, #992\n\n")
	fmt.Println("  When adding breaking changes")
	fmt.Println("    * All breaking changes have to be mentioned in footer with the")
	fmt.Println("      description of the change, justification and migration notes:")
	fmt.Println("        BREAKING CHANGE:")
	fmt.Println("\n        `port-runner` command line option has changed to `runner-port`,")
	fmt.Println("        so that it is consistent with the configuration file syntax.")
	fmt.Println("\n        To migrate your project, change all the commands, where you use")
	fmt.Printf("        `--port-runner` to `--runner-port`.\n\n")
	printLineCursor()

	input := getInput()

	return input
}

// AskForInput ...
func AskForInput() *msgformat.Message {
	messageSubject := askForSubject()
	messageBody := askForBody()
	messageFooter := askForFooter()
	return &msgformat.Message{
		Subject: *messageSubject,
		Body:    messageBody,
		Footer:  messageFooter,
	}
}
