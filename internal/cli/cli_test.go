package commandline

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type option struct {
	flag  string
	value string
}

func TestFlagParsing(t *testing.T) {
	assert := assert.New(t)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = generateArgs()

	conf, _ := CreateConfFromCli()

	assert.EqualValues(555, conf.Capacity, "Expected the flag --capacity to set conf.Capacity to 555, got %d", conf.Capacity)
	assert.Equal("mb", conf.CapacityUnit, "Expected the flag --capacity-unit to set conf.CapacityUnit to \"mb\", got %q", conf.CapacityUnit)
	assert.Equal("localhost", conf.Hostname, "Expected the flag --hostname to set conf.Hostname to \"localhost\", got %q", conf.Hostname)
	assert.EqualValues(8080, conf.Port, "Expected the flag --port to set conf.Port to 8080, got %d", conf.Port)
	assert.Equal("https://jsonplaceholder.typicode.com", conf.ApiUrl, "Expected the flag --api-url to set conf.ApiUrl to \"https://jsonplaceholder.typicode.com\", got %q", conf.ApiUrl)
	assert.Equal("logfile.log", conf.LogFilePath, "Expected the flag --logfile to set conf.LogFilePath to \"logfile.log\", got %q", conf.LogFilePath)
	assert.Equal([]string{"/posts", "/posts/:id"}, conf.Cache["GET"], "Expected the flag --cache:GET with comma-separated values to set conf.Cache[\"GET\"] to both of %v, got %v", []string{"/posts", "/posts/:id"}, conf.Cache["GET"])
	assert.Equal([]string{"/posts", "/posts/:id"}, conf.Cache["HEAD"], "Expected the flag --cache:HEAD to set conf.Cache[\"HEAD\"] to %v, got %v", []string{"/posts", "/posts/:id"}, conf.Cache["HEAD"])
	assert.Equal([]string{"/posts"}, conf.Bust["GET"]["/todos"], "Expected the flag --bust:GET to set conf.Bust[\"GET\"][\"/todos\"] to %v, got %v", []string{"/posts"}, conf.Bust["GET"]["/todos"])
	assert.Equal([]string{"/posts"}, conf.Bust["HEAD"]["/todos"], "Expected the flag --bust:HEAD to set conf.Bust[\"HEAD\"][\"/todos\"] to %v, got %v", []string{"/posts"}, conf.Bust["HEAD"]["/todos"])
	assert.Equal([]string{"/posts"}, conf.Bust["POST"]["/posts"], "Expected the flag --bust:POST to set conf.Bust[\"POST\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["POST"]["/posts"])
	assert.Equal([]string{"^GET:/posts", "^HEAD:/posts"}, conf.Bust["PUT"]["/posts"], "Expected the flag --bust:PUT with multiple patterns to set conf.Bust[\"PUT\"][\"/posts\"] to both of %v, got %v", []string{"^GET:/posts", "^HEAD:/posts"}, conf.Bust["PUT"]["/posts"])
	assert.Equal([]string{"/posts/:id"}, conf.Bust["PUT"]["/posts/:id"], "Expected the flag --bust:PUT to set conf.Bust[\"PUT\"][\"/posts/:id\"] to %v, got %v", []string{"/posts/:id"}, conf.Bust["PUT"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["DELETE"]["/posts/:id"], "Expected the flag --bust:DELETE to set conf.Bust[\"DELETE\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["DELETE"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["PATCH"]["/posts/:id"], "Expected the flag --bust:PATCH to set conf.Bust[\"PATCH\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["PATCH"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["TRACE"]["/posts/:id"], "Expected the flag --bust:TRACE to set conf.Bust[\"TRACE\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["TRACE"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["CONNECT"]["/posts"], "Expected the flag --bust:CONNECT to set conf.Bust[\"CONNECT\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["CONNECT"]["/posts"])
	assert.Equal([]string{"/posts"}, conf.Bust["OPTIONS"]["/posts"], "Expected the flag --bust:OPTIONS to set conf.Bust[\"OPTIONS\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["OPTIONS"]["/posts"])
}

func TestConfigFileParsing(t *testing.T) {
	assert := assert.New(t)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--config", "./testdata/test.config.json"}

	conf, _ := CreateConfFromCli()

	assert.EqualValues(555, conf.Capacity, "Expected the prop 'capacity' to set conf.Capacity to 555, got %d", conf.Capacity)
	assert.Equal("mb", conf.CapacityUnit, "Expected the prop 'capacityUnit' to set conf.CapacityUnit to \"mb\", got %q", conf.CapacityUnit)
	assert.Equal("localhost", conf.Hostname, "Expected the prop 'hostname' to set conf.Hostname to \"localhost\", got %q", conf.Hostname)
	assert.EqualValues(8080, conf.Port, "Expected the prop 'port' to set conf.Port to 8080, got %d", conf.Port)
	assert.Equal("https://jsonplaceholder.typicode.com", conf.ApiUrl, "Expected the prop 'apiUrl' to set conf.ApiUrl to \"https://jsonplaceholder.typicode.com\", got %q", conf.ApiUrl)
	assert.Equal("logfile.log", conf.LogFilePath, "Expected the prop 'logFilePath' to set conf.LogFilePath to \"logfile.log\", got %q", conf.LogFilePath)
	assert.Equal([]string{"/posts", "/posts/:id"}, conf.Cache["GET"], "Expected the prop cache.GET to set conf.Cache[\"GET\"] to %v, got %v", []string{"/posts", "/posts/:id"}, conf.Cache["GET"])
	assert.Equal([]string{"/posts", "/posts/:id"}, conf.Cache["HEAD"], "Expected the prop cache.HEAD to set conf.Cache[\"HEAD\"] to %v, got %v", []string{"/posts", "/posts/:id"}, conf.Cache["HEAD"])
	assert.Equal([]string{"/posts"}, conf.Bust["GET"]["/todos"], "Expected the prop bust.GET to set conf.Bust[\"GET\"][\"/todos\"] to %v, got %v", []string{"/posts"}, conf.Bust["GET"]["/todos"])
	assert.Equal([]string{"/posts"}, conf.Bust["HEAD"]["/todos"], "Expected the prop bust.HEAD to set conf.Bust[\"HEAD\"][\"/todos\"] to %v, got %v", []string{"/posts"}, conf.Bust["HEAD"]["/todos"])
	assert.Equal([]string{"/posts"}, conf.Bust["POST"]["/posts"], "Expected the prop bust.POST to set conf.Bust[\"POST\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["POST"]["/posts"])
	assert.Equal([]string{"^GET:/posts", "^HEAD:/posts"}, conf.Bust["PUT"]["/posts"], "Expected the prop bust.PUT to set conf.Bust[\"PUT\"][\"/posts\"] to %v, got %v", []string{"^GET:/posts", "^HEAD:/posts"}, conf.Bust["PUT"]["/posts"])
	assert.Equal([]string{"/posts/:id"}, conf.Bust["PUT"]["/posts/:id"], "Expected the prop bust.PUT to set conf.Bust[\"PUT\"][\"/posts/:id\"] to %v, got %v", []string{"/posts/:id"}, conf.Bust["PUT"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["DELETE"]["/posts/:id"], "Expected the prop bust.DELETE to set conf.Bust[\"DELETE\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["DELETE"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["PATCH"]["/posts/:id"], "Expected the prop bust.PATCH to set conf.Bust[\"PATCH\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["PATCH"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["TRACE"]["/posts/:id"], "Expected the prop bust.TRACE to set conf.Bust[\"TRACE\"][\"/posts/:id\"] to %v, got %v", []string{"/posts"}, conf.Bust["TRACE"]["/posts/:id"])
	assert.Equal([]string{"/posts"}, conf.Bust["CONNECT"]["/posts"], "Expected the prop bust.CONNECT to set conf.Bust[\"CONNECT\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["CONNECT"]["/posts"])
	assert.Equal([]string{"/posts"}, conf.Bust["OPTIONS"]["/posts"], "Expected the prop bust.OPTIONS to set conf.Bust[\"OPTIONS\"][\"/posts\"] to %v, got %v", []string{"/posts"}, conf.Bust["OPTIONS"]["/posts"])
}

func TestFlagsOverwriteConfigFile(t *testing.T) {
	assert := assert.New(t)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--config", "./testdata/test.config.json", "--api-url", "https://test.com"}

	conf, err := CreateConfFromCli()

	assert.NoError(err, "Expected new config to be created without any errors")
	assert.Equal("https://test.com", conf.ApiUrl, "Expected the passed flag (--api-url) to overwrite the prop (apiUrl) specified in the config file, but got %q", conf.ApiUrl)
}

func generateArgs() []string {
	return []string{"cmd",
		"--capacity", "555",
		"--capacity-unit", "mb",
		"--hostname", "localhost",
		"--port", "8080",
		"--api-url", "https://jsonplaceholder.typicode.com/",
		"--logfile", "logfile.log",
		"--cache:GET", "/posts",
		"--cache:GET", "/posts/:id",
		"--cache:HEAD", "/posts,/posts/:id",
		"--bust:GET", "/todos=>/posts",
		"--bust:HEAD", "/todos=>/posts",
		"--bust:POST", "/posts=>/posts",
		"--bust:PUT", "/posts=>^GET:/posts||^HEAD:/posts",
		"--bust:PUT", "/posts/:id=>/posts/:id",
		"--bust:DELETE", "/posts/:id=>/posts",
		"--bust:PATCH", "/posts/:id=>/posts",
		"--bust:TRACE", "/posts/:id=>/posts",
		"--bust:CONNECT", "/posts=>/posts",
		"--bust:OPTIONS", "/posts=>/posts",
	}
}
