package input

import (
	"errors"
	"fmt"
)

type DirPath struct {
	RootPath string
	Folder   string
	FullPath string
}

const ServerDirPath = 1
const DesktopDirPath = 2

func buildDirPath(dirPathDigit int, showName ShowName) (DirPath, error) {
	var rootPath string
	var error error

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = configuration.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = configuration.DesktopDirPath
		break
	default:
		error = errors.New(fmt.Sprintf("No DirPath matches with %d. Using default DirPath - '%s'", dirPathDigit, configuration.ServerDirPath))
		rootPath = configuration.ServerDirPath
	}

	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	return DirPath{
		RootPath: rootPath,
		Folder:   showFolder,
		FullPath: fmt.Sprintf("%s/%s", rootPath, showFolder),
	}, error
}
