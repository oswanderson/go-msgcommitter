package msgformat

// Message holds all information that should contain in the commit message.
type Message struct {
	Subject MessageSubject
	Body    MessageBody
	Footer  MessageFooter
}
