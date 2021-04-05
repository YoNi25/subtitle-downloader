package input

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

	showName, err := readShowName(reader)
	if err != nil {
		colors.Red.Printf("%s\n", err)
		os.Exit(1)
	}

	dirPathDigit, err := readDirPath(reader)
	if err != nil {
		colors.Red.Printf("%s\n", err)
		os.Exit(1)
	}

	language, err := readLanguage(reader)
	if err != nil {
		colors.Red.Printf("%s\n", err)
		os.Exit(1)
	}

	input, err := buildInput(showName, dirPathDigit, language)
	if err != nil {
		typeOf := reflect.TypeOf(err)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			colors.Red.Printf("%s\n", err)
			os.Exit(3)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			colors.Yellow.Printf("%s\n", err)
		}
	}

	err = confirmInput(reader, input)
	if err != nil {
		colors.Red.Printf("%s\n", err)
		os.Exit(2)
	}

	return input
}

func readShowName(reader *bufio.Reader) (string, error) {
	colors.Green.Println("Indicate the show's episode name")

	showNameInput, err := reader.ReadString('\n')

	if err != nil {
		return "", &utils.Error{fmt.Sprintf("Read showName failed - '%s'", err)}
	}

	return convertCRLFtoLF(showNameInput), nil
}

func readDirPath(reader *bufio.Reader) (int, error) {
	colors.Green.Println("Indicate the directory path where the file should be download")
	colors.White.Printf("[%d] - %s\n", ServerDirPath, configuration.ServerDirPath)
	colors.White.Printf("[%d] - %s\n", DesktopDirPath, configuration.DesktopDirPath)

	dirPathInput, err := reader.ReadString('\n')
	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read DirPath failed - '%s'", err)}
	}

	dirPathInput = convertCRLFtoLF(dirPathInput)
	if len(dirPathInput) == 0 {
		return -1, nil
	}

	dirPathDigit, _ := strconv.Atoi(dirPathInput)
	return dirPathDigit, nil
}

func readLanguage(reader *bufio.Reader) (int, error) {
	colors.Green.Printf("Indicate the subtitles' Language\n")
	colors.White.Printf("[%d] - French\n", French)
	colors.White.Printf("[%d] - English\n", English)

	languageInput, err := reader.ReadString('\n')

	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read Language failed - '%s'", err)}
	}

	languageInput = convertCRLFtoLF(languageInput)
	if len(languageInput) == 0 {
		return -1, nil
	}

	languageDigit, _ := strconv.Atoi(languageInput)
	return languageDigit, nil
}

func confirmInput(reader *bufio.Reader, input Input) error {

	colors.Blue.Println()
	colors.Blue.Println("---------------------SUMMARY---------------------")
	colors.Blue.Printf("Download %s.%s\n", input.ShowName.Fullname, configuration.SubtitleExtension)
	colors.Blue.Printf("Chosen Language : %s\n", input.Language)
	colors.Blue.Printf("Directory path : %s\n", input.DirPath.FullPath)
	colors.Blue.Println("-------------------------------------------------")
	colors.Green.Println("Confirm that choice ? [Yn]")

	confirm, err := reader.ReadString('\n')

	if err != nil {
		return &utils.Error{fmt.Sprintf("Read confirmation failed - '%s'", err)}
	}

	confirm = convertCRLFtoLF(confirm)

	if !(len(confirm) == 0 || strings.ToUpper(confirm) == "Y") {
		return &utils.Error{fmt.Sprintf("Confirmation failed ! Invalid answer '%s'", confirm)}
	}
	return nil
}

func convertCRLFtoLF(toConvert string) string {
	return strings.Replace(toConvert, "\n", "", -1)
}

func buildInput(showName string, dirPathDigit int, languageDigit int) (Input, error) {
	showNameStruct, showNameError := buildShowName(showName)
	var warnings utils.Warnings

	if showNameError != nil {
		return Input{}, &utils.Error{showNameError.Error()}
	}

	dirPathStruct, dirPathError := buildDirPath(dirPathDigit, showNameStruct)
	if dirPathError != nil {
		warnings = append(warnings, utils.Warning{dirPathError.Error()})
	}

	language, languageError := buildLanguage(languageDigit, configuration.DefaultLanguage)
	if languageError != nil {
		warnings = append(warnings, utils.Warning{languageError.Error()})
	}

	return Input{
		ShowName: showNameStruct,
		DirPath:  dirPathStruct,
		Language: language,
	}, warnings
}
