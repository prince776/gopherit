package main

import (
	"flag"
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
	gitClient, err := internal.NewGitClient(pwd, command)
	if err != nil {
		fmt.Println("Error: Couldn't start git client:", err)
		return
	}

	switch command {
	case internal.InitCmd:
		err = gitClient.Init()
	case internal.ValidateCmd:
		err = gitClient.Validate()
	case internal.CatFileCmd:
		err = gitClient.CatFile(os.Args[2])
	case internal.HashObjectCmd:
		os.Args = append(os.Args[0:1], os.Args[2:]...)
		write := flag.Bool("w", false, "write the hash to object store")
		objType := flag.String("t", "", "type of object to hash")
		flag.Parse()
		filepath := flag.Arg(0)
		err = gitClient.HashOject(filepath, *write, *objType)
	default:
		fmt.Println("Error: Bad command")
	}
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
