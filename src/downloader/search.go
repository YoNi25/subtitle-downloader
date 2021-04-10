package downloader

import (
	"fmt"
	"github.com/matcornic/addic7ed"
	"strings"
	"utils"
)

type searchSubtitle struct {
	name     string
	language string
	version  string
}

type subtitleSearcher struct {
	colors   utils.ColorsStruct
}

// NewSubtitleSearcher return a new subtitleSearcher struct
func NewSubtitleSearcher(colors utils.ColorsStruct) *subtitleSearcher {
	construct := new(subtitleSearcher)
	construct.colors = colors

	return construct
}

func (subtitleSearch *subtitleSearcher) retrieveShow(searchSubtitle searchSubtitle) (addic7ed.Subtitle, error) {
	subtitleSearch.colors.Blue.Printf("\nLooking For subtitles for '%s'\n", searchSubtitle.name)
	subtitleSearch.colors.Blue.Printf("The process could take some time\n\n")

	var warnings utils.Warnings

	c := addic7ed.New()
	showName, subtitle, err := c.SearchBest(searchSubtitle.name, searchSubtitle.language)
	if err != nil {
		return addic7ed.Subtitle{}, &utils.Error{fmt.Sprintf("Fail to retrieve subtitles - %s", err)}
	}

	subtitleSearch.colors.Blue.Printf("ℹ️  TV Show %s found !\n", searchSubtitle.name)
	subtitleSearch.colors.Blue.Println("------------------")
	subtitleSearch.colors.Blue.Println(showName)          // Output: Shameless (US) - 08x11 - A Gallagher Pedicure
	subtitleSearch.colors.Blue.Println(subtitle)          // Output: the best suitable subtitle in English language
	subtitleSearch.colors.Blue.Println(subtitle.Version)  // Output: BATV
	subtitleSearch.colors.Blue.Println(subtitle.Language) // Output: English
	subtitleSearch.colors.Blue.Println("------------------")

	if !strings.Contains(strings.ToUpper(subtitle.Version), strings.ToUpper(searchSubtitle.version)) {
		warnings = append(warnings, utils.Warning{fmt.Sprintf("The subtitle versions seem to be different. Check compatibility : %s", subtitle.Link)})
	}
	return subtitle, warnings
}
