// +build all packageInput builder language

package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func Test_BuildLanguage(t *testing.T) {
	sut := NewLanguageBuilder(utils.Config.LanguagesConfig)

	flagtests := []struct {
		languageDigit    int
		expectedLanguage string
	}{
		{english, "English"},
		{french, "French"},
	}

	for _, test := range flagtests {
		language, err := sut.build(test.languageDigit)
		assert.Equal(t, test.expectedLanguage, language)
		assert.Nil(t, err)
	}
}

func Test_BuildLanguageWithDefaultValue(t *testing.T) {
	sut := NewLanguageBuilder(utils.Config.LanguagesConfig)

	language, err := sut.build(-1)
	assert.Equal(t, "Zulu", language)
	assert.Equal(t, "no Language matches with -1. Using default Language - 'Zulu'", err.Error())
}
