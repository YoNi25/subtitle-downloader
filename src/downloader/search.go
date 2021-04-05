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

func retrieveShow(searchSubtitle searchSubtitle) (addic7ed.Subtitle, error) {
	colors.Blue.Printf("\nLooking For subtitles for '%s'\n", searchSubtitle.name)
	colors.Blue.Printf("The process could take some time\n\n")

	var warnings utils.Warnings

	c := addic7ed.New()
	showName, subtitle, err := c.SearchBest(searchSubtitle.name, searchSubtitle.language)
	if err != nil {
		return addic7ed.Subtitle{}, &utils.Error{fmt.Sprintf("Fail to retrieve subtitles - %s", err)}
	}

	colors.Blue.Printf("ℹ️  TV Show %s found !\n", searchSubtitle.name)
	colors.Blue.Println("------------------")
	colors.Blue.Println(showName)          // Output: Shameless (US) - 08x11 - A Gallagher Pedicure
	colors.Blue.Println(subtitle)          // Output: the best suitable subtitle in English language
	colors.Blue.Println(subtitle.Version)  // Output: BATV
	colors.Blue.Println(subtitle.Language) // Output: English
	colors.Blue.Println("------------------")

	if !strings.Contains(strings.ToUpper(subtitle.Version), strings.ToUpper(searchSubtitle.version)) {
		warnings = append(warnings, utils.Warning{fmt.Sprintf("The subtitle versions seem to be different. Check compatibility : %s", subtitle.Link)})
	}
	return subtitle, warnings
}
