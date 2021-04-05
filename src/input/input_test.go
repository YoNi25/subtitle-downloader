package input

import (
	"github.com/fatih/color"
	"utils"
)

func init() {
	colors = utils.ColorsStruct{
		Yellow: color.New(),
		Red : color.New(),
	}
}