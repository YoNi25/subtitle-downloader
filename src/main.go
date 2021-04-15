package main

import (
	"command"
	"os"
)

func main() {
	subtitleDownloader := command.NewSubtitleDownloader(os.Stdin)
	subtitleDownloader.Execute()
}