package msgformat

// Type ...
type Type struct {
	Name string
	Desc string
}

// Scope ...
type Scope struct {
	Name string
	Desc string
}

// Subject ...
type Subject struct {
	Name string
	Desc string
}

// MessageSubject ...
type MessageSubject struct {
	Type    string
	Scope   string
	Subject string
}
