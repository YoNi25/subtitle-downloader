package input

import (
	"errors"
	"fmt"
)

const French = 1
const English = 2

func buildLanguage(languageDigit int, defaultLanguage string) (string, error) {

	var language string
	var error error = nil

	switch languageDigit {
	case French:
		language = "French"
		break
	case English:
		language = "English"
		break
	default:
		error = errors.New(fmt.Sprintf("Unable to find Language %d. Using Default Language - '%s'", languageDigit, defaultLanguage))
		language = defaultLanguage
	}

	return language, error
}
