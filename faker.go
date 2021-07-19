package main

import (
	"fmt"
	. "github.com/danieltaub96/git-faker/git"
	. "github.com/danieltaub96/git-faker/log"
	. "github.com/danieltaub96/git-faker/object"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print git-faker version",
	}

	app := &cli.App{
		Name:    "git-faker",
		Version: version,
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "gitPath",
			Value:       ".",
			DefaultText: "Current dir",
			Usage:       "Git directory to work with",
		},
		&cli.StringFlag{
			Name:        "startDate",
			Value:       "20/01/1970",
			DefaultText: "20/01/1970",
			Usage:       "Git commit start `date`",
		},
		&cli.StringFlag{
			Name:        "endDate",
			Value:       "20/01/1980",
			DefaultText: "20/01/1980",
			Usage:       "Git commit end `date`",
		},
		&cli.StringFlag{
			Name:        "commitsPerDay",
			Value:       "2-10",
			DefaultText: "2-10",
		},
		&cli.StringFlag{
			Name:        "hours",
			Value:       "0-23",
			DefaultText: "0-23",
		},
		&cli.BoolFlag{
			Name:        "workdaysOnly",
			DefaultText: "false",
		},
		&cli.BoolFlag{
			Name:        "rewriteHistory",
			DefaultText: "false",
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			DefaultText: "false",
		},
	}

	app.Action = func(c *cli.Context) error {
		var output string

		InitLogger()

		log.Info("Reading user args...")
		var userInput, inputErr = InitUserInput(c)

		if inputErr != nil {
			return inputErr
		}

		// enable verbose mode if needed
		CheckVerboseMode(userInput)

		// Check for git repository
		CheckGitInit(userInput)

		var repo, _ = OpenGitRepo(userInput)

		// Rewrite history check
		if userInput.RewriteHistory {
			CleanGitWorktree(userInput)
		}

		// Build commit history
		BuildCommitHistory(userInput, repo)

		fmt.Println(output)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		//log.Fatal(err)
	}
}

func InitUserInput(c *cli.Context) (UserInput, cli.ExitCoder) {
	var startDate, startDateErr = time.Parse("2/1/2006", c.String("startDate"))
	var endDate, endDateErr = time.Parse("2/1/2006", c.String("endDate"))

	if err := ValidateErr(startDateErr, "failed to parse startDate"); err != nil {
		return UserInput{}, err
	}

	if err := ValidateErr(endDateErr, "failed to parse endDate"); err != nil {
		return UserInput{}, err
	}

	var commitsPerDay = c.String("commitsPerDay")
	var commitsRange = Range{}

	if strings.Contains(commitsPerDay, "-") {
		var splitedRange = strings.Split(commitsPerDay, "-")
		commitsRange.Start, _ = strconv.Atoi(splitedRange[0])
		commitsRange.End, _ = strconv.Atoi(splitedRange[1])
	}

	var hours = c.String("hours")
	var hoursRange = Range{}

	if strings.Contains(hours, "-") {
		var splitedRange = strings.Split(hours, "-")
		hoursRange.Start, _ = strconv.Atoi(splitedRange[0])
		hoursRange.End, _ = strconv.Atoi(splitedRange[1])
	}

	u := UserInput{
		GitPath:        c.String("gitPath"),
		StartDate:      startDate,
		EndDate:        endDate,
		WorkdaysOnly:   c.Bool("workdaysOnly"),
		RewriteHistory: c.Bool("rewriteHistory"),
		CommitsPerDay:  commitsRange,
		Hours:          hoursRange,
		Verbose:        c.Bool("verbose"),
	}

	log.Debugln("loaded arguments", u)

	return u, nil
}

func ValidateErr(err error, message string) cli.ExitCoder {
	if err != nil {
		log.Errorln(message)
		os.Exit(1)
		return cli.Exit(err, 1)
	}

	return nil
}
