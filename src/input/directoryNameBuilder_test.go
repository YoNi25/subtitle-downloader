// +build all packageInput builder directoryName

package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildDirectoryName(t *testing.T) {
	sut := NewDirectoryNameBuilder()
	flagtests := []struct {
		directoryName               string
		expectedDirectoryName DirectoryName
	}{
		{
			"The.Falcon.and.the.Winter.Soldier/S01",
			DirectoryName{
				TvShow: "The Falcon and the Winter Soldier",
				Season: "S01",
			},
		},
		{
			"The Flash 2014/S07",
			DirectoryName{
				TvShow: "The Flash 2014",
				Season: "S07",
			},
		},
	}

	for _, test := range flagtests {
		directoryNameStruct, err := sut.build(test.directoryName)

		assert.Equal(t, test.expectedDirectoryName, directoryNameStruct)
		assert.Empty(t, err)
	}
}

func Test_InvalidDirectoryName(t *testing.T) {
	sut := NewDirectoryNameBuilder()

	directoryNameStruct, err := sut.build("an-invalid-directory-name")

	assert.Equal(t, DirectoryName{}, directoryNameStruct)
	assert.Equal(t, "unable to parse directory name 'an-invalid-directory-name'", err.Error())
}

func Test_readPathForDirectoryName(t *testing.T) {
	sut := DirectoryName{
		TvShow: "A Tv Show",
		Season: "S02",
	}
	assert.Equal(t, "A Tv Show/S02", sut.path())
}