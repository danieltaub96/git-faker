package git

import (
	. "github.com/danieltaub96/git-faker/object"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OpenGitRepoDosentInit(t *testing.T) {
	_, err := OpenGitRepo(UserInput{GitPath: t.TempDir()})
	assert.Error(t, err, "Error in testing OpenGitRepo")
}

func Test_OpenGitRepoOk(t *testing.T) {
	var tempDir = t.TempDir()
	git.PlainInit(tempDir, true)

	_, err := OpenGitRepo(UserInput{GitPath: tempDir})
	assert.NoError(t, err, "Error in testing OpenGitRepo")
}

func Test_OpenGitRepoDirNotExists(t *testing.T) {
	_, err := OpenGitRepo(UserInput{GitPath: "dir_not_exists"})
	assert.Error(t, err, "Error in testing OpenGitRepo")

}

func Test_CleanGitWorktreeOk(t *testing.T) {
}
