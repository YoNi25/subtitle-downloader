package downloader

import (
	"input"
	"os"
	"reflect"
	"utils"
)

var colors utils.ColorsStruct

func DownloadSubtitles(input input.Input) {

	colors = utils.Colors

	subtitle, retrieveShowErr := retrieveShow(
		searchSubtitle{
			name:     input.ShowName.Fullname,
			language: input.Language,
			version:  input.ShowName.Version,
		})

	if retrieveShowErr != nil {
		typeOf := reflect.TypeOf(retrieveShowErr)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			colors.Red.Printf("%s\n", retrieveShowErr)
			os.Exit(4)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			colors.Yellow.Printf("%s\n", retrieveShowErr)
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
			colors.Red.Printf("%s\n", downloadErr)
			os.Exit(5)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			colors.Yellow.Printf("%s\n", downloadErr)
		}
	}

	colors.Green.Printf("🎉 Subtitle %s/%s.%s downloaded\n", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension)
}
