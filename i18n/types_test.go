package i18n

import (
	"testing"

	"github.com/franela/goblin"
)

func TestTypesTranslation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Tests for typeTranslation struct", func() {
		g.Describe("Tests for toType method", func() {
			typeTsl := typeTranslation{
				Name: map[string]string{
					EnUS: "type",
					PtBR: "tipo",
				},
				Desc: map[string]string{
					EnUS: "description",
					PtBR: "descrição",
				},
			}

			g.It("should return a valid msgformat.Type with translation en-us/en-us", func() {
				result := typeTsl.toType(EnUS, EnUS)

				g.Assert(result.Name).Equal("type")
				g.Assert(result.Desc).Equal("description")
			})

			g.It("should return a valid msgformat.Type with translation en-us/pt-br", func() {
				result := typeTsl.toType(EnUS, PtBR)

				g.Assert(result.Name).Equal("type")
				g.Assert(result.Desc).Equal("descrição")
			})
		})

		g.Describe("Tests for getName method", func() {
			typeTsl := typeTranslation{
				Name: map[string]string{EnUS: "feat"},
				Desc: map[string]string{
					EnUS: "new feature for the user, not a new feature for the app/script",
					PtBR: "novas features para o usuário, não para a aplicação/script",
				},
			}

			g.It("should return a name for language en-us", func() {
				name := typeTsl.getName(EnUS)

				g.Assert(name).Equal("feat")
			})

			g.It("should return a name for language en-us if another invalid language is used", func() {
				name := typeTsl.getName("Jp-JP")

				g.Assert(name).Equal("feat")
			})
		})

		g.Describe("Tests for getDescription method", func() {
			typeTsl := typeTranslation{
				Name: map[string]string{EnUS: "feat"},
				Desc: map[string]string{
					EnUS: "new feature for the user, not a new feature for the app/script",
					PtBR: "novas features para o usuário, não para a aplicação/script",
				},
			}

			g.It("should return a description for language pt-br", func() {
				description := typeTsl.getDescription(PtBR)

				g.Assert(description).Equal("novas features para o usuário, não para a aplicação/script")
			})

			g.It("should return a description for language en-us if another invalid language is used", func() {
				description := typeTsl.getDescription("Jp-JP")

				g.Assert(description).Equal("new feature for the user, not a new feature for the app/script")
			})
		})
	})

	g.Describe("Tests for GetTypes function", func() {
		g.It("should return an array of length 7 and no empty types", func() {
			result := GetTypes(EnUS, EnUS)

			g.Assert(len(result)).Equal(7)

			for _, subjectType := range result {
				g.Assert((subjectType != Type{})).IsTrue()
			}
		})

		g.It("should return correct translations for en-us/en-us", func() {
			result := GetTypes(EnUS, EnUS)
			counter := 0

			for _, subjectType := range result {
				switch name := subjectType.Name; name {
				case "feat":
					g.Assert(subjectType.Desc).Equal("new feature for the user, not a new feature for the app/script")
					counter++
				case "doc":
					g.Assert(subjectType.Desc).Equal("changes to the documentation")
					counter++
				case "refactor":
					g.Assert(subjectType.Desc).Equal("refactoring production code, eg. renaming a variable")
					counter++
				default:
				}
			}

			g.Assert(counter).Equal(3)
		})

		g.It("should return correct translations for en-us/pt-br", func() {
			result := GetTypes(EnUS, PtBR)
			counter := 0

			for _, subjectType := range result {
				switch name := subjectType.Name; name {
				case "feat":
					g.Assert(subjectType.Desc).Equal("novas features para o usuário, não para a aplicação/script")
					counter++
				case "style":
					g.Assert(subjectType.Desc).Equal("formatação, ausência de pronto e vírgulas, etc; nada para o código em produção")
					counter++
				case "refactor":
					g.Assert(subjectType.Desc).Equal("refatoração de código em produção, como por exemplo, renomeação de variável")
					counter++
				default:
				}
			}

			g.Assert(counter).Equal(3)
		})
	})
}
