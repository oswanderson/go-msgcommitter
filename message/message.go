package message

import (
	"fmt"
	"go-msgcommitter/msgformat"
	"strings"
)

// Build generates the message in a specific string format and returns it.
func Build(m msgformat.Message) string {
	var builder strings.Builder

	if len(m.Subject.Scope) == 0 {
		builder.WriteString(fmt.Sprintf("%s: %s\n", m.Subject.Type, m.Subject.Text))
	} else {
		builder.WriteString(
			fmt.Sprintf("%s(%s): %s\n", m.Subject.Type, m.Subject.Scope, m.Subject.Text),
		)
	}

	if len(m.Body) > 0 {
		builder.WriteString(fmt.Sprintf("\n%s\n", m.Body))
	}

	if len(m.Footer) > 0 {
		builder.WriteString(fmt.Sprintf("\n%s", m.Footer))
	}

	return builder.String()
}
