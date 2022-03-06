package input

import (
	"reflect"
	"utils"
)

// Downloadable Interface to describe the type of items that could be download (Subtitle, Directory)
type Downloadable interface {
	Path() string
	TypeOf() string
}

// SubtitleToDownload Structure that contains structured information needed to download subtitles
type SubtitleToDownload struct {
	ShowName ShowName
	Language string
	DirPath  DirPath
	Downloadable
}

// Path Retrieve the subtitle path from SubtitleToDownload (implement Downloadable)
func (s SubtitleToDownload) Path() string {
	return s.DirPath.FullPath
}

// TypeOf Retrieve the typeOf SubtitleToDownload (implement Downloadable)
func (s SubtitleToDownload) TypeOf() string {
	return "subtitle"
}

// DirectoryToDownload Structure that contains structured information needed to download all subtitles for a directory
type DirectoryToDownload struct {
	DirectoryName DirectoryName
	Language      string
	DirPath       DirPath
	Downloadable
}

// Path Retrieve the subtitle path from DirectoryToDownload (implement Downloadable)
func (d DirectoryToDownload) Path() string {
	return d.DirPath.FullPath
}

// TypeOf Retrieve the typeOf DirectoryToDownload (implement Downloadable)
func (d DirectoryToDownload) TypeOf() string {
	return "directory"
}

//BuildSubtitleToDownload Generate a SubtitleToDownload object from a DirectoryToDownload and a file name
func (d DirectoryToDownload) BuildSubtitleToDownload(reader *Reader, name string) SubtitleToDownload {
	showName, err := reader.showNameBuilder.build(name)

	if err != nil {
		typeOf := reflect.TypeOf(err)
		if (typeOf == reflect.TypeOf(&utils.Error{})) {
			reader.colors.Red.Printf("%s\n", err)
		} else if typeOf == reflect.TypeOf(utils.Warnings{}) {
			reader.colors.Yellow.Printf("%s\n", err)
		}
	}

	return SubtitleToDownload{
		ShowName: showName,
		Language: d.Language,
		DirPath:  d.DirPath,
	}
}
