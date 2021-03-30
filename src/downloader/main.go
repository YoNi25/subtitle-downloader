package downloader

import (
	"input"
	"utils"
)

var colors utils.ColorsStruct

func DownloadSubtitles(input input.Input) {

	colors = utils.Colors

	subtitle := retrieveShow(
		searchSubtitle{
			name:     input.ShowName.Fullname,
			language: input.Language,
		})

	downloadShowsSubtitles(subtitleToDownload{
		subtitle:  subtitle,
		name:      input.ShowName.Fullname,
		extension: utils.Config.SubtitleExtension,
		dirPath:   input.DirPath.FullPath,
	})
}
