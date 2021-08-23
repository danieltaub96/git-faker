package git

import (
	. "github.com/danieltaub96/git-faker/util/array"
	. "github.com/danieltaub96/git-faker/util/file"
	. "github.com/danieltaub96/git-faker/util/time"
	. "github.com/danieltaub96/git-faker/util/yaml"
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
func CheckGitInit(u UserInput) error {
	if _, err := os.Stat(path.Join(u.GitPath, ".git")); os.IsNotExist(err) {
		log.Errorln("Git is not init for the current directory")
		return err
	} else {
		log.Infof("Git path verified: %s\n", AbsPath(u.GitPath))
	}

	return nil
}

func CleanGitWorktree(u UserInput) {
	deleteCmd := exec.Command("rm", "-rf", u.GitPath)

	if _, err := deleteCmd.Output(); err != nil {
		log.Errorln("Unable to delete git worktree", err)
		return
	} else {
		log.Infoln("Cleaned git worktree")
	}

	gitCmd := exec.Command("git", "init", u.GitPath)

	if _, err := gitCmd.Output(); err != nil {
		log.Infoln("Unable to exec git init")
		return
	} else {
		log.Infoln("Git init in dir ", AbsPath(u.GitPath))
	}
}

func BuildCommitHistory(u UserInput, repo *git.Repository) error {
	var startDate = u.StartDate
	var endDate = u.EndDate

	w, _ := repo.Worktree()

	rand.Seed(time.Now().UnixNano())

	minCommits := u.CommitsPerDay.Start
	maxCommits := u.CommitsPerDay.End

	for i := 1; i <= u.Days; i++ {

		randomDate, randateErr := RandomDate(startDate, endDate, u.WorkdaysOnly)
		if randateErr != nil {
			log.Errorf("Error while try to generate random dates")
		}

		var commitsInDay = rand.Intn(maxCommits-minCommits) + minCommits

		for i := 1; i <= commitsInDay; i++ {

			randomDateWithHours, randHoursErr := RandomHours(randomDate, u.Hours)
			if randHoursErr != nil {
				log.Errorf("Error while try to generate random dates")
			}

			if err := commit(w, randomDateWithHours); err != nil {
				log.Errorln("Error while trying to create commit", err)
				return err
			}
		}
	}

	return nil
}

func commit(w *git.Worktree, when time.Time) error {
	var message = RandomValueFromArray(GetDataFile().Messages)
	var name = RandomValueFromArray(GetDataFile().Names)
	var email = RandomValueFromArray(GetDataFile().Emails)

	_, commitErr := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  when,
		}})

	log.Debugf("git commit with detailes: \nmessage: %s\nname: %s\nemail: %s\ntime: %s\n", message, name, email, when)

	return commitErr
}
