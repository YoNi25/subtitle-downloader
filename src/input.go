package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ShowName struct {
	tvShow   string
	season   string
	episode  string
	fullname string
}

type DirPath struct {
	rootPath string
	folder   string
	fullPath string
}

type Input struct {
	showName ShowName
	language string
	dirPath  DirPath
}

const ServerDirPath = 1
const DesktopDirPath = 2

var colors Colors
var configuration Configuration

func ReadInputArgs(config Configuration, c Colors) Input {

	colors = c
	configuration = config

	reader := bufio.NewReader(os.Stdin)

	showName := readShowName(reader)
	dirPathDigit := readDirPath(reader)
	language := readLanguage(reader)

	input := buildInput(showName, dirPathDigit, language)

	confirmInput(reader, input)

	return input
}

func readShowName(reader *bufio.Reader) string {
	colors.green.Println("Indicate the show's episode name")
	showNameInput, err := reader.ReadString('\n')

	if err != nil {
		colors.red.Printf("Read showName failed - '%s'\n", err)
		os.Exit(1)
	}

	return convertCRLFtoLF(showNameInput)
}

func readDirPath(reader *bufio.Reader) int {

	colors.green.Println("Indicate the directory path where the file should be download")
	colors.white.Printf("[%d] - %s\n", ServerDirPath, configuration.ServerDirPath)
	colors.white.Printf("[%d] - %s\n", DesktopDirPath, configuration.DesktopDirPath)
	dirPathInput, err := reader.ReadString('\n')
	if err != nil {
		colors.red.Printf("Read DirPath failed - '%s'\n", err)
		os.Exit(1)
	}

	dirPathDigit, _ := strconv.Atoi(convertCRLFtoLF(dirPathInput))
	return dirPathDigit
}

func readLanguage(reader *bufio.Reader) string {
	colors.green.Printf("Indicate the subtitles' Language (Default : %s)\n", configuration.DefaultLanguage)

	languageInput, err := reader.ReadString('\n')

	if err != nil {
		colors.red.Printf("Read Language failed - '%s'\n", err)
		os.Exit(1)
	}

	return convertCRLFtoLF(languageInput)
}

func confirmInput(reader *bufio.Reader, input Input) {

	colors.green.Printf("Download %s.%s to %s. Confirm that choice ? [Yn]\n", input.showName.fullname, configuration.SubtitleExtension, input.dirPath.fullPath)
	confirm, err := reader.ReadString('\n')

	if err != nil {
		colors.red.Printf("Read confirmation failed - '%s'\n", err)
		os.Exit(1)
	}

	confirm = convertCRLFtoLF(confirm)

	if !(len(confirm) == 0 || strings.ToUpper(confirm) != "Y") {
		colors.red.Printf("Confirmation failed ! Invalid answer '%s'\n", confirm)
		os.Exit(2)
	}
}

func convertCRLFtoLF(toConvert string) string {
	return strings.Replace(toConvert, "\n", "", -1)
}

func buildInput(showName string, dirPathDigit int, language string) Input {

	showNameStruct := buildShowName(showName)
	dirPathStruct := buildDirPath(dirPathDigit, showNameStruct)
	language = buildLanguage(language)

	return Input{
		showName: showNameStruct,
		dirPath:  dirPathStruct,
		language: language,
	}
}

func buildShowName(showNameStr string) ShowName {
	showNamePattern := regexp.MustCompile(`(?i)(?P<tvShow>.*)\.(?P<season>S\d+)(?P<episode>E\d+)`)

	match := showNamePattern.FindStringSubmatch(showNameStr)
	if len(match) == 0 {
		colors.red.Printf("Unable to parse '%s'", showNameStr)
		os.Exit(3)
	}

	result := make(map[string]string)
	for i, name := range showNamePattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return ShowName{
		tvShow:   strings.Title(strings.ReplaceAll(result["tvShow"], ".", " ")),
		season:   strings.Title(result["season"]),
		episode:  strings.Title(result["episode"]),
		fullname: showNameStr,
	}
}

func buildLanguage(language string) string {
	if len(language) > 0 {
		return language
	}
	return configuration.DefaultLanguage
}

func buildDirPath(dirPathDigit int, showName ShowName) DirPath {
	rootPath := configuration.ServerDirPath
	showFolder := fmt.Sprintf("%s/%s", showName.tvShow, showName.season)

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = configuration.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = configuration.DesktopDirPath
		break
	}

	return DirPath{
		rootPath: rootPath,
		folder:   showFolder,
		fullPath: fmt.Sprintf("%s/%s", rootPath, showFolder),
	}
}
