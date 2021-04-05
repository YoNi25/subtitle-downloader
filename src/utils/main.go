package utils

// Colors expose the configuration to be used everywhere else
var Colors ColorsStruct
// Config expose the configuration to be used everywhere else
var Config Configuration

// Init initialize utils global variables
func Init() {
	Colors = initializeColors()
	Config = initializeConfig()
}