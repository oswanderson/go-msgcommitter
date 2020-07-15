package msgformat

// Message holds all information that should contain in the commit message.
type Message struct {
	Subject Subject
	Body    string
	Footer  string
}
