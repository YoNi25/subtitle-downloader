package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BuildLanguage(t *testing.T) {
	flagtests := []struct {
		languageDigit    int
		expectedLanguage string
	}{
		{1, "French"},
		{2, "English"},
	}

	for _, test := range flagtests {
		language, err := buildLanguage(test.languageDigit, "French")
		assert.Equal(t, test.expectedLanguage, language)
		assert.Nil(t, err)
	}
}

func Test_BuildLanguageWithDefaultValue(t *testing.T) {
	language, err := buildLanguage(-1, "Zulu")
	assert.Equal(t, "Zulu", language)
	assert.Equal(t, "No Language matches with -1. Using default Language - 'Zulu'", err.Error())
}
