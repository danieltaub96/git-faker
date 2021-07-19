package util

import (
	"github.com/danieltaub96/git-faker/object"
	"testing"
	"time"
)

var (
	DefaultHoursRange = object.Range{Start: 0, End: 23}
)

func Test_isWorkday(t *testing.T) {
	var sunday = time.Date(1996, 9, 1, 1, 1, 1, 1, time.UTC)
	var monday = time.Date(1996, 9, 2, 1, 1, 1, 1, time.UTC)

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test Weekend", args: args{t: sunday}, want: false},
		{name: "Test Workday", args: args{t: monday}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWorkday(tt.args.t); got != tt.want {
				t.Errorf("isWorkday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randate_in_between(t *testing.T) {
	startDate := createTimeFromString("01/07/2021")
	endDate := createTimeFromString("03/07/2021")
	got, _ := Randate(startDate, endDate, DefaultHoursRange, false)

	if !(got.After(startDate) && got.Before(endDate)) {
		t.Errorf("Generated date not in between desierd dates")
	}
}

func Test_randate_workdays(t *testing.T) {
	startDate := createTimeFromString("24/07/2021")
	endDate := createTimeFromString("25/07/2021")

	_, err := Randate(startDate, endDate, DefaultHoursRange, true)

	if err != nil {
		t.Errorf("Generated random date althoug it suppose not to")
	}
}

func createTimeFromString(timeStr string) time.Time {
	parse, err := time.Parse("2/1/2006", timeStr)
	if err != nil {
		return time.Time{}
	}

	return parse
}
