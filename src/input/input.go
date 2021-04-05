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
		colors.Red.Printf("❌ Read showName failed - '%s'\n", err)
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
		colors.Red.Printf("❌ Read DirPath failed - '%s'\n", err)
		os.Exit(1)
	}

	dirPathInput = convertCRLFtoLF(dirPathInput)
	if len(dirPathInput) == 0 {
		dirPathInput = strconv.Itoa(ServerDirPath)
	}

	dirPathDigit, _ := strconv.Atoi(dirPathInput)
	return dirPathDigit
}

func readLanguage(reader *bufio.Reader) int {
	colors.Green.Printf("Indicate the subtitles' Language\n")
	colors.White.Printf("[%d] - French\n", French)
	colors.White.Printf("[%d] - English\n", English)
	languageInput, err := reader.ReadString('\n')

	if err != nil {
		colors.Red.Printf("❌ Read Language failed - '%s'\n", err)
		os.Exit(1)
	}

	languageInput = convertCRLFtoLF(languageInput)
	if len(languageInput) == 0 {
		languageInput = "-1"
	}

	languageDigit, _ := strconv.Atoi(languageInput)
	return languageDigit
}

func confirmInput(reader *bufio.Reader, input Input) {

	colors.Blue.Println()
	colors.Blue.Println("---------------------SUMMARY---------------------")
	colors.Blue.Printf("Download %s.%s\n", input.ShowName.Fullname, configuration.SubtitleExtension)
	colors.Blue.Printf("Chosen Language : %s\n", input.Language)
	colors.Blue.Printf("Directory path : %s\n", input.DirPath.FullPath)
	colors.Blue.Println("-------------------------------------------------")
	colors.Green.Println("Confirm that choice ? [Yn]")

	confirm, err := reader.ReadString('\n')

	if err != nil {
		colors.Red.Printf("❌ Read confirmation failed - '%s'\n", err)
		os.Exit(1)
	}

	confirm = convertCRLFtoLF(confirm)

	if !(len(confirm) == 0 || strings.ToUpper(confirm) == "Y") {
		colors.Red.Printf("❌ Confirmation failed ! Invalid answer '%s'\n", confirm)
		os.Exit(2)
	}
}

func convertCRLFtoLF(toConvert string) string {
	return strings.Replace(toConvert, "\n", "", -1)
}

func buildInput(showName string, dirPathDigit int, languageDigit int) Input {

	showNameStruct, err := buildShowName(showName)
	if err != nil {
		colors.Red.Printf("❌ %s\n", err)
		os.Exit(3)
	}

	dirPathStruct := buildDirPath(dirPathDigit, showNameStruct)

	language, err := buildLanguage(languageDigit, configuration.DefaultLanguage)
	if err != nil {
		colors.Red.Printf("⚠️  %s\n", err)
	}

	return Input{
		ShowName: showNameStruct,
		DirPath:  dirPathStruct,
		Language: language,
	}
}
