package internal

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func (g *GitRepo) Init() error {
	if fileOrDirExists(g.gitDir) {
		return fmt.Errorf("git repo already exists (.git dir(or file) exists)")
	}

	createDir(g.repoPath("objects"))
	createDir(g.repoPath("branches"))
	createDir(g.repoPath("refs", "heads"))
	createDir(g.repoPath("refs", "tags"))

	os.WriteFile(g.repoPath("description"), []byte("Unnamed repository; edit this file 'description' to name the repository.\n"), os.ModePerm)
	os.WriteFile(g.repoPath("HEAD"), []byte("ref: refs/heads/master\n"), os.ModePerm)

	g.defaultConfigFile().SaveTo(g.repoPath("config"))
	return nil
}

func (g *GitRepo) defaultConfigFile() *ini.File {
	cfg := ini.Empty()
	cfg.NewSection("core")
	cfg.Section("core").NewKey("repositoryformatversion", "0")
	cfg.Section("core").NewKey("filemode", "false")
	cfg.Section("core").NewKey("bare", "false")

	return cfg
}
