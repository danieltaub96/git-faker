package cli

import (
	"fmt"
	. "github.com/danieltaub96/git-faker/git"
	. "github.com/danieltaub96/git-faker/log"
	. "github.com/danieltaub96/git-faker/object"
	util "github.com/danieltaub96/git-faker/util/file"
	. "github.com/danieltaub96/git-faker/util/yaml"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func InitCliApp(version string) {
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
		&cli.IntFlag{
			Name:        "days",
			Value:       3,
			DefaultText: "3",
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
		&cli.StringFlag{
			Name:  "data-file",
			Value: "data-file ",
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			DefaultText: "false",
		},
		&cli.BoolFlag{
			Name:        "force",
			Aliases:     []string{"f"},
			DefaultText: "false",
		},
	}

	app.Action = func(c *cli.Context) error {
		InitLogger()

		log.Info("Reading user args...")
		var userInput, inputErr = initUserInput(c)

		if inputErr != nil {
			return inputErr
		}

		// enable verbose mode if needed
		CheckVerboseMode(userInput)

		InitDataFile(userInput.DataFile)

		// Check for git repository
		if err := CheckGitInit(userInput); err != nil {
			exit(1, err)
		}

		var repo, err = OpenGitRepo(userInput)

		if err != nil {
			exit(1, err)
		}

		// ask permissions
		if !userInput.Force {
			var gitPathAbs = util.AbsPath(userInput.GitPath)

			fmt.Printf("Do you want to perform an git changes on this git repo: %s [y/n]\n", gitPathAbs)

			// var then variable name then variable type
			var userPermissions string

			// Taking input from user
			fmt.Scanln(&userPermissions)

			if userPermissions != "y" {
				os.Exit(0)
			}
		}

		log.Infoln("Starting process git-faker")

		// Rewrite history check
		if userInput.RewriteHistory {
			CleanGitWorktree(userInput)
		}

		// Build commit history
		log.Infoln("Building commit history")
		if err := BuildCommitHistory(userInput, repo); err != nil {
			exit(1, err)
		}
		log.Infoln("Finished building commit history")


		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		exit(1, err)
	}

	log.Infof("Finished git-faker process")
}

func initUserInput(c *cli.Context) (UserInput, cli.ExitCoder) {
	var startDate, startDateErr = time.Parse("2/1/2006", c.String("startDate"))
	var endDate, endDateErr = time.Parse("2/1/2006", c.String("endDate"))

	if err := validateErr(startDateErr, "failed to parse startDate"); err != nil {
		return UserInput{}, err
	}

	if err := validateErr(endDateErr, "failed to parse endDate"); err != nil {
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
		Days:           c.Int("days"),
		WorkdaysOnly:   c.Bool("workdaysOnly"),
		RewriteHistory: c.Bool("rewriteHistory"),
		CommitsPerDay:  commitsRange,
		Hours:          hoursRange,
		DataFile:       c.String("data-file"),
		Verbose:        c.Bool("verbose"),
		Force:          c.Bool("force"),
	}

	log.Debugln("loaded arguments", u)

	return u, nil
}

func validateErr(err error, message string) cli.ExitCoder {
	if err != nil {
		exit(1, message)
	}

	return nil
}

func exit(code int, i interface{}) {
	log.Errorln("Error while running git-faker, Aborting...", i)
	os.Exit(1)
}