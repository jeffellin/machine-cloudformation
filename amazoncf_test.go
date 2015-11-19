package amazoncf

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/docker/machine/commands/mcndirs"
	//"github.com/docker/machine/drivers/amazonec2/amz"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/stretchr/testify/assert"
)

const (
	testSSHPort           = 22
	testDockerPort        = 2376
	testStoreDir          = ".store-test"
	machineTestName       = "test-host"
	machineTestDriverName = "none"
	machineTestStorePath  = "/test/path"
	machineTestCaCert     = "test-cert"
	machineTestPrivateKey = "test-key"
)

type DriverOptionsMock struct {
	Data map[string]interface{}
}

func (d DriverOptionsMock) String(key string) string {
	return d.Data[key].(string)
}

func (d DriverOptionsMock) StringSlice(key string) []string {
	return d.Data[key].([]string)
}

func (d DriverOptionsMock) Int(key string) int {
	return d.Data[key].(int)
}

func (d DriverOptionsMock) Bool(key string) bool {
	return d.Data[key].(bool)
}

func cleanup() error {
	return os.RemoveAll(testStoreDir)
}

func getTestStorePath() (string, error) {
	tmpDir, err := ioutil.TempDir("", "machine-test-")
	if err != nil {
		return "", err
	}
	mcndirs.BaseDir = tmpDir
	return tmpDir, nil
}

func getDefaultTestDriverFlags() *DriverOptionsMock {
	return &DriverOptionsMock{
		Data: map[string]interface{}{
			"name":                               "test",
			"url":                                "unix:///var/run/docker.sock",
			"swarm":                              false,
			"swarm-host":                         "",
			"swarm-master":                       false,
			"swarm-discovery":                    "",
			"cloudformation-use-private-address": true,
			"cloudformation-url":                 "someurl",
			"cloudformation-keypath":             "somepath",
			"cloudformation-keypairname":         "somekeypair",
			"cloudformation-ssh-user":            "somesshuser",
			"cloudformation-parameters":          "para",
		},
	}
}

func getTestDriver() (*Driver, error) {
	storePath, err := getTestStorePath()
	if err != nil {
		return nil, err
	}
	defer cleanup()

	d := NewDriver(machineTestName, storePath)
	d.SetConfigFromFlags(getDefaultTestDriverFlags())
	drv := d.(*Driver)
	return drv, nil
}

func TestSetConfigFromFlags(t *testing.T) {
	driver, err := getTestDriver()
	if err != nil {
		t.Fatal(err)
	}

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"amazonec2-region": "us-west-2",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)
}
