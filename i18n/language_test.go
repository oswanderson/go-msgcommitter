package i18n

import (
	"testing"

	"github.com/franela/goblin"
)

func TestLanguages(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Tests for validateLanguage function", func() {
		g.It("should return language en-us for name and description", func() {
			name, desc := validateLanguage(EnUS, EnUS)

			g.Assert(name).Equal(EnUS)
			g.Assert(desc).Equal(EnUS)
		})

		g.It("should return language en-us for name and pt-br for description", func() {
			name, desc := validateLanguage(EnUS, PtBR)

			g.Assert(name).Equal(EnUS)
			g.Assert(desc).Equal(PtBR)
		})

		g.It("should return language en-us for name and description for any unmapped language", func() {
			name, desc := validateLanguage("Jp-JP", "Md-CH")

			g.Assert(name).Equal(EnUS)
			g.Assert(desc).Equal(EnUS)
		})
	})
}
