package downloader

import (
	"github.com/matcornic/addic7ed"
	"os"
	"strings"
)

type searchSubtitle struct {
	name     string
	language string
	version  string
}

func retrieveShow(searchSubtitle searchSubtitle) addic7ed.Subtitle {
	colors.Blue.Printf("\nLooking For subtitles for '%s'\n", searchSubtitle.name)
	colors.Blue.Printf("The process could take some time\n\n")

	c := addic7ed.New()
	showName, subtitle, err := c.SearchBest(searchSubtitle.name, searchSubtitle.language)
	if err != nil {
		colors.Red.Printf("❌ Fail to retrieve subtitles - %s\n", err)
		os.Exit(4)
	}
	colors.Blue.Printf("ℹ️  TV Show %s found !\n", searchSubtitle.name)
	colors.Blue.Println("------------------")
	colors.Blue.Println(showName)          // Output: Shameless (US) - 08x11 - A Gallagher Pedicure
	colors.Blue.Println(subtitle)          // Output: the best suitable subtitle in English language
	colors.Blue.Println(subtitle.Version)  // Output: BATV
	colors.Blue.Println(subtitle.Language) // Output: English
	colors.Blue.Println("------------------")

	if !strings.Contains(subtitle.Version, searchSubtitle.version) {
		colors.Yellow.Println()
		colors.Yellow.Printf("⚠️  The subtitle versions seem to be different. Check compatibility : %s\n\n", subtitle.Link)
	}
	return subtitle
}
