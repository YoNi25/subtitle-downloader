package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

const configFileSource = "../../config/config.json.dist"
const configFileDestination = "../../config/config.test.json"

func init() {
	os.Setenv("ENV", "test")
}

func setup() {
	prepareConfigFile()
}

func teardown() {
	removeConfigFile()
}

func prepareConfigFile() {
	input, err := ioutil.ReadFile(configFileSource)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(configFileDestination, input, 0644)
	if err != nil {
		panic(err)
	}
}

func removeConfigFile() {
	err := os.RemoveAll(configFileDestination)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
