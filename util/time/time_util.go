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
	return time.Unix(sec, 0)
}

func IsWorkday(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
}

func IsInHours(t time.Time, hours Range) bool {
	return t.Hour() >= hours.Start && t.Hour() <= hours.End
}

func Randate(startDate time.Time, endDate time.Time, hours Range, workdaysOnly bool) (time.Time, error) {
	var randomDate = GenerateDateInBetween(startDate, endDate)
	retryCounter := 10000

	for !IsInHours(randomDate, hours) && retryCounter > 0 {
		workdayRetryCounter := 1000
		retryCounter--

		randomDate = GenerateDateInBetween(startDate, endDate)

		if workdaysOnly {
			for !IsWorkday(randomDate) && workdayRetryCounter > 0 {
				workdayRetryCounter--
				randomDate = GenerateDateInBetween(startDate, endDate)
			}
		}
	}

	if retryCounter == 0 {
		log.Errorf("cannot generate dates between, maybe the dates is limited in such way it cant be generated?")
	}

	return randomDate, nil
}
