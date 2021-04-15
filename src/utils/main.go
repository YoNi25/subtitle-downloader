package utils

// Colors expose the configuration to be used everywhere else
var Colors ColorsStruct
// Config expose the configuration to be used everywhere else
var Config Configuration

// NewColors Instantiate a Colors struct
func NewColors() ColorsStruct{
	return initializeColors()
}

// NewConfig Instantiate a Configuration struct
func NewConfig() Configuration {
	return initializeConfig()
}