package i18n

const (
	// EnUS is English from US
	EnUS = "en-us"
	// PtBR is Portuguese from Brazil
	PtBR = "pt-br"
)

func validateLanguage(nameTranslation string, descTranslation string) (string, string) {
	name, desc := EnUS, EnUS

	switch descTranslation {
	case EnUS:
	case PtBR:
		desc = PtBR
	}

	return name, desc
}
