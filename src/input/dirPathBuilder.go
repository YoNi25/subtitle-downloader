package input

import (
	"fmt"
	"utils"
)

const defaultDirPathValue int = -1

//DirPath The structure that contains all information needed to download a subtitle
type DirPath struct {
	RootPath  string
	Folder    string
	FullPath  string
	Extension string
}

// DirPathBuilder Structure used to instantiate properties needed to build a DirPath
type DirPathBuilder struct {
	defaultDirPath    string
	subtitleExtension string
	mapping           map[int]string
}

// NewDirPathBuilder return a new DirPathBuilder structure
func NewDirPathBuilder(config utils.DirPathsConfig, subtitleExtension string) *DirPathBuilder {
	construct := new(DirPathBuilder)

	construct.defaultDirPath = config.Default
	construct.subtitleExtension = subtitleExtension
	construct.mapping = convertAndSortMapping(config.Available)

	return construct
}

func (builder *DirPathBuilder) build(dirPathDigit int, showName ShowName) (DirPath, error) {
	var error error

	rootPath, ok := builder.mapping[dirPathDigit]
	if !ok {
		error = fmt.Errorf("No DirPath matches with %d. Using default DirPath - '%s'", dirPathDigit, builder.defaultDirPath)
		rootPath = builder.defaultDirPath
	}

	showFolder := fmt.Sprintf("%s/%s", showName.TvShow, showName.Season)

	return DirPath{
		RootPath:  rootPath,
		Folder:    showFolder,
		FullPath:  fmt.Sprintf("%s/%s", rootPath, showFolder),
		Extension: builder.subtitleExtension,
	}, error
}
