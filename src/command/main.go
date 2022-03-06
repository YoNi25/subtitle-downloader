package command

import (
	"downloader"
	"errors"
	"flag"
	"fmt"
	"input"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"utils"
)

var directory bool
var useDefaultValues bool

//SubtitleDownloader Define the main structure command
type SubtitleDownloader struct {
	reader io.Reader
	colors utils.ColorsStruct
	config utils.Configuration
}

//NewSubtitleDownloader Instantiate the Main class SubtitleDownloader
func NewSubtitleDownloader(reader io.Reader) *SubtitleDownloader {
	construct := new(SubtitleDownloader)
	construct.reader = reader
	construct.colors = utils.NewColors()
	construct.config = utils.NewConfig()

	return construct
}

//Execute Ask questions to user and download corresponding subtitle
func (sd *SubtitleDownloader) Execute() {

	flag.Parse()
	inputInstance := input.NewInputReader(sd.reader, sd.colors, sd.config, useDefaultValues)

	var toDownload input.Downloadable
	if directory == true {
		toDownload = inputInstance.BuildDirectoryToDownloadFromInputs()
	} else {
		toDownload = inputInstance.BuildSubtitleToDownloadFromInputs()
	}
	downloadSubtitles(inputInstance, sd, toDownload)
}

func downloadSubtitles(inputInstance *input.Reader, sd *SubtitleDownloader, toDownload input.Downloadable) {
	downloaderInstance := downloader.NewDownloader(sd.colors)

	if toDownload.TypeOf() == "directory" {

		file, _ := os.Stat(toDownload.Path())
		if file.IsDir() {
			files, _ := ioutil.ReadDir(toDownload.Path())
			directoryToDownload := toDownload.(input.DirectoryToDownload)
			for _, f := range files {
				// Ignore srt files
				matched, _ := filepath.Match(fmt.Sprintf("*.%s", sd.config.SubtitleExtension), filepath.Base(f.Name()))
				if !matched {
					tvShowName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
					subtitleFileName := fmt.Sprintf("%s%c%s.%s", toDownload.Path(), os.PathSeparator, tvShowName, sd.config.SubtitleExtension)
					// Check if srt already exist for the given file
					if _, err := os.Stat(subtitleFileName); errors.Is(err, os.ErrNotExist) {
						subtitleToDownload := directoryToDownload.BuildSubtitleToDownload(inputInstance, tvShowName)
						downloaderInstance.DownloadSubtitles(subtitleToDownload)
					} else {
						warning := utils.Warning{Message: fmt.Sprintf("file %s already exists", subtitleFileName)}
						sd.colors.Yellow.Printf("%s\n", warning.Error())
					}
				}
			}
		}
	} else {
		downloaderInstance.DownloadSubtitles(toDownload.(input.SubtitleToDownload))
	}
}

func init() {
	flag.BoolVar(&useDefaultValues, "fast", false, "Use default values for Language and Dir Path")
	flag.BoolVar(&directory, "directory", false, "Download all subtitles for a given directory")
}
