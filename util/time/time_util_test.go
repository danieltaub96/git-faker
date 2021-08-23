package util

import (
	"github.com/danieltaub96/git-faker/object"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	DefaultHoursRange = object.Range{Start: 0, End: 23}
	DefaultTime       = time.Now()
)

func TestIsInHoursOK(t *testing.T) {
	isInHours := IsInHours(DefaultTime, DefaultHoursRange)
	assert.True(t, isInHours)
}

func TestIsInHoursNotInLimitRange(t *testing.T) {
	testTime := time.Date(DefaultTime.Year(), DefaultTime.Month(), DefaultTime.Day(), 10, 0, 0, 0, DefaultTime.Location())
	isInHours := IsInHours(testTime, object.Range{Start: 1, End: 2})

	assert.False(t, isInHours)
}


func TestIsInHoursIsInLimitRangeOk(t *testing.T) {
	testTime := time.Date(DefaultTime.Year(), DefaultTime.Month(), DefaultTime.Day(), 10, 0, 0, 0, DefaultTime.Location())
	isInHours := IsInHours(testTime, object.Range{Start: 1, End: 12})

	assert.True(t, isInHours)
}

func TestIsWorkdayOk(t *testing.T) {
	testTime := time.Date(DefaultTime.Year(), DefaultTime.Month(), 2, 10, 0, 0, 0, DefaultTime.Location())

	isWorkday := IsWorkday(testTime)
	assert.True(t, isWorkday)
}

func TestIsNotWorkdayOk(t *testing.T) {
	testTime := time.Date(DefaultTime.Year(), DefaultTime.Month(), 7, 10, 0, 0, 0, DefaultTime.Location())

	isWorkday := IsWorkday(testTime)
	assert.False(t, isWorkday)
}

func TestGenerateHoursInBetweenOk(t *testing.T) {
	generatedTime := GenerateHoursInBetween(DefaultTime, object.Range{
		Start: 3,
		End:   5,
	})

	assert.NotNil(t, generatedTime)
	assert.GreaterOrEqual(t, generatedTime.Hour(), 3)
	assert.LessOrEqual(t, generatedTime.Hour(), 5)
}

func TestGenerateHoursInBetweenOk2(t *testing.T) {
	generatedTime := GenerateHoursInBetween(DefaultTime, object.Range{
		Start: 3,
		End:   4,
	})

	assert.NotNil(t, generatedTime)
	assert.Equal(t, generatedTime.Hour(), 3)
}