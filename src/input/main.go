package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"utils"
)

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
	showNameBuilder  *ShowNameBuilder
	dirPathBuilder   *DirPathBuilder
	languageBuilder  *LanguageBuilder
}

// NewInputReader return a new Reader struct
func NewInputReader(readBuffer io.Reader, colors utils.ColorsStruct, config utils.Configuration, useDefaultValues bool) *Reader {
	construct := new(Reader)
	construct.reader = bufio.NewReader(readBuffer)
	construct.colors = colors
	construct.config = config
	construct.useDefaultValues = useDefaultValues
	construct.showNameBuilder = NewShowNameBuilder()
	construct.dirPathBuilder = NewDirPathBuilder(config.DirPathsConfig, config.SubtitleExtension)
	construct.languageBuilder = NewLanguageBuilder(config.LanguagesConfig)

	return construct
}

// ReadInputArgs Prompt all question to the user, then wrap them to a structure that will be used later
func (i *Reader) ReadInputArgs() SubtitleToDownload {

	showName, err := i.readShowName()
	var dirPathDigit = defaultDirPathValue
	var languageDigit = defaultLanguageValue

	if err != nil {
		i.colors.Red.Printf("%s\n", err)
		os.Exit(1)
	}

	if i.useDefaultValues != true {
		dirPathDigit, err = i.readDirPath()
		if err != nil {
			i.colors.Red.Printf("%s\n", err)
			os.Exit(1)
		}

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

	for key, value := range i.dirPathBuilder.GetSortedMapping() {
		i.colors.White.Printf("[%d] - %s\n", key, value)
	}

	dirPathInput, err := i.reader.ReadString('\n')
	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read DirPath failed - '%s'", err)}
	}

	dirPathInput = convertCRLFtoLF(dirPathInput)
	if len(dirPathInput) == 0 {
		return defaultDirPathValue, nil
	}

	dirPathDigit, _ := strconv.Atoi(dirPathInput)
	return dirPathDigit, nil
}

func (i *Reader) readLanguage() (int, error) {
	i.colors.Green.Printf("Indicate the subtitles' Language\n")
	for key, value := range i.languageBuilder.GetSortedMapping()  {
		i.colors.White.Printf("[%d] - %s\n", key, value)
	}

	languageInput, err := i.reader.ReadString('\n')

	if err != nil {
		return -1, &utils.Error{fmt.Sprintf("Read Language failed - '%s'", err)}
	}

	languageInput = convertCRLFtoLF(languageInput)
	if len(languageInput) == 0 {
		return defaultLanguageValue, nil
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

	showNameStruct, showNameError := i.showNameBuilder.build(showName)
	var warnings utils.Warnings

	if showNameError != nil {
		return SubtitleToDownload{}, &utils.Error{showNameError.Error()}
	}

	dirPathStruct, dirPathError := i.dirPathBuilder.build(dirPathDigit, showNameStruct)
	if dirPathError != nil {
		warnings = append(warnings, utils.Warning{dirPathError.Error()})
	}

	languageString, languageError := i.languageBuilder.build(languageDigit)
	if languageError != nil {
		warnings = append(warnings, utils.Warning{languageError.Error()})
	}

	return SubtitleToDownload{
		ShowName: showNameStruct,
		DirPath:  dirPathStruct,
		Language: languageString,
	}, warnings
}

func convertAndSortMapping(mapping map[string]string) map[int]string {
	sortedMapping := make(map[int]string)

	keys := make([]string, 0, len(mapping))
	for k := range mapping {
		keys = append(keys, k)
	}
	i := 0
	sort.Strings(keys)
	for _, element := range keys {
		i += 1
		sortedMapping[i] = mapping[element]
	}
	return sortedMapping
}