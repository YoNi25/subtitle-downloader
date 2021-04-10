package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func Test_BuildLanguage(t *testing.T) {
	sut := NewLanguage(utils.Config)

	flagtests := []struct {
		languageDigit    int
		expectedLanguage string
	}{
		{1, "French"},
		{2, "English"},
	}

	for _, test := range flagtests {
		language, err := sut.buildLanguage(test.languageDigit)
		assert.Equal(t, test.expectedLanguage, language)
		assert.Nil(t, err)
	}
}

func Test_BuildLanguageWithDefaultValue(t *testing.T) {
	sut := NewLanguage(utils.Config)

	language, err := sut.buildLanguage(-1)
	assert.Equal(t, "Zulu", language)
	assert.Equal(t, "No Language matches with -1. Using default Language - 'Zulu'", err.Error())
}
