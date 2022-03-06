package input

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const showNamePattern = "(?P<tvShow>.*)\\.(?P<season>S\\d+)(?P<episode>E\\d+).*-(?P<version>[^\\[]*)(\\[(?P<source>.*)\\])?"

//ShowName The structure that contains all Show information
type ShowName struct {
	TvShow   string
	Season   string
	Episode  string
	Version  string
	Source   string
	Fullname string
}

// Return a generated path for the TVShow
func (showName *ShowName) path() string {
	return fmt.Sprintf("%s%c%s", showName.TvShow, os.PathSeparator, showName.Season)
}

// ShowNameBuilder Structure used to instanciate properties needed to build the ShowName
type ShowNameBuilder struct {
	showNamePattern *regexp.Regexp
}

// NewShowNameBuilder return a new ShowNameBuilder structure
func NewShowNameBuilder() *ShowNameBuilder {
	construct := new(ShowNameBuilder)
	construct.showNamePattern = regexp.MustCompile(`(?i)` + showNamePattern)

	return construct
}

func (builder *ShowNameBuilder) build(showNameStr string) (ShowName, error) {

	match := builder.showNamePattern.FindStringSubmatch(showNameStr)
	if len(match) == 0 {
		return ShowName{}, fmt.Errorf("unable to parse Show name '%s'", showNameStr)
	}

	result := builder.mapRegexpToNamedVariables(match)

	tvShowName := strings.Title(strings.ReplaceAll(result["tvShow"], ".", " "))
	season := strings.Title(result["season"])
	return ShowName{
		tvShowName,
		season,
		strings.Title(result["episode"]),
		strings.ToUpper(result["version"]),
		result["source"],
		showNameStr,
	}, nil
}

func (builder *ShowNameBuilder) mapRegexpToNamedVariables(match []string) map[string]string {
	result := make(map[string]string)
	for i, name := range builder.showNamePattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}
