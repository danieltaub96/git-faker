package object

import "time"

type UserInput struct {
	GitPath        string
	StartDate      time.Time
	EndDate        time.Time
	CommitsPerDay  Range
	Hours          Range
	WorkdaysOnly   bool
	RewriteHistory bool
	Verbose        bool
}

type Range struct {
	Start int
	End   int
}
