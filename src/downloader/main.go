package downloader

import (
	"input"
	"os"
	"reflect"
	"utils"
)

// Downloader Structure used to instanciate properties needed to download subtitles
type Downloader struct {
	colors utils.ColorsStruct
}

// NewDownloader return a new Downloader struct
func NewDownloader(colors utils.ColorsStruct) *Downloader {
	construct := new(Downloader)
	construct.colors = colors

	return construct
}

//DownloadSubtitles Use addic7ed client to search and download the best subtitle for the given SubtitleToDownload
func (downloader *Downloader) DownloadSubtitles(input input.SubtitleToDownload) {

	subtitleSearcher := NewSubtitleSearcher(downloader.colors)
	subtitleDownloader := NewSubtitleDownloader(downloader.colors)

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
		extension: input.DirPath.Extension,
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
