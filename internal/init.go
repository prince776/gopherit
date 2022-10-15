package internal

import (
	"fmt"
	"os"
)

func (g *GitRepo) Init(path string) error {
	fmt.Println("Doing git init", path)
	if fileOrDirExists(g.gitDir) {
		return fmt.Errorf("git repo already exists (.git dir(or file) exists)")
	}

	createDir(g.repoPath("objects"))
	createDir(g.repoPath("branches"))
	createDir(g.repoPath("refs", "heads"))
	createDir(g.repoPath("refs", "tags"))

	os.WriteFile(g.repoPath("description"), []byte("Unnamed repository; edit this file 'description' to name the repository.\n"), os.ModePerm)
	os.WriteFile(g.repoPath("HEAD"), []byte("ref: refs/heads/master\n"), os.ModePerm)

	return nil
}
