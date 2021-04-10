package downloader

import (
	"input"
	"os"
	"reflect"
	"utils"
)

// downloader Structure used to instanciate properties needed to download subtitles
type downloader struct {
	colors utils.ColorsStruct}

// NewInput return a new downloader struct
func NewDownloader(colors utils.ColorsStruct) *downloader {
	construct := new(downloader)
	construct.colors = colors

	return construct
}

//DownloadSubtitles Use addic7ed client to search and download the best subtitle for the given SubtitleToDownload
func (downloader *downloader) DownloadSubtitles(input input.SubtitleToDownload) {

	subtitleSearcher := NewSubtitleSearcher(downloader.colors)
	subtitleDownloader := NewSubtitleDownloader(downloader.colors);

	subtitle, retrieveShowErr := subtitleSearcher.retrieveShow(searchSubtitle{
		name:     input.ShowName.Fullname,
		language: input.Language,
		version:  input.ShowName.Version,
	})
	if retrieveShowErr != nil {
		typeOf := reflect.TypeOf(retrieveShowErr)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			downloader.colors.Red.Printf("%s\n", retrieveShowErr)
			os.Exit(4)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			downloader.colors.Yellow.Printf("%s\n", retrieveShowErr)
		}
	}

	subtitleToDownload := subtitleToDownload{
		subtitle:  subtitle,
		name:      input.ShowName.Fullname,
		extension: utils.Config.SubtitleExtension,
		dirPath:   input.DirPath.FullPath,
	}
	downloadErr := subtitleDownloader.downloadShowsSubtitles(subtitleToDownload)

	if downloadErr != nil {
		typeOf := reflect.TypeOf(downloadErr)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			downloader.colors.Red.Printf("%s\n", downloadErr)
			os.Exit(5)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			downloader.colors.Yellow.Printf("%s\n", downloadErr)
		}
	}

	downloader.colors.Green.Printf("🎉 Subtitle %s/%s.%s downloaded\n", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension)
}
