package input

import (
	"fmt"
	"regexp"
	"strings"
)

//ShowName The structure that contains all Show information
type ShowName struct {
	TvShow   string
	Season   string
	Episode  string
	Version  string
	Source   string
	Fullname string
}

const showNamePattern = "(?P<tvShow>.*)\\.(?P<season>S\\d+)(?P<episode>E\\d+).*-(?P<version>[^\\[]*)(\\[(?P<source>.*)\\])?"

func buildShowName(showNameStr string) (ShowName, error) {
	showNamePattern := regexp.MustCompile(`(?i)` + showNamePattern)

	match := showNamePattern.FindStringSubmatch(showNameStr)
	if len(match) == 0 {
		return ShowName{}, fmt.Errorf("Unable to parse Show name '%s", showNameStr)
	}

	result := make(map[string]string)
	for i, name := range showNamePattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return ShowName{
		strings.Title(strings.ReplaceAll(result["tvShow"], ".", " ")),
		strings.Title(result["season"]),
		strings.Title(result["episode"]),
		strings.ToUpper(result["version"]),
		result["source"],
		showNameStr,
	}, nil
}
