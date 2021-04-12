package command

import (
	"downloader"
	"flag"
	"input"
	"io"
	"utils"
)

var useDefaultValues bool

type subtitleDownloader struct {
	reader io.Reader
	colors utils.ColorsStruct
	config utils.Configuration
}

//NewSubtitleDownloader Instantiate the Main class SubtitleDownloader
func NewSubtitleDownloader(reader io.Reader) *subtitleDownloader {
	construct := new(subtitleDownloader)
	construct.reader = reader
	construct.colors = utils.NewColors()
	construct.config = utils.NewConfig()

	return construct
}

//Execute Ask questions to user and download corresponding subtitle
func (sd *subtitleDownloader) Execute() {

	initializeCommandArgs()

	inputInstance := input.NewInputReader(sd.reader, sd.colors, sd.config, useDefaultValues)
	downloaderInstance := downloader.NewDownloader(sd.colors)

	subtitleToDownload := inputInstance.ReadInputArgs()

	downloaderInstance.DownloadSubtitles(subtitleToDownload)
}

func initializeCommandArgs() {
	flag.BoolVar(&useDefaultValues, "fast", false, "Use default values for Language and Dir Path")
	flag.Parse()
}