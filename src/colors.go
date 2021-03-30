package main

import "github.com/fatih/color"

type Colors struct {
	green  *color.Color
	red    *color.Color
	white  *color.Color
	yellow *color.Color
}

func InitializeColors() Colors {
	return Colors{
		green:  color.New(color.FgHiGreen, color.Bold).Add(color.Concealed),
		red:    color.New(color.FgHiRed, color.Bold),
		white:  color.New(color.FgHiWhite, color.Faint),
		yellow: color.New(color.FgHiYellow, color.Bold),
	}
}
