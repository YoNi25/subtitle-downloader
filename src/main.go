package main

import (
	"downloader"
	"flag"
	"input"
	"os"
	"utils"
)

var useDefaultValues bool

func main() {
	utils.Init()
	initializeCommandArgs()

	inputInstance := input.NewInput(utils.Colors, utils.Config, os.Stdin, useDefaultValues)
	downloaderInstance := downloader.NewDownloader(utils.Colors)

	inputs := inputInstance.ReadInputArgs()

	downloaderInstance.DownloadSubtitles(inputs)
}

func initializeCommandArgs() {

	flag.BoolVar(&useDefaultValues, "fast", false, "Use default values for Language and Dir Path")
	flag.Parse()
}