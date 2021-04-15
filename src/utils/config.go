package utils

import (
	"github.com/tkanos/gonfig"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// DirPathsConfig Define a structure for the DirPaths configurations
// Available contains all paths that could be used to download a subtitle
// Default is the default path to use
type DirPathsConfig struct {
	Available map[string]string `json:"Available"`
	Default string `json:"Default"`
}

// LanguagesConfig Define a structure for the Language configurations
// Available contains all languages that could be used to download a subtitle
// Default is the default language to use
type LanguagesConfig struct {
	Available map[string]string `json:"Available"`
	Default string `json:"Default"`
}

// Configuration Define a structure for the configuration that could be used everywhere
type Configuration struct {
	DirPathsConfig    DirPathsConfig  `json:"DirPaths"`
	LanguagesConfig   LanguagesConfig `json:"Languages"`
	SubtitleExtension string          `json:"SubtitleExtension"`
}

func initializeConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getConfigFileName(), &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"../../", "config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
