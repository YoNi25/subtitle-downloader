package input

import (
	"fmt"
	"sort"
	"utils"
)

const defaultLanguageValue int = -1

// LanguageBuilder Structure used to instanciate properties needed to build the LanguageBuilder string
type LanguageBuilder struct {
	defaultLanguage string
	mapping         map[int]string
}

// NewLanguageBuilder return a new LanguageBuilder structure
func NewLanguageBuilder(config utils.LanguagesConfig) *LanguageBuilder {
	construct := new(LanguageBuilder)
	construct.defaultLanguage = config.Default

	construct.mapping = convertAndSortMapping(config.Available)

	return construct
}

func (builder *LanguageBuilder) build(languageDigit int) (string, error) {

	var language string
	var error error

	language, ok := builder.mapping[languageDigit]
	if !ok {
		error = fmt.Errorf("No Language matches with %d. Using default Language - '%s'", languageDigit, builder.defaultLanguage)
		language = builder.defaultLanguage
	}

	return language, error
}

// GetSortedMapping Sort the Languages' mapping by numeric
func (builder *LanguageBuilder) GetSortedMapping()map[int]string {
	sortedMapping := make(map[int]string)

	keys := make([]int, 0, len(builder.mapping))
	for k := range builder.mapping {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, element := range keys {
		sortedMapping[element] = builder.mapping[element]
	}
	return sortedMapping
}