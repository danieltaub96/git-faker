package util

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestIsFileExistsOk(t *testing.T) {
	f, _ := ioutil.TempFile("", "test_tmp_file")

	var exists = IsFileExists(f.Name())
	assert.True(t, exists, "Array ok test failed, should not be empty")
}

func TestIsFileExistsBad(t *testing.T) {
	var exists = IsFileExists("fake_file")
	assert.False(t, exists, "Array ok test failed, should not be empty")
}
