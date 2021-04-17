// +build all packageCommand

package command

import (
	"bytes"
	"github.com/stretchr/testify/assert"
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
}
