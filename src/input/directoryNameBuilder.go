package input

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const directoryNamePattern = "(?P<tvShow>.*)\\/(?P<season>S\\d+)"

//DirectoryName The structure that contains all Directory information
type DirectoryName struct {
	TvShow string
	Season string
}

// Return a generated path for the TVShow
func(DirectoryName *DirectoryName) path() string {
	return fmt.Sprintf("%s%c%s", DirectoryName.TvShow, os.PathSeparator, DirectoryName.Season)
}

// DirectoryNameBuilder Structure used to instantiate properties needed to build the DirectoryNameBuilder
type DirectoryNameBuilder struct {
	directoryNamePattern *regexp.Regexp
}

// NewDirectoryNameBuilder return a new DirectoryNameBuilder structure
func NewDirectoryNameBuilder() *DirectoryNameBuilder {
	construct := new(DirectoryNameBuilder)
	construct.directoryNamePattern = regexp.MustCompile(`(?i)` + directoryNamePattern)

	return construct
}

func (builder *DirectoryNameBuilder) build(directoryName string) (DirectoryName, error) {

	match := builder.directoryNamePattern.FindStringSubmatch(directoryName)
	if len(match) == 0 {
		return DirectoryName{}, fmt.Errorf("unable to parse directory name '%s'", directoryName)
	}

	result := builder.mapRegexpToNamedVariables(match)

	tvShowName := strings.ReplaceAll(result["tvShow"], ".", " ")
	season := result["season"]
	return DirectoryName{
		tvShowName,
		season,
	}, nil
}

func (builder *DirectoryNameBuilder) mapRegexpToNamedVariables(match []string) map[string]string {
	result := make(map[string]string)
	for i, name := range builder.directoryNamePattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}