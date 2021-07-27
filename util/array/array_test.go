package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyRandomValueFromArray(t *testing.T) {
	var emptyArr []string

	var val = RandomValueFromArray(emptyArr)
	assert.Empty(t, val, "Array empty test failed, should be empty string")
}

func TestOkRandomValueFromArray(t *testing.T) {
	var arr = []string{"git", "faker", "array", "test"}

	var val = RandomValueFromArray(arr)
	assert.NotEmpty(t, val, "Array ok test failed, should not be empty")
}
