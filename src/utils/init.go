package utils

var Config Configuration
var Colors ColorsStruct

func Init() {
	Config = initializeConfig()
	Colors = initializeColors()
}