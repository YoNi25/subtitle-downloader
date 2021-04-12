// +build all packageDownloader searcher

package downloader

import (
	"github.com/matcornic/addic7ed"
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func Test_retrieveShow(t *testing.T) {
	sut := NewSubtitleSearcher(utils.Colors)
	searchSubtitle := searchSubtitle{
		"The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re]",
		"English",
		"KOGi",
	}

	expectedSubtitle := addic7ed.Subtitle{
		"English",
		"KOGi+Telly",
		"http://www.addic7ed.com/original/164736/1",
	}

	subtitles, err := sut.retrieveShow(searchSubtitle)
	assert.Equal(t, expectedSubtitle, subtitles)
	assert.Nil(t, err)
}

func Test_retrieveShowAndWarnWhenVersionsAreDifferents(t *testing.T) {
	sut := NewSubtitleSearcher(utils.Colors)
	searchSubtitle := searchSubtitle{
		"The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-PHOENiX[eztv.re]",
		"English",
		"PHOENiX",
	}

	expectedSubtitle := addic7ed.Subtitle{
		"English",
		"KOGi+Telly",
		"http://www.addic7ed.com/original/164736/1",
	}

	expectedErrors := utils.Warnings{
		{"The subtitle versions seem to be different. Check compatibility : http://www.addic7ed.com/original/164736/1"},
	}

	subtitles, err := sut.retrieveShow(searchSubtitle)
	assert.Equal(t, expectedSubtitle, subtitles)
	assert.Equal(t, expectedErrors, err)
}

func Test_failToRetrieveShow(t *testing.T) {
	sut := NewSubtitleSearcher(utils.Colors)
	searchSubtitle := searchSubtitle{
		"TV.Show.S02E01-PHOENix",
		"English",
		"PHOENiX",
	}

	expectedError := &utils.Error{"Fail to retrieve subtitles - show not found for filename TV.Show.S02E01-PHOENix"}

	subtitles, err := sut.retrieveShow(searchSubtitle)
	assert.Equal(t, addic7ed.Subtitle{}, subtitles)
	assert.Equal(t, expectedError, err)
}
