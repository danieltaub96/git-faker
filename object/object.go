package object

import (
	"time"
)

type UserInput struct {
	GitPath        string
	StartDate      time.Time
	EndDate        time.Time
	Days           int
	CommitsPerDay  Range
	Hours          Range
	WorkdaysOnly   bool
	RewriteHistory bool
	DataFile       string
	Verbose        bool
	Force          bool
}

type Range struct {
	Start int
	End   int
}

type DataFile struct {
	Names    []string `yaml:"names"`
	Messages []string `yaml:"messages"`
	Emails   []string `yaml:"emails"`
}

func (d *DataFile) SetDefaultMessages() {
	*d = DataFile{
		Names:    d.Names,
		Emails:   d.Emails,
		Messages: []string{"Hello World"},
	}
}
func (d *DataFile) SetDefaultNames() {
	*d = DataFile{
		Messages: d.Messages,
		Emails:   d.Emails,
		Names:    []string{"Daniel Taub"},
	}
}
func (d *DataFile) SetDefaultEmails() {
	*d = DataFile{
		Names:    d.Names,
		Messages: d.Messages,
		Emails:   []string{"bla@gmail.com"},
	}
}
func (d *DataFile) LoadDefaults() {
	if d == nil {
		*d = DataFile{}
	}

	d.SetDefaultMessages()
	d.SetDefaultNames()
	d.SetDefaultEmails()
}
