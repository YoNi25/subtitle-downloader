package downloader

import (
	"fmt"
	addic7ed "github.com/matcornic/addic7ed"
	"os"
)

type subtitleToDownload struct {
	subtitle  addic7ed.Subtitle
	name      string
	extension string
	dirPath   string
}

func downloadShowsSubtitles(subtitleToDownload subtitleToDownload) {
	colors.Green.Printf("⬇     Download srt for %s\n", subtitleToDownload.name)

	if _, err := os.Stat(subtitleToDownload.dirPath)
		os.IsNotExist(err) {
		colors.Yellow.Printf("🚧     Missing directory %s. Creating ...\n", subtitleToDownload.dirPath)
		os.MkdirAll(subtitleToDownload.dirPath, 0755)
	}

	subtitle := subtitleToDownload.subtitle
	err := subtitle.DownloadTo(fmt.Sprintf("%s/%s.%s", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension))
	if err != nil {
		colors.Red.Printf("Fail to download subtitles - %s", err)
		os.Exit(5)
	}

	colors.Green.Printf("🎉     Subtitle %s/%s.%s downloaded\n", subtitleToDownload.dirPath, subtitleToDownload.name, subtitleToDownload.extension)
}
