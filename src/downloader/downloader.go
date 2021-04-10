package downloader

import (
	"fmt"
	addic7ed "github.com/matcornic/addic7ed"
	"os"
	"utils"
)

type subtitleToDownload struct {
	subtitle  addic7ed.Subtitle
	name      string
	extension string
	dirPath   string
}

type subtitleDownloader struct {
	colors   utils.ColorsStruct
}

// NewSubtitleDownloader return a new subtitleDownloader struct
func NewSubtitleDownloader(colors utils.ColorsStruct) *subtitleDownloader {
	construct := new(subtitleDownloader)
	construct.colors = colors

	return construct
}

func (subtitleDownloader *subtitleDownloader) downloadShowsSubtitles(subtitleToDownload subtitleToDownload) error {
	var warnings utils.Warnings

	subtitleDownloader.colors.Green.Printf("📥 Download srt for %s\n", subtitleToDownload.name)

	if _, err := os.Stat(subtitleToDownload.dirPath)
		os.IsNotExist(err) {
		warnings = append(warnings, utils.Warning{fmt.Sprintf("Missing directory %s. Creating ...", subtitleToDownload.dirPath)})
		os.MkdirAll(subtitleToDownload.dirPath, 0755)
	}

	subtitle := subtitleToDownload.subtitle
	err := subtitle.DownloadTo(fmt.Sprintf("%s/%s.%s", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension))
	if err != nil {
		return &utils.Error{fmt.Sprintf("Fail to download subtitles - %s", err)}
	}

	return warnings
}
