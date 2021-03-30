package input

import "fmt"

type DirPath struct {
	RootPath string
	Folder   string
	FullPath string
}

const ServerDirPath = 1
const DesktopDirPath = 2

func buildDirPath(dirPathDigit int, showName ShowName) DirPath {
	rootPath := configuration.ServerDirPath
	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = configuration.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = configuration.DesktopDirPath
		break
	}

	return DirPath{
		RootPath: rootPath,
		Folder:   showFolder,
		FullPath: fmt.Sprintf("%s/%s", rootPath, showFolder),
	}
}
