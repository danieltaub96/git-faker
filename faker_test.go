package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
	"reflect"
	"testing"
	"time"
)

func TestCheckGitInit(t *testing.T) {
	type args struct {
		u UserInput
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test git repository"},
		{name: "Test regular directory"},
		{name: "Test file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestBuildCommitHistory(t *testing.T) {
	type args struct {
		u    UserInput
		repo *git.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCheckGitInit1(t *testing.T) {
	type args struct {
		u UserInput
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCheckTimeErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestCleanGitWorktree(t *testing.T) {
	type args struct {
		u    UserInput
		repo *git.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestInitUserInput(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name string
		args args
		want UserInput
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitUserInput(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitUserInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenGitRepo(t *testing.T) {
	type args struct {
		u UserInput
	}
	tests := []struct {
		name string
		args args
		want *git.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OpenGitRepo(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenGitRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := isWorkday(tt.args.t); got != tt.want {
				t.Errorf("isWorkday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randate(t *testing.T) {
	type args struct {
		startDate    time.Time
		endDate      time.Time
		workdaysOnly bool
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randate(tt.args.startDate, tt.args.endDate, tt.args.workdaysOnly); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randate() = %v, want %v", got, tt.want)
			}
		})
	}
}
