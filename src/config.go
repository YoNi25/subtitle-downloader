package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Configuration struct {
	ServerDirPath  string
	DesktopDirPath string
	DefaultLanguage string
	SubtitleExtension string
}

func InitializeConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getConfigFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}

	return configuration
}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"../", "config/", "config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
