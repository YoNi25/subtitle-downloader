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
	utils.Colors.Blue.Printf("\nLooking For subtitles for '%s'\n", searchSubtitle.name)
	utils.Colors.Blue.Printf("The process could take some time\n\n")

	var warnings utils.Warnings

	c := addic7ed.New()
	showName, subtitle, err := c.SearchBest(searchSubtitle.name, searchSubtitle.language)
	if err != nil {
		return addic7ed.Subtitle{}, &utils.Error{fmt.Sprintf("Fail to retrieve subtitles - %s", err)}
	}

	utils.Colors.Blue.Printf("ℹ️  TV Show %s found !\n", searchSubtitle.name)
	utils.Colors.Blue.Println("------------------")
	utils.Colors.Blue.Println(showName)          // Output: Shameless (US) - 08x11 - A Gallagher Pedicure
	utils.Colors.Blue.Println(subtitle)          // Output: the best suitable subtitle in English language
	utils.Colors.Blue.Println(subtitle.Version)  // Output: BATV
	utils.Colors.Blue.Println(subtitle.Language) // Output: English
	utils.Colors.Blue.Println("------------------")

	if !strings.Contains(strings.ToUpper(subtitle.Version), strings.ToUpper(searchSubtitle.version)) {
		warnings = append(warnings, utils.Warning{fmt.Sprintf("The subtitle versions seem to be different. Check compatibility : %s", subtitle.Link)})
	}
	return subtitle, warnings
}
