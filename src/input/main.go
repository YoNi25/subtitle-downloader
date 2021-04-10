package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"utils"
)

const defaultDirPathValue int = -1
const defaultLanguageValue int = -1

// SubtitleToDownload Structure that contains structured information needed to download subtitles
type SubtitleToDownload struct {
	ShowName ShowName
	Language string
	DirPath  DirPath
}

// Reader Structure used to instanciate properties needed to read user's inputs
type Reader struct {
	colors           utils.ColorsStruct
	config           utils.Configuration
	reader           *bufio.Reader
	useDefaultValues bool
}

// NewInputReader return a new Reader struct
func NewInputReader(readBuffer io.Reader, colors utils.ColorsStruct, config utils.Configuration, useDefaultValues bool) *Reader {
	construct := new(Reader)
	construct.reader = bufio.NewReader(readBuffer)
	construct.colors = colors
	construct.config = config
	construct.useDefaultValues = useDefaultValues

	return construct
}

// ReadInputArgs Prompt all question to the user, then wrap them to a structure that will be used later
func (i *Reader) ReadInputArgs() SubtitleToDownload {

	showName, err := i.readShowName()
	if err != nil {
		i.colors.Red.Printf("%s\n", err)
		os.Exit(1)
	}

	var dirPathDigit int
	if i.useDefaultValues == true {
		dirPathDigit = i.useDefaultDirPathValue()
	} else {
		dirPathDigit, err = i.readDirPath()
		if err != nil {
			i.colors.Red.Printf("%s\n", err)
			os.Exit(1)
		}
	}

	var languageDigit int
	if i.useDefaultValues == true {
		languageDigit = i.useDefaultLanguageValue()
	} else {
		languageDigit, err = i.readLanguage()
		if err != nil {
			i.colors.Red.Printf("%s\n", err)
			os.Exit(1)
		}
	}

	subtitleToDownload, err := i.buildSubtitleToDownload(showName, dirPathDigit, languageDigit)
	if err != nil {
		typeOf := reflect.TypeOf(err)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			i.colors.Red.Printf("%s\n", err)
			os.Exit(3)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			i.colors.Yellow.Printf("%s\n", err)
		}
	}

	i.displaySubtitleToDownloadInformation(subtitleToDownload)
	if i.useDefaultValues != true {
		err = i.confirmInput()
		if err != nil {
			i.colors.Red.Printf("%s\n", err)
			os.Exit(2)
		}
	}

	return subtitleToDownload
}

func (i *Reader) readShowName() (string, error) {
	i.colors.Green.Println("Indicate the show's episode name")

	showNameInput, err := i.reader.ReadString('\n')

	if err != nil {
		return "", &utils.Error{fmt.Sprintf("Read showName failed - '%s'", err)}
	}

	return convertCRLFtoLF(showNameInput), nil
}

func (i *Reader) readDirPath() (int, error) {
	i.colors.Green.Println("Indicate the directory path where the file should be download")
	i.colors.White.Printf("[%d] - %s\n", ServerDirPath, i.config.ServerDirPath)
	i.colors.White.Printf("[%d] - %s\n", DesktopDirPath, i.config.DesktopDirPath)

	dirPathInput, err := i.reader.ReadString('\n')
	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read DirPath failed - '%s'", err)}
	}

	dirPathInput = convertCRLFtoLF(dirPathInput)
	if len(dirPathInput) == 0 {
		return i.useDefaultDirPathValue(), nil
	}

	dirPathDigit, _ := strconv.Atoi(dirPathInput)
	return dirPathDigit, nil
}

func (i *Reader) readLanguage() (int, error) {
	i.colors.Green.Printf("Indicate the subtitles' Language\n")
	i.colors.White.Printf("[%d] - French\n", French)
	i.colors.White.Printf("[%d] - English\n", English)

	languageInput, err := i.reader.ReadString('\n')

	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read Language failed - '%s'", err)}
	}

	languageInput = convertCRLFtoLF(languageInput)
	if len(languageInput) == 0 {
		return i.useDefaultLanguageValue(), nil
	}

	languageDigit, _ := strconv.Atoi(languageInput)
	return languageDigit, nil
}

func (i *Reader) confirmInput() error {
	i.colors.Green.Println("Confirm that choice ? [Yn]")

	confirm, err := i.reader.ReadString('\n')

	if err != nil {
		return &utils.Error{fmt.Sprintf("Read confirmation failed - '%s'", err)}
	}

	confirm = convertCRLFtoLF(confirm)

	if !(len(confirm) == 0 || strings.ToUpper(confirm) == "Y") {
		return &utils.Error{fmt.Sprintf("Confirmation failed ! Invalid answer '%s'", confirm)}
	}
	return nil
}

func (i *Reader) displaySubtitleToDownloadInformation(subtitleToDownload SubtitleToDownload) {
	i.colors.Blue.Println()
	i.colors.Blue.Println("---------------------SUMMARY---------------------")
	i.colors.Blue.Printf("Download %s.%s\n", subtitleToDownload.ShowName.Fullname, i.config.SubtitleExtension)
	i.colors.Blue.Printf("Chosen Language : %s\n", subtitleToDownload.Language)
	i.colors.Blue.Printf("Directory path : %s\n", subtitleToDownload.DirPath.FullPath)
	i.colors.Blue.Println("-------------------------------------------------")
}

func convertCRLFtoLF(toConvert string) string {
	return strings.Replace(toConvert, "\n", "", -1)
}

func (i *Reader) buildSubtitleToDownload(showName string, dirPathDigit int, languageDigit int) (SubtitleToDownload, error) {

	showNameBuilder := NewShowNameBuilder();
	dirPathBuilder := NewDirPathBuilder(i.config)
	languageBuilder := NewLanguageBuilder(i.config)


	showNameStruct, showNameError := showNameBuilder.build(showName)
	var warnings utils.Warnings

	if showNameError != nil {
		return SubtitleToDownload{}, &utils.Error{showNameError.Error()}
	}

	dirPathStruct, dirPathError := dirPathBuilder.build(dirPathDigit, showNameStruct)
	if dirPathError != nil {
		warnings = append(warnings, utils.Warning{dirPathError.Error()})
	}

	languageString, languageError := languageBuilder.build(languageDigit)
	if languageError != nil {
		warnings = append(warnings, utils.Warning{languageError.Error()})
	}

	return SubtitleToDownload{
		ShowName: showNameStruct,
		DirPath:  dirPathStruct,
		Language: languageString,
	}, warnings
}

func (i *Reader) useDefaultDirPathValue() int{
	return defaultDirPathValue;
}
func (i *Reader) useDefaultLanguageValue() int {
	return defaultLanguageValue
}