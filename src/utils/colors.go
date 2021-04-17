package utils

import "github.com/fatih/color"

// ColorsStruct Define a structure of Colors that could be used everywhere
type ColorsStruct struct {
	Green  *color.Color
	Red    *color.Color
	White  *color.Color
	Yellow *color.Color
	Blue   *color.Color
}

func initializeColors() ColorsStruct {
	return ColorsStruct{
		Green:  color.New(color.FgHiGreen, color.Bold).Add(color.Concealed),
		Red:    color.New(color.FgHiRed, color.Bold),
		White:  color.New(color.FgHiWhite, color.Faint),
		Yellow: color.New(color.FgHiYellow, color.Bold),
		Blue:   color.New(color.FgHiCyan, color.Faint).Add(color.Italic).Add(color.Bold),
	}
}
