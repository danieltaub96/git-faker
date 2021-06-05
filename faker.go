package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Handle args flags
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "gitPath",
			Value:       ".",
			DefaultText: "Current dir",
			Usage:       "Git directory to work with",
		},
		&cli.StringFlag{
			Name:        "startDate",
			Value:       "01/01/1970",
			DefaultText: "01/01/1970",
			Usage:       "Git commit start `date`",
		},
		&cli.StringFlag{
			Name:        "endDate",
			Value:       "01/01/1980",
			DefaultText: "01/01/1980",
			Usage:       "Git commit end `date`",
		},
		&cli.StringFlag{
			Name:        "commitsPerDay",
			Value:       "1-5",
			DefaultText: "1-5",
		},
		&cli.BoolFlag{
			Name:        "workdaysOnly",
			DefaultText: "false",
		},
		&cli.BoolFlag{
			Name:        "rewriteHistory",
			DefaultText: "false",
		},
	}

	app.Action = func(c *cli.Context) error {
		var output string

		var userInput = InitUserInput(c)

		// Check for git repository
		CheckGitInit(userInput)

		var repo = OpenGitRepo(userInput)

		// Rewrite history check
		if userInput.rewriteHistory {
			CleanGitWorktree(userInput, repo)
		}

		// Build commit history
		//BuildCommitHistory(userInput)

		fmt.Println(output)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func CleanGitWorktree(u UserInput, repo *git.Repository) {
	var w, err = repo.Worktree()

	if err == nil {
		w.Checkout(&git.CheckoutOptions{
			Branch: "git-faker",
			Create: true,
		})

		w.Add(".")

		w.Commit("Git faker orphan branch", &git.CommitOptions{})

		repo.DeleteBranch("master")
	}
}

func OpenGitRepo(u UserInput) *git.Repository {
	repo, _ := git.PlainOpen(u.gitPath)

	return repo
}

func InitUserInput(c *cli.Context) UserInput {
	var startDate, startDateErr = time.Parse("15/01/2006", c.String("startDate"))
	var endDate, endDateErr = time.Parse("15/01/2006", c.String("endDate"))

	CheckTimeErr(startDateErr)
	CheckTimeErr(endDateErr)

	var commitsPerDay = c.String("commitsPerDay")
	var commitsRange = Range{}

	if strings.Contains(commitsPerDay, "-") {
		var splitedRange = strings.Split(commitsPerDay, "-")
		commitsRange.start, _ = strconv.Atoi(splitedRange[0])
		commitsRange.end, _ = strconv.Atoi(splitedRange[1])
	}

	return UserInput{
		gitPath:        c.String("gitPath"),
		startDate:      startDate,
		endDate:        endDate,
		workdaysOnly:   c.Bool("workdaysOnly"),
		rewriteHistory: c.Bool("rewriteHistory"),
		commitsPerDay:  commitsRange,
	}
}

func CheckTimeErr(err error) {
	if err != nil {
		print("error while trying to parse date")
		panic(err)
	}
}

func CheckGitInit(u UserInput) {
	if _, err := os.Stat(path.Join(u.gitPath, ".git")); os.IsNotExist(err) {
		print("Git is not init for the current directory")
	}
}

func BuildCommitHistory(u UserInput, repo *git.Repository) {
	var startDate = u.startDate
	var endDate = u.endDate

	w, _ := repo.Worktree()

	rand.Seed(time.Now().UnixNano())

	minCommits := u.commitsPerDay.start
	maxCommits := u.commitsPerDay.end

	var commitsInDay = rand.Intn(maxCommits-minCommits) + minCommits
	log.Printf("Commits: %d\n", commitsInDay)
	for i := 1; i <= commitsInDay; i++ {
		log.Println(randate(startDate, endDate, u.workdaysOnly))

		log.Println("Creating commit...")
		w.Commit("Random string", &git.CommitOptions{
			Author: &object.Signature{
				Name:  "John Doe",
				Email: "john@doe.org",
				When:  randate(startDate, endDate, u.workdaysOnly),
			}})
	}
}

func randate(startDate time.Time, endDate time.Time, workdaysOnly bool) time.Time {
	min := startDate.Unix()
	max := endDate.Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	var randomDate = time.Unix(sec, 0)

	if workdaysOnly {
		for !isWorkday(randomDate) {
			sec := rand.Int63n(delta) + min
			randomDate = time.Unix(sec, 0)
		}
	}

	return randomDate
}

func isWorkday(t time.Time) bool {
	return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
}
