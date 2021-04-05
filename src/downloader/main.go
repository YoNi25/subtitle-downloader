package downloader

import (
	"input"
	"os"
	"reflect"
	"utils"
)

//DownloadSubtitles Use addic7ed client to search and download the best subtitle for the given Input
func DownloadSubtitles(input input.Input) {

	subtitle, retrieveShowErr := retrieveShow(
		searchSubtitle{
			name:     input.ShowName.Fullname,
			language: input.Language,
			version:  input.ShowName.Version,
		})

	if retrieveShowErr != nil {
		typeOf := reflect.TypeOf(retrieveShowErr)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			utils.Colors.Red.Printf("%s\n", retrieveShowErr)
			os.Exit(4)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			utils.Colors.Yellow.Printf("%s\n", retrieveShowErr)
		}
	}

	subtitleToDownload := subtitleToDownload{
		subtitle:  subtitle,
		name:      input.ShowName.Fullname,
		extension: utils.Config.SubtitleExtension,
		dirPath:   input.DirPath.FullPath,
	}

	downloadErr := downloadShowsSubtitles(subtitleToDownload)

	if downloadErr != nil {
		typeOf := reflect.TypeOf(downloadErr)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			utils.Colors.Red.Printf("%s\n", downloadErr)
			os.Exit(5)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			utils.Colors.Yellow.Printf("%s\n", downloadErr)
		}
	}

	utils.Colors.Green.Printf("🎉 Subtitle %s/%s.%s downloaded\n", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension)
}
