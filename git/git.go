package git

import (
	util "github.com/danieltaub96/git-faker/util/time"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"time"
)
import . "github.com/danieltaub96/git-faker/object"

func OpenGitRepo(u UserInput) (*git.Repository, error) {
	repo, err := git.PlainOpen(u.GitPath)

	return repo, err
}
func CheckGitInit(u UserInput) {
	if _, err := os.Stat(path.Join(u.GitPath, ".git")); os.IsNotExist(err) {
		print("Git is not init for the current directory")
	}
}

func CleanGitWorktree(u UserInput) {
	deleteCmd := exec.Command("rm", "-rf", u.GitPath)
	if _, err := deleteCmd.Output(); err != nil {
		log.Infoln("Unable to delete git worktree")
		return
	}

	gitCmd := exec.Command("git", "init", u.GitPath)

	if _, err := gitCmd.Output(); err != nil {
		log.Infoln("Unable to exec git init")
		return
	}

	log.Infoln("Deleted successfully git dir")
}

func BuildCommitHistory(u UserInput, repo *git.Repository) {
	var startDate = u.StartDate
	var endDate = u.EndDate

	w, _ := repo.Worktree()

	rand.Seed(time.Now().UnixNano())

	minCommits := u.CommitsPerDay.Start
	maxCommits := u.CommitsPerDay.End

	var commitsInDay = rand.Intn(maxCommits-minCommits) + minCommits
	for i := 1; i <= commitsInDay; i++ {
		randomDate, randateErr := util.Randate(startDate, endDate, u.Hours, u.WorkdaysOnly)
		if randateErr != nil {
			log.Errorf("Error while try to generate random dates")
		}

		w.Commit("Random string", &git.CommitOptions{
			Author: &object.Signature{
				Name:  "John Doe",
				Email: "john@doe.org",
				When:  randomDate,
			}})
	}
}
