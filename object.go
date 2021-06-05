package main

import "time"

type UserInput struct {
	gitPath        string
	startDate      time.Time
	endDate        time.Time
	commitsPerDay  Range
	workdaysOnly   bool
	rewriteHistory bool
}

type Range struct {
	start int
	end   int
}
