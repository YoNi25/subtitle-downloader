package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"utils"
)

type Input struct {
	ShowName ShowName
	Language string
	DirPath  DirPath
}

var colors utils.ColorsStruct
var configuration utils.Configuration

func ReadInputArgs() Input {

	colors = utils.Colors
	configuration = utils.Config

	reader := bufio.NewReader(os.Stdin)

	showName := readShowName(reader)
	dirPathDigit := readDirPath(reader)
	language := readLanguage(reader)

	input := buildInput(showName, dirPathDigit, language)

	confirmInput(reader, input)

	return input
}

func readShowName(reader *bufio.Reader) string {
	colors.Green.Println("Indicate the show's episode name")
	showNameInput, err := reader.ReadString('\n')

	if err != nil {
		colors.Red.Printf("Read showName failed - '%s'\n", err)
		os.Exit(1)
	}

	return convertCRLFtoLF(showNameInput)
}

func readDirPath(reader *bufio.Reader) int {

	colors.Green.Println("Indicate the directory path where the file should be download")
	colors.White.Printf("[%d] - %s\n", ServerDirPath, configuration.ServerDirPath)
	colors.White.Printf("[%d] - %s\n", DesktopDirPath, configuration.DesktopDirPath)
	dirPathInput, err := reader.ReadString('\n')
	if err != nil {
		colors.Red.Printf("Read DirPath failed - '%s'\n", err)
		os.Exit(1)
	}

	if len(dirPathInput) == 0 {
		dirPathInput = strconv.Itoa(ServerDirPath)
	}

	dirPathDigit, _ := strconv.Atoi(convertCRLFtoLF(dirPathInput))
	return dirPathDigit
}

func readLanguage(reader *bufio.Reader) string {
	colors.Green.Printf("Indicate the subtitles' Language (Default : %s)\n", configuration.DefaultLanguage)

	languageInput, err := reader.ReadString('\n')

	if err != nil {
		colors.Red.Printf("Read Language failed - '%s'\n", err)
		os.Exit(1)
	}

	language := convertCRLFtoLF(languageInput)

	if len(language) > 0 {
		return language
	}
	return configuration.DefaultLanguage
}

func confirmInput(reader *bufio.Reader, input Input) {

	colors.Green.Printf("Download %s.%s to %s. Confirm that choice ? [Yn]\n", input.ShowName.Fullname, configuration.SubtitleExtension, input.DirPath.FullPath)
	confirm, err := reader.ReadString('\n')

	if err != nil {
		colors.Red.Printf("Read confirmation failed - '%s'\n", err)
		os.Exit(1)
	}

	confirm = convertCRLFtoLF(confirm)

	if !(len(confirm) == 0 || strings.ToUpper(confirm) != "Y") {
		colors.Red.Printf("Confirmation failed ! Invalid answer '%s'\n", confirm)
		os.Exit(2)
	}
}

func convertCRLFtoLF(toConvert string) string {
	return strings.Replace(toConvert, "\n", "", -1)
}

func buildInput(showName string, dirPathDigit int, language string) Input {

	showNameStruct := buildShowName(showName)
	dirPathStruct := buildDirPath(dirPathDigit, showNameStruct)

	return Input{
		ShowName: showNameStruct,
		DirPath:  dirPathStruct,
		Language: language,
	}
}