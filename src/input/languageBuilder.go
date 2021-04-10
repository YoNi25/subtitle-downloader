package input

import (
	"fmt"
	"utils"
)

// French const used on the prompt to display the option value
const French = 1

// English const used on the prompt to display the option value
const English = 2

// LanguageBuilder Structure used to instanciate properties needed to build the LanguageBuilder string
type LanguageBuilder struct {
	config utils.Configuration
}

// NewLanguageBuilder return a new LanguageBuilder structure
func NewLanguageBuilder(config utils.Configuration) *LanguageBuilder {
	construct := new(LanguageBuilder)
	construct.config = config

	return construct
}

func (builder *LanguageBuilder) build(languageDigit int) (string, error) {

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
		error = fmt.Errorf("No Language matches with %d. Using default Language - '%s'", languageDigit, builder.config.DefaultLanguage)
		language = builder.config.DefaultLanguage
	}

	return language, error
}
