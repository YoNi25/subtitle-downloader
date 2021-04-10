package input

import (
	"fmt"
	"utils"
)

// ServerDirPath const used on the prompt to display the option value
const ServerDirPath = 1

// DesktopDirPath const used on the prompt to display the option value
const DesktopDirPath = 2

//DirPath The structure that contains all information needed to download a subtitle
type DirPath struct {
	RootPath string
	Folder   string
	FullPath string
	Extension string
}

// DirPathBuilder Structure used to instanciate properties needed to build a DirPath
type DirPathBuilder struct {
	config utils.Configuration
}

// NewDirPathBuilder return a new DirPathBuilder structure
func NewDirPathBuilder(config utils.Configuration) *DirPathBuilder {
	construct := new(DirPathBuilder)
	construct.config = config

	return construct
}

func (builder *DirPathBuilder) build(dirPathDigit int, showName ShowName) (DirPath, error) {
	var rootPath string
	var error error

	switch dirPathDigit {
	case ServerDirPath:
		rootPath = builder.config.ServerDirPath
		break
	case DesktopDirPath:
		rootPath = builder.config.DesktopDirPath
		break
	default:
		error = fmt.Errorf("No DirPath matches with %d. Using default DirPath - '%s'", dirPathDigit, utils.Config.ServerDirPath)
		rootPath = builder.config.ServerDirPath
	}

	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	return DirPath{
		RootPath:  rootPath,
		Folder:    showFolder,
		FullPath:  fmt.Sprintf("%s/%s", rootPath, showFolder),
		Extension: builder.config.SubtitleExtension,
	}, error
}
