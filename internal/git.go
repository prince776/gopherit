package internal

import (
	"fmt"
	"path/filepath"
)

type GitClient interface {
	Init() error
	Validate() error
}

type GitRepo struct {
	worktree string
	gitDir   string
	confFile string
}

func NewGitClient(path string, cmd string) (GitClient, error) {
	if !dirExists(path) {
		return nil, fmt.Errorf("%v directory doesn't exist", path)
	}

	if cmd != InitCmd {
		var err error
		path, err = findGitPath(path)
		if err != nil {
			return nil, err
		}
	}
	repo := &GitRepo{
		worktree: path,
		gitDir:   filepath.Join(path, ".git"),
	}
	repo.confFile = repo.repoPath("config")
	return repo, nil
}

// get path of file relative to gitDir
func (g *GitRepo) repoPath(paths ...string) string {
	paths = append([]string{g.gitDir}, paths...)
	return filepath.Join(paths...)
}
