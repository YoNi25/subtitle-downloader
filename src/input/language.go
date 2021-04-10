package input

import (
	"fmt"
	"utils"
)

// French const used on the prompt to display the option value
const French = 1

// English const used on the prompt to display the option value
const English = 2

// language Structure used to instanciate properties needed to build the language string
type language struct {
	config utils.Configuration
}

func NewLanguage(config utils.Configuration) *language {
	construct := new(language)
	construct.config = config

	return construct
}

func (l *language) buildLanguage(languageDigit int) (string, error) {

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
		error = fmt.Errorf("No Language matches with %d. Using default Language - '%s'", languageDigit, l.config.DefaultLanguage)
		language = l.config.DefaultLanguage
	}

	return language, error
}
