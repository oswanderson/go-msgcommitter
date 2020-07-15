package message

import (
	"fmt"
	"go-msgcommitter/msgformat"
	"strings"
)

// Build generates the message in a specific string format and returns it.
func Build(m msgformat.Message) string {
	var builder strings.Builder

	subject := fmt.Sprintf("%s(%s): %s\n", m.Subject.Type, m.Subject.Scope, m.Subject.Text)
	builder.WriteString(subject)

	body := fmt.Sprintf("\n%s\n", m.Body)
	builder.WriteString(body)

	footer := fmt.Sprintf("\n%s", m.Footer)
	builder.WriteString(footer)

	return builder.String()
}
