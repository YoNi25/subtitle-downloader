package input

import (
	"fmt"
	"utils"
)

//DirPath The structure that contains all information needed to download a subtitle
type DirPath struct {
	RootPath string
	Folder   string
	FullPath string
	Extension string
}

// dirPath Structure used to instanciate properties needed to build a DirPath
type dirPath struct {
	config utils.Configuration
}

// ServerDirPath const used on the prompt to display the option value
const ServerDirPath = 1

// DesktopDirPath const used on the prompt to display the option value
const DesktopDirPath = 2

func NewDirPath(config utils.Configuration) *dirPath {
	construct := new(dirPath)
	construct.config = config

	return construct
}

func (dirPath *dirPath) buildDirPath(dirPathDigit int, showName ShowName) (DirPath, error) {
	var rootPath string
	var error error

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = dirPath.config.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = dirPath.config.DesktopDirPath
		break
	default:
		error = fmt.Errorf("No DirPath matches with %d. Using default DirPath - '%s'", dirPathDigit, utils.Config.ServerDirPath)
		rootPath = dirPath.config.ServerDirPath
	}

	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	return DirPath{
		RootPath: rootPath,
		Folder:   showFolder,
		FullPath: fmt.Sprintf("%s/%s", rootPath, showFolder),
		Extension: dirPath.config.SubtitleExtension,
	}, error
}
