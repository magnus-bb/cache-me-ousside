package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Make sure all props are set on the config object properlu
func TestLoadProps(t *testing.T) {
	assert := assert.New(t)

	configPath := "testdata/test.config.json5"

	assert.FileExists(configPath, "Expected test configuration file to exist for test to work")

	config := Load(configPath)

	assert.NotZero(config.Capacity, "Expected required prop config.Capacity to be loaded correctly as a non-zero value")
	assert.Positive(config.Capacity, "Expected required prop config.Capacity to be loaded correctly as a positive number")

	assert.NotZero(config.ApiUrl, "Expected required prop config.ApiUrl to be loaded correctly as a non-zero value")

	// TODO: change to be a map of []string with GET and HEAD (and more?)
	assert.NotEmpty(config.Cache, "Expected config.Cache to not be empty when given a valid config file")

	assert.NotEmpty(config.BustMap, "Expected config.BustMap to not be empty when given a valid config file")

	assert.NotEmpty(config.BustMap["POST"]["/posts"], "Expected config.BustMap's POST /posts endpoint to not be empty when given a valid config file")

	assert.NotEmpty(config.BustMap["PUT"]["/posts/:slug"], "Expected config.BustMap's PUT /posts/:slug endpoint to not be empty when given a valid config file")

	assert.NotEmpty(config.BustMap["DELETE"]["/posts/:id"], "Expected config.BustMap's DELETE /posts/:id endpoint to not be empty when given a valid config file")
}

func TestBadPathPanic(t *testing.T) {
	configPath := "testdata/does.not.exist.json5"

	assert.NoFileExists(t, configPath, "Expected test configuration file to not exist for test to work")

	assert.Panics(t, func() { Load(configPath) }, "Expected config.Load to panic when the config file does not exist")
}

// TODO: edit this when more required props are added to config
func TestRequiredProps(t *testing.T) {
	missingProps := []string{
		"capacity",
		"apiUrl",
		"cacheMap", // TODO: create a version where the prop exists but there are empty slices etc
	}

	for _, prop := range missingProps {
		configPath := "testdata/missing." + prop + ".json5"

		assert.FileExists(t, configPath, "Expected test configuration file to exist for test to work")

		assert.Panics(t, func() { Load(configPath) }, "Expected config.Load to panic when the file: %s is missing the required prop: %s", configPath, prop)
	}
}

func TestTrimTrailingSlash(t *testing.T) {
	configPath := "testdata/test.config.json5"

	assert.FileExists(t, configPath, "Expected test configuration file to exist for test to work")

	config := Load(configPath)

	assert.Equal(t, config.ApiUrl, "https://jsonplaceholder.typicode.com", "Expected config.Load to remove trailing slashes from the apiUrl")
}