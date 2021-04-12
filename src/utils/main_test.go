package utils

import (
	"os"
)

func init() {
	os.Setenv("ENV", "test")
}
