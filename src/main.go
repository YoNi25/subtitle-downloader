package main

import (
	"downloader"
	"input"
	"utils"
)

func main() {
	utils.Init()

	inputs := input.ReadInputArgs()

	downloader.DownloadSubtitles(inputs)
}
