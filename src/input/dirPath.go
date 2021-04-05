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
}

// ServerDirPath const used on the prompt to display the option value
const ServerDirPath = 1

// DesktopDirPath const used on the prompt to display the option value
const DesktopDirPath = 2

func buildDirPath(dirPathDigit int, showName ShowName) (DirPath, error) {
	var rootPath string
	var error error

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = utils.Config.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = utils.Config.DesktopDirPath
		break
	default:
		error = fmt.Errorf("No DirPath matches with %d. Using default DirPath - '%s'", dirPathDigit, utils.Config.ServerDirPath)
		rootPath = utils.Config.ServerDirPath
	}

	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	return DirPath{
		RootPath: rootPath,
		Folder:   showFolder,
		FullPath: fmt.Sprintf("%s/%s", rootPath, showFolder),
	}, error
}
