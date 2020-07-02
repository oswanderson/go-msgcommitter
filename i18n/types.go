package i18n

import "go-msgcommitter/msgformat"

type typeTranslation struct {
	Name map[string]string
	Desc map[string]string
}

func (t *typeTranslation) getName(translation string) string {
	if name, ok := t.Name[translation]; ok {
		return name
	}

	return t.Name[EnUS]
}

func (t *typeTranslation) getDescription(translation string) string {
	if description, ok := t.Desc[translation]; ok {
		return description
	}

	return t.Desc[EnUS]
}

func (t *typeTranslation) toType(nameTranslation string, descTranslation string) msgformat.Type {
	return msgformat.Type{
		Name: t.getName(nameTranslation),
		Desc: t.getDescription(descTranslation),
	}
}

var translations = [...]typeTranslation{
	{
		Name: map[string]string{EnUS: "feat"},
		Desc: map[string]string{
			EnUS: "new feature for the user, not a new feature for the app/script",
			PtBR: "novas features para o usuário, não para a aplicação/script",
		},
	},
	{
		Name: map[string]string{EnUS: "fix"},
		Desc: map[string]string{
			EnUS: "bug fix for the user, not a fix to the code",
			PtBR: "correções de bugs para o usuário, não para a aplicação/script",
		},
	},
	{
		Name: map[string]string{EnUS: "doc"},
		Desc: map[string]string{
			EnUS: "changes to the documentation",
			PtBR: "mudanças na documentação",
		},
	},
	{
		Name: map[string]string{EnUS: "style"},
		Desc: map[string]string{
			EnUS: "formatting, missing semicolons, etc; no production code change",
			PtBR: "formatação, ausência de pronto e vírgulas, etc; nada para o código em produção",
		},
	},
	{
		Name: map[string]string{EnUS: "refactor"},
		Desc: map[string]string{
			EnUS: "refactoring production code, eg. renaming a variable",
			PtBR: "refatoração de código em produção, como por exemplo, renomeação de variável",
		},
	},
	{
		Name: map[string]string{EnUS: "test"},
		Desc: map[string]string{
			EnUS: "adding missing tests, refactoring tests; no production code change",
			PtBR: "adição de testes faltantes, refatoração de testes; nada para o código em produção",
		},
	},
	{
		Name: map[string]string{EnUS: "chore"},
		Desc: map[string]string{
			EnUS: "updating grunt tasks etc; no production code change",
			PtBR: "atualização de tarefas que executam lint, testes unitários, compilação, etc; nada para o código em produção",
		},
	},
}

// GetTypes returns a list of default message subject types with infos according to the
// language defined on environment variable XXXX.
func GetTypes(nameTranslation string, descTranslation string) [7]msgformat.Type {
	types := [len(translations)]msgformat.Type{}

	for index, typeTsl := range translations {
		validatedName, validatedDesc := validateLanguage(nameTranslation, descTranslation)
		types[index] = typeTsl.toType(validatedName, validatedDesc)
	}

	return types
}
