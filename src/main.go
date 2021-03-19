package main

import (
	"flag"
)

var showStr string
var language string
var dirPath string

func main() {
	initializeCommandArgs()

	subtitle := RetrieveShow(SearchSubtitle{
		name:     showStr,
		language: language,
	})

	DownloadShowsSubtitles(SubtitleToDownload{
		subtitle:  subtitle,
		name:      showStr,
		extension: "srt",
		dirPath:   dirPath,
	})
}

func initializeCommandArgs() {

	flag.StringVar(&showStr, "showStr", "", "The Show's episode name")
	flag.StringVar(&language, "language", "French", "The wanted subtitle language")
	flag.StringVar(&dirPath, "dirPath", ".", "The dir path where the file should be download")
	flag.Parse()

	if len(showStr) == 0 {
		panic("Missing parameters. Please run ./subtitle-downloader -h to know available parameters")
	}

}
