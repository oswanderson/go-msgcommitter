package message

// SubjectType ...
type SubjectType struct {
	Name string
	Desc string
}

var acceptedSubjectTypes = []SubjectType{
	{
		Name: "feat",
		Desc: "new feature for the user, not a new feature for the app/script",
	},
	{
		Name: "fix",
		Desc: "bug fix for the user, not a fix to the code",
	},
	{
		Name: "doc",
		Desc: "changes to the documentation",
	},
	{
		Name: "style",
		Desc: "formatting, missing semicolons, etc; no production code change",
	},
	{
		Name: "refactor",
		Desc: "refactoring production code, eg. renaming a variable",
	},
	{
		Name: "test",
		Desc: "adding missing tests, refactoring tests; no production code change",
	},
	{
		Name: "chore",
		Desc: "updating grunt tasks etc; no production code change",
	},
}

// GetSubjectTypes ...
func GetSubjectTypes() []SubjectType {
	subjectTypes := []SubjectType{}

	for _, subjectType := range acceptedSubjectTypes {
		subjectTypes = append(subjectTypes, subjectType)
	}

	return subjectTypes
}

// IsValidSubjectType ...
func IsValidSubjectType(value string) bool {
	for _, subjectType := range acceptedSubjectTypes {
		if value == subjectType.Name {
			return true
		}
	}

	return false
}
