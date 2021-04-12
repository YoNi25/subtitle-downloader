package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_initializeConfig(t *testing.T) {
	expectedConfig := Configuration{
		DirPathsConfig: DirPathsConfig{
			Available: map[string]string{
				"Server":  "/a/server/dir/path",
				"Desktop": "../../downloaded",
				"Current": ".",
			},
			Default: "../../downloaded",
		},
		LanguagesConfig: LanguagesConfig{
			Available: map[string]string{
				"French":  "French",
				"English": "English",
			},
			Default: "French",
		},
		SubtitleExtension: "srt",
	}

	config := initializeConfig()

	assert.Equal(t, expectedConfig, config)
}

func Test_failWhenConfigFileDoesntExist(t *testing.T) {
	os.Setenv("ENV", "env")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = initializeConfig()
}
