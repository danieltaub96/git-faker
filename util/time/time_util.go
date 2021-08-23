package util

import (
	. "github.com/danieltaub96/git-faker/object"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func GenerateDateInBetween(startDate time.Time, endDate time.Time) time.Time {
	min := startDate.Unix()
	max := endDate.Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0)

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GenerateHoursInBetween(t time.Time, hours Range) time.Time {
	timeMin := time.Date(t.Year(), t.Month(), t.Day(), hours.Start, 0, 0, t.Nanosecond(), t.Location()).Unix()
	timeMax := time.Date(t.Year(), t.Month(), t.Day(), hours.End, 0, 0, t.Nanosecond(), t.Location()).Unix()
	delta := timeMax - timeMin

	sec := rand.Int63n(delta) + timeMin

	return time.Unix(sec, 0)
}

func IsWorkday(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
}

func IsInHours(t time.Time, hours Range) bool {
	return t.Hour() >= hours.Start && t.Hour() <= hours.End
}

func RandomDate(startDate time.Time, endDate time.Time, workdaysOnly bool) (time.Time, error) {
	var randomDate = GenerateDateInBetween(startDate, endDate)
	workdayRetryCounter := 1000

	if workdaysOnly {
		for !IsWorkday(randomDate) && workdayRetryCounter > 0 {
			workdayRetryCounter--
			randomDate = GenerateDateInBetween(startDate, endDate)
		}
	}

	if workdayRetryCounter == 0 {
		log.Errorf("cannot generate dates between, maybe the dates is limited in such way it cant be generated?")
	}

	return randomDate, nil
}

func RandomHours(t time.Time, hours Range) (time.Time, error) {
	var randomDate = GenerateHoursInBetween(t, hours)
	hoursRetryCounter := 1000

	for !IsInHours(randomDate, hours) && hoursRetryCounter > 0 {
		hoursRetryCounter--
		randomDate = GenerateHoursInBetween(t, hours)
	}

	if hoursRetryCounter == 0 {
		log.Errorf("cannot generate dates between, maybe the dates is limited in such way it cant be generated?")
	}

	return randomDate, nil
}
