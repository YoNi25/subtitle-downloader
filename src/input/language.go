package input

import (
	"fmt"
)
// French const used on the prompt to display the option value
const French = 1

// English const used on the prompt to display the option value
const English = 2

func buildLanguage(languageDigit int, defaultLanguage string) (string, error) {

	var language string
	var error error

	switch languageDigit {
	case French:
		language = "French"
		break
	case English:
		language = "English"
		break
	default:
		error = fmt.Errorf("No Language matches with %d. Using default Language - '%s'", languageDigit, defaultLanguage)
		language = defaultLanguage
	}

	return language, error
}
