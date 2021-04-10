package downloader

import (
	"github.com/matcornic/addic7ed"
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func Test_createMissingDirectoryAndDownloadSubtitles(t *testing.T) {
	sut := NewSubtitleDownloader(utils.Colors)
	subtitles := addic7ed.Subtitle{
		"English",
		"KOGi+Telly",
		"http://www.addic7ed.com/original/164736/1",
	}

	subtitleToDownload := subtitleToDownload{
		subtitle:  subtitles,
		name:      "The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re]",
		extension: "srt",
		dirPath:   DownloadDirPath,
	}

	expectedFilePath := "../../downloaded/The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re].srt"
	expectedErrors := utils.Warnings{
		{"Missing directory ../../downloaded. Creating ..."},
	}

	err := sut.downloadShowsSubtitles(subtitleToDownload)
	assert.FileExists(t, expectedFilePath)
	assert.Equal(t, expectedErrors, err)
}

func Test_DownloadShowSubtitles(t *testing.T) {
	sut := NewSubtitleDownloader(utils.Colors)

	subtitles := addic7ed.Subtitle{
		"English",
		"AMZN.WEB-DL.NTb",
		"http://www.addic7ed.com/original/162669/6",
	}

	subtitleToDownload := subtitleToDownload{
		subtitle:  subtitles,
		name:      "Vikings.S06E11.King.Of.Kings.720p.AMZN.WEBRip.DDP5.1.x264-NTb[eztv.re]",
		extension: "srt",
		dirPath:   DownloadDirPath,
	}

	expectedFilePath := "../../downloaded/Vikings.S06E11.King.Of.Kings.720p.AMZN.WEBRip.DDP5.1.x264-NTb[eztv.re].srt"

	err := sut.downloadShowsSubtitles(subtitleToDownload)
	assert.FileExists(t, expectedFilePath)
	assert.Nil(t, err)
}

func Test_failToDownloadShowSubtitles(t *testing.T) {
	sut := NewSubtitleDownloader(utils.Colors)

	subtitles := addic7ed.Subtitle{
		"English",
		"AMZN.WEB-DL.NTb",
		"www.addic7ed.com/original/-1/-1",
	}

	subtitleToDownload := subtitleToDownload{
		subtitle:  subtitles,
		name:      "Vikings.S06E11.King.Of.Kings.720p.AMZN.WEBRip.DDP5.1.x264-NTb[eztv.re]",
		extension: "srt",
		dirPath:   DownloadDirPath,
	}

	expectedError := &utils.Error{"Fail to download subtitles - Unable to reach addic7ed server: Get \"www.addic7ed.com/original/-1/-1\": unsupported protocol scheme \"\""}
	err := sut.downloadShowsSubtitles(subtitleToDownload)
	assert.Equal(t, expectedError, err)
}
