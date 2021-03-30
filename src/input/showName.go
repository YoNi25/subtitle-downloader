package input

import (
	"os"
	"regexp"
	"strings"
)

type ShowName struct {
	TvShow   string
	Season   string
	Episode  string
	Fullname string
}


func buildShowName(showNameStr string) ShowName {
	showNamePattern := regexp.MustCompile(`(?i)(?P<tvShow>.*)\.(?P<season>S\d+)(?P<episode>E\d+)`)

	match := showNamePattern.FindStringSubmatch(showNameStr)
	if len(match) == 0 {
		colors.Red.Printf("Unable to parse '%s'", showNameStr)
		os.Exit(3)
	}

	result := make(map[string]string)
	for i, name := range showNamePattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return ShowName{
		TvShow:   strings.Title(strings.ReplaceAll(result["tvShow"], ".", " ")),
		Season:   strings.Title(result["season"]),
		Episode:  strings.Title(result["episode"]),
		Fullname: showNameStr,
	}
}