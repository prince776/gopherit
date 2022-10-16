package internal

import (
	"fmt"
	"path/filepath"
)

type GitClient interface {
	Init() error
	Validate() error
	CatFile(object string) error
}

type GitRepo struct {
	worktree string
	gitDir   string
	confFile string
}

func NewGitClient(path string, cmd string) (GitClient, error) {
	fmt.Println("Git client initiated at:", path, " with command: ", cmd)
	if !dirExists(path) {
		return nil, fmt.Errorf("%v directory doesn't exist", path)
	}

	if cmd != InitCmd {
		var err error
		path, err = findGitPath(path)
		if err != nil {
			return nil, err
		}
		path = filepath.Clean(filepath.Join(path, ".."))
	}
	repo := &GitRepo{
		worktree: path,
		gitDir:   filepath.Join(path, ".git"),
	}
	repo.confFile = repo.repoPath("config")
	_, err := repo.validateRepo()
	return repo, err
}

// get path of file relative to gitDir
func (g *GitRepo) repoPath(paths ...string) string {
	paths = append([]string{g.gitDir}, paths...)
	return filepath.Join(paths...)
}
