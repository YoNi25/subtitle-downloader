package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func Test_buildDirPath(t *testing.T) {
	flagtests := []struct {
		dirPathDigit          int
		ShowName              ShowName
		expectedDirPathStruct DirPath
	}{
		{
			1,
			ShowName{
				TvShow: "A Tv Show",
				Season: "S02",
			},
			DirPath{
				RootPath: "/server/dir/path",
				Folder: "A Tv Show/S02",
				FullPath: "/server/dir/path/A Tv Show/S02",
				Extension: "srt",
			},
		},
		{
			2,
			ShowName{
				TvShow: "A Tv Show",
				Season: "S02",
			},
			DirPath{
				RootPath: "/desktop/dir/path",
				Folder: "A Tv Show/S02",
				FullPath: "/desktop/dir/path/A Tv Show/S02",
				Extension: "srt",
			},
		},
	}

	sut := NewDirPathBuilder(utils.Config)
	for _, test := range flagtests {
		dirPathStruct, err := sut.build(test.dirPathDigit, test.ShowName)
		assert.Equal(t, test.expectedDirPathStruct, dirPathStruct)
		assert.Empty(t, err)
	}
}

func Test_BuildDirPathWithDefaultValue(t *testing.T) {
	sut := NewDirPathBuilder(utils.Config)

	expectedDirPathStruct := DirPath{
		RootPath: "/server/dir/path",
		Folder: "A Tv Show/S02",
		FullPath: "/server/dir/path/A Tv Show/S02",
		Extension: "srt",
	}
	dirPathStruct, err := sut.build(-1, ShowName{
		TvShow: "A Tv Show",
		Season: "S02",
	})
	assert.Equal(t, expectedDirPathStruct, dirPathStruct)
	assert.Equal(t, "No DirPath matches with -1. Using default DirPath - '/server/dir/path'", err.Error())
}