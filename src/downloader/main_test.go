// +build all packageDownloader

package downloader

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"input"
	"os"
	"testing"
	"utils"
)

const DownloadDirPath = "../../downloaded"

func init() {
	utils.Colors = utils.ColorsStruct{
		Green:  color.New(),
		White:  color.New(),
		Yellow: color.New(),
		Red:    color.New(),
		Blue:   color.New(),
	}

	utils.Config = utils.Configuration{
		SubtitleExtension: "srt",
	}
}

func setup() {
	removeDownloadDirectory()
}

func teardown() {
	removeDownloadDirectory()
}

func removeDownloadDirectory() {
	err := os.RemoveAll(DownloadDirPath)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func Test_downloadSubtitles(t *testing.T) {
	sut := NewDownloader(utils.Colors)
	subtitleToDownload := input.SubtitleToDownload{
		ShowName: input.ShowName{
			TvShow:   "The Falcon And The Winter Soldier",
			Season:   "S01",
			Episode:  "E02",
			Version:  "KOGi",
			Source:   "eztv.re",
			Fullname: "The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re]",
		},
		DirPath: input.DirPath{
			RootPath: DownloadDirPath,
			Folder:   "The Falcon And The Winter Soldier/S01",
			FullPath: DownloadDirPath +"/The Falcon And The Winter Soldier/S01",
		},
		Language: "French",
	}

	expectedFilePath := DownloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re].srt"

	sut.DownloadSubtitles(subtitleToDownload);
	assert.FileExists(t, expectedFilePath)
}
