package main

import (
	"fmt"
	"os"

	"github.com/prince776/gopherit/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need a command: git <COMMAND>, ex: git commit")
		return
	}
	command := os.Args[1]

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gitClient, err := internal.NewGitClient(pwd)
	if err != nil {
		fmt.Println("Error: Couldn't start git client:", err)
		return
	}

	switch command {
	case internal.InitCmd:
		err = gitClient.Init()
	case internal.ValidateCmd:
		err = gitClient.Validate()
	default:
		fmt.Println("Error: Bad command")
	}
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
