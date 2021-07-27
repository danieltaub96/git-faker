package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultInitDataFile(t *testing.T) {
	InitDataFile("")

	var dataFile = GetDataFile()
	assert.NotNil(t, dataFile, "DataFile should not be null")
}

func TestInitTwiceNotChanging(t *testing.T) {
	InitDataFile("d1")
	var d1 = GetDataFile()
	InitDataFile("d2")
	var d2 = GetDataFile()

	assert.Equal(t, d1, d2, "DataFile should not be change")
}

func TestReadDataFileOK(t *testing.T) {
	var d, err = ReadDataFile("test_data.yml")

	assert.NoError(t, err, "No errors should be accord while reading valid yaml file")
	assert.NotNil(t, d, "Data file should be init from test yml file")
}

func TestReadDataFileNotExists(t *testing.T) {
	var d, err = ReadDataFile("fake.yml")

	assert.Error(t, err, "errors should be while reading invalid yaml file")
	assert.Nil(t, d, "Data file should be nil from fake yml file")
}
