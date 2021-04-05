package input

import (
	"bufio"
	"bytes"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func init() {
	utils.Colors = utils.ColorsStruct{
		Green:  color.New(),
		White:  color.New(),
		Yellow: color.New(),
		Red:    color.New(),
		Blue:   color.New(),
	}
	utils.Config =  utils.Configuration{
		ServerDirPath:   "/server/dir/path",
		DesktopDirPath:  "/desktop/dir/path",
		DefaultLanguage: "Zulu",
	}
}

func Test_ReadShowName(t *testing.T) {
	expectedShowName := "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME"
	var stdin bytes.Buffer
	stdin.Write([]byte("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME\n"))
	reader := bufio.NewReader(&stdin)

	showName, err := readShowName(reader)

	assert.Equal(t, expectedShowName, showName)
	assert.Nil(t, err)
}

func Test_ReadDirPath(t *testing.T) {
	flagtests := []struct {
		inputDirPath         string
		expectedDirPathDigit int
	}{
		{"1\n", 1},
		{"\n", -1},
	}
	var stdin bytes.Buffer

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputDirPath))
		reader := bufio.NewReader(&stdin)
		languageDigit, err := readLanguage(reader)
		assert.Equal(t, test.expectedDirPathDigit, languageDigit)
		assert.Nil(t, err)
	}
}

func Test_ReadLanguage(t *testing.T) {
	flagtests := []struct {
		inputLanguage         string
		expectedLanguageDigit int
	}{
		{"1\n", 1},
		{"\n", -1},
	}
	var stdin bytes.Buffer

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputLanguage))
		reader := bufio.NewReader(&stdin)
		languageDigit, err := readLanguage(reader)
		assert.Equal(t, test.expectedLanguageDigit, languageDigit)
		assert.Nil(t, err)
	}
}

func Test_ReadConfirmation(t *testing.T) {
	flagtests := []struct {
		inputConfirmation string
		expectedError     error
	}{
		{"\n", nil},
		{"y\n", nil},
		{"Y\n", nil},
		{"n\n", &utils.Error{"Confirmation failed ! Invalid answer 'n'"}},
	}
	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath: "/server/dir/path",
		Folder:   "Age Of Samurai Battle For Japan/S01",
		FullPath: "/server/dir/path/Age Of Samurai Battle For Japan/S01",
	}
	input := Input{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "French",
	}

	var stdin bytes.Buffer

	for _, test := range flagtests {
		stdin.Write([]byte(test.inputConfirmation))
		reader := bufio.NewReader(&stdin)
		err := confirmInput(reader, input)
		assert.Equal(t, test.expectedError, err)
	}
}

func Test_BuildInput(t *testing.T) {
	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath: "/server/dir/path",
		Folder:   "Age Of Samurai Battle For Japan/S01",
		FullPath: "/server/dir/path/Age Of Samurai Battle For Japan/S01",
	}
	expectedInput := Input{
		ShowName: showName,
		DirPath:  dirPath,
		Language: "French",
	}

	input, err := buildInput("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME", ServerDirPath, French)
	assert.Equal(t, expectedInput, input)
	assert.Nil(t, err)
}

func Test_FailBuildInput(t *testing.T) {
	_, err := buildInput("a-wrong-tv-show-name", ServerDirPath, French)
	assert.Equal(t, err, &utils.Error{"Unable to parse Show name 'a-wrong-tv-show-name"})
}

func Test_BuildInputWithWarnings(t *testing.T) {
	showName := ShowName{
		TvShow:   "Age Of Samurai Battle For Japan",
		Season:   "S01",
		Episode:  "E01",
		Version:  "EXTREME",
		Source:   "",
		Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
	}
	dirPath := DirPath{
		RootPath: "/server/dir/path",
		Folder:   "Age Of Samurai Battle For Japan/S01",
		FullPath: "/server/dir/path/Age Of Samurai Battle For Japan/S01",
	}

	flagtests := []struct {
		dirPathDigit     int
		languageDigit    int
		expectedInput    Input
		expectedWarnings error
	}{
		{
			-1,
			French,
			Input{
				ShowName: showName,
				DirPath:  dirPath,
				Language: "French",
			},
			utils.Warnings{{"No DirPath matches with -1. Using default DirPath - '/server/dir/path'"}},
		},
		{
			ServerDirPath,
			-1,
			Input{
				ShowName: showName,
				DirPath:  dirPath,
				Language: "Zulu",
			},
			utils.Warnings{{"No Language matches with -1. Using default Language - 'Zulu'"}},
		},
		{
			42,
			42,
			Input{
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
		input, err := buildInput("Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME", test.dirPathDigit, test.languageDigit)
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
