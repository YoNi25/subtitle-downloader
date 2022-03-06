// +build all packageCommand

package command

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const downloadDirPath = "../../downloaded"

var stdin bytes.Buffer

func init() {
	os.Setenv("ENV", "test")
}

func setup() {
	removeDownloadDirectory()
	os.MkdirAll(downloadDirPath, os.ModePerm)
}

func teardown() {
	removeDownloadDirectory()
}

func removeDownloadDirectory() {
	err := os.RemoveAll(downloadDirPath)
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

func Test_command(t *testing.T) {
	stdin.Write([]byte("The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re]\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("y\n"))

	sut := NewSubtitleDownloader(&stdin)

	expectedFilePath := downloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re].srt"
	sut.Execute()
	assert.FileExists(t, expectedFilePath)
	teardown()
}

func Test_download_all_dir_subtitles(t *testing.T) {
	os.MkdirAll(downloadDirPath+"/The Falcon And The Winter Soldier/S01/", os.ModePerm)
	os.Create(downloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E01.720p.WEB.h264-KOGi[eztv.re].mkv")
	os.Create(downloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re].mkv")
	firstEpisodeExpectedFilePath := downloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E01.720p.WEB.h264-KOGi[eztv.re].srt"
	secondEpisodeExpectedFilePath := downloadDirPath + "/The Falcon And The Winter Soldier/S01/The.Falcon.and.The.Winter.Soldier.S01E02.720p.WEB.h264-KOGi[eztv.re].srt"

	os.Args = append(os.Args, "--directory")

	stdin.Write([]byte("The Falcon.And.The.Winter.Soldier/S01\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("2\n"))
	stdin.Write([]byte("y\n"))

	sut := NewSubtitleDownloader(&stdin)

	sut.Execute()
	files, _ := ioutil.ReadDir(downloadDirPath + "/The Falcon And The Winter Soldier/S01/")
	assert.Equal(t, 4, len(files))
	assert.FileExists(t, firstEpisodeExpectedFilePath)
	assert.FileExists(t, secondEpisodeExpectedFilePath)
	teardown()
}
