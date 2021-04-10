package input

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

var stdin bytes.Buffer

func init() {
	utils.Colors = utils.ColorsStruct{
		Green:  color.New(),
		White:  color.New(),
		Yellow: color.New(),
		Red:    color.New(),
		Blue:   color.New(),
	}
	utils.Config = utils.Configuration{
		ServerDirPath:     "/server/dir/path",
		DesktopDirPath:    "/desktop/dir/path",
		DefaultLanguage:   "Zulu",
		SubtitleExtension: "srt",
	}
}

func Test_ReadShowName(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	expectedShowName := "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME"
	stdin.Write([]byte("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME\n"))

	showName, err := sut.readShowName()

	assert.Equal(t, expectedShowName, showName)
	assert.Nil(t, err)
}

func Test_ReadDirPath(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	flagtests := []struct {
		inputDirPath         string
		expectedDirPathDigit int
	}{
		{"1\n", 1},
		{"\n", -1},
	}

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputDirPath))
		languageDigit, err := sut.readLanguage()
		assert.Equal(t, test.expectedDirPathDigit, languageDigit)
		assert.Nil(t, err)
	}
}

func Test_ReadLanguage(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)
	flagtests := []struct {
		inputLanguage         string
		expectedLanguageDigit int
	}{
		{"1\n", 1},
		{"\n", -1},
	}

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputLanguage))
		languageDigit, err := sut.readLanguage()

		assert.Equal(t, test.expectedLanguageDigit, languageDigit)
		assert.Nil(t, err)
	}
}

func Test_ReadConfirmation(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	flagtests := []struct {
		inputConfirmation string
		expectedError     error
	}{
		{"\n", nil},
		{"y\n", nil},
		{"Y\n", nil},
		{"n\n", &utils.Error{"Confirmation failed ! Invalid answer 'n'"}},
	}

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputConfirmation))
		err := sut.confirmInput()

		assert.Equal(t, test.expectedError, err)
	}
}

func Test_BuildSubtitleToDownload(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath:  "/server/dir/path",
		Folder:    "Age Of Samurai Battle For Japan/S01",
		FullPath:  "/server/dir/path/Age Of Samurai Battle For Japan/S01",
		Extension: "srt",
	}
	expectedInput := SubtitleToDownload{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "French",
	}

	subtitleToDownload, err := sut.buildSubtitleToDownload("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME", ServerDirPath, French)

	assert.Equal(t, expectedInput, subtitleToDownload)
	assert.Nil(t, err)
}

func Test_FailBuildSubtitleToDownload(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)
	_, err := sut.buildSubtitleToDownload("a-wrong-tv-show-name", ServerDirPath, French)

	assert.Equal(t, err, &utils.Error{"Unable to parse Show name 'a-wrong-tv-show-name"})
}

func Test_BuildInputWithWarnings(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath:  "/server/dir/path",
		Folder:    "Age Of Samurai Battle For Japan/S01",
		FullPath:  "/server/dir/path/Age Of Samurai Battle For Japan/S01",
		Extension: "srt",
	}

	flagtests := []struct {
		dirPathDigit     int
		languageDigit    int
		expectedInput    SubtitleToDownload
		expectedWarnings error
	}{
		{
			-1,
			French,
			SubtitleToDownload{
				ShowName: showName,
				DirPath:  dirPath,
				Language: "French",
			},
			utils.Warnings{{"No DirPath matches with -1. Using default DirPath - '/server/dir/path'"}},
		},
		{
			ServerDirPath,
			-1,
			SubtitleToDownload{
				ShowName: showName,
				DirPath:  dirPath,
				Language: "Zulu",
			},
			utils.Warnings{{"No Language matches with -1. Using default Language - 'Zulu'"}},
		},
		{
			42,
			42,
			SubtitleToDownload{
				ShowName: showName,
				DirPath:  dirPath,
				Language: "Zulu",
			},
			utils.Warnings{
				{"No DirPath matches with 42. Using default DirPath - '/server/dir/path'"},
				{"No Language matches with 42. Using default Language - 'Zulu'"},
			},
		},
	}

	for _, test := range flagtests {
		input, err := sut.buildSubtitleToDownload("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME", test.dirPathDigit, test.languageDigit)

		assert.Equal(t, test.expectedInput, input)
		assert.Equal(t, test.expectedWarnings, err)
	}
}

func Test_convertCRLFtoLF(t *testing.T) {
	flagtests := []struct {
		toConvert      string
		expectedString string
	}{
		{"\n", ""},
		{"hello\n", "hello"},
		{"Hi", "Hi"},
	}
	for _, test := range flagtests {
		convertedString := convertCRLFtoLF(test.toConvert)

		assert.Equal(t, test.expectedString, convertedString)
	}
}

func Test_readInputs(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath:  "/server/dir/path",
		Folder:    "Age Of Samurai Battle For Japan/S01",
		FullPath:  "/server/dir/path/Age Of Samurai Battle For Japan/S01",
		Extension: "srt",
	}
	expectedInput := SubtitleToDownload{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "French",
	}

	stdin.Write([]byte("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME\n"))
	stdin.Write([]byte("1\n"))
	stdin.Write([]byte("1\n"))
	stdin.Write([]byte("y\n"))

	subtitleToDownload := sut.ReadInputArgs()

	assert.Equal(t, expectedInput, subtitleToDownload)
}
func Test_readInputsWithEmptyValues(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, false)

	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath:  "/server/dir/path",
		Folder:    "Age Of Samurai Battle For Japan/S01",
		FullPath:  "/server/dir/path/Age Of Samurai Battle For Japan/S01",
		Extension: "srt",
	}
	expectedInput := SubtitleToDownload{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "Zulu",
	}

	stdin.Write([]byte("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME\n"))
	stdin.Write([]byte("\n"))
	stdin.Write([]byte("\n"))
	stdin.Write([]byte("\n"))

	subtitleToDownload := sut.ReadInputArgs()

	assert.Equal(t, expectedInput, subtitleToDownload)
}

func Test_readInputsWithUsingDefaultValues(t *testing.T) {
	sut := NewInputReader(&stdin, utils.Colors, utils.Config, true)

	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath:  "/server/dir/path",
		Folder:    "Age Of Samurai Battle For Japan/S01",
		FullPath:  "/server/dir/path/Age Of Samurai Battle For Japan/S01",
		Extension: "srt",
	}
	expectedInput := SubtitleToDownload{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "Zulu",
	}

	stdin.Write([]byte("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME\n"))

	subtitleToDownload := sut.ReadInputArgs()

	assert.Equal(t, expectedInput, subtitleToDownload)
}
