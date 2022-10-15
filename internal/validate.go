package internal

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func (g *GitRepo) Validate() error {
	ok, err := g.validateRepo()
	if err != nil {
		return err
	}
	if ok {
		fmt.Println("Repo is valid")
	} else {
		fmt.Println("Not a valid repo")
	}
	return nil
}

func (g *GitRepo) validateRepo() (bool, error) {
	if !dirExists(g.gitDir) {
		return false, fmt.Errorf("not a git repository: .git directory doesn't exist")
	}
	if !fileExists(g.confFile) {
		return false, fmt.Errorf("missing configuration file")
	}
	cfg, err := ini.Load(g.confFile)
	if err != nil {
		return false, fmt.Errorf("error loading configuration file %v", err)
	}
	if formatVersion := cfg.Section("core").Key("repositoryformatversion").String(); formatVersion != "0" {
		return false, fmt.Errorf("unsupported repository format version %v", formatVersion)
	}
	return true, nil
}
