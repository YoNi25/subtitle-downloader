package input

const French = 1
const English = 2

func buildLanguage(languageDigit int) string {

	var language string

	switch languageDigit {
	case French:
		language = "French"
		break
	case English:
		language = "English"
		break
	default:
		colors.Red.Printf("Unable to find language %d. Using default Language - '%s'\n", languageDigit, configuration.DefaultLanguage)
		language = configuration.DefaultLanguage
	}

	return language
}
