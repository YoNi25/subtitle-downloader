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
	inputs := inputInstance.ReadInputArgs()

	downloader.DownloadSubtitles(inputs)
}

func initializeCommandArgs() {

	flag.BoolVar(&useDefaultValues, "fast", false, "Use default values for Language and Dir Path")
	flag.Parse()
}