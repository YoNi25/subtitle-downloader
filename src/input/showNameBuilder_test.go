// +build all packageInput builder showName

package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildShowName(t *testing.T) {
	sut := NewShowNameBuilder()
	flagtests := []struct {
		showName               string
		expectedShowNameStruct ShowName
	}{
		{
			"The.Falcon.and.the.Winter.Soldier.S01E01.WEBRip.x264-PHOENiX[eztv.re]",
			ShowName{
				TvShow:   "The Falcon And The Winter Soldier",
				Season:   "S01",
				Episode:  "E01",
				Version:  "PHOENIX",
				Source:   "eztv.re",
				Fullname: "The.Falcon.and.the.Winter.Soldier.S01E01.WEBRip.x264-PHOENiX[eztv.re]",
			},
		},
		{
			"The.Flash.2014.S07E03.HDTV.x264-cakes[eztv.re]",
			ShowName{
				TvShow:   "The Flash 2014",
				Season:   "S07",
				Episode:  "E03",
				Version:  "CAKES",
				Source:   "eztv.re",
				Fullname: "The.Flash.2014.S07E03.HDTV.x264-cakes[eztv.re]",
			},
		},
		{
			"Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
			ShowName{
				TvShow:   "Age Of Samurai Battle For Japan",
				Season:   "S01",
				Episode:  "E01",
				Version:  "EXTREME",
				Source:   "",
				Fullname: "Age.of.Samurai.Battle.for.Japan.S01E01.VOSTFR.WEB.XviD-EXTREME",
			},
		},
	}

	for _, test := range flagtests {
		showNameStruct, err := sut.build(test.showName)

		assert.Equal(t, test.expectedShowNameStruct, showNameStruct)
		assert.Empty(t, err)
	}
}

func Test_InvalidShowName(t *testing.T) {
	sut := NewShowNameBuilder()

	showNameStruct, err := sut.build("an-invalid-show-name")

	assert.Equal(t, ShowName{}, showNameStruct)
	assert.Equal(t, "Unable to parse Show name 'an-invalid-show-name'", err.Error())
}
