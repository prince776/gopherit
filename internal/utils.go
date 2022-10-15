package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func dirExists(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) && stat.IsDir() {
		return true
	}
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileOrDirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func createDir(paths ...string) (string, error) {
	path := filepath.Join(paths...)
	if fileExists(path) {
		return path, fmt.Errorf("file %v exists, can't make directory", path)
	}
	os.MkdirAll(path, os.ModePerm)
	return path, nil
}

func createFile(paths ...string) (string, error) {
	path := filepath.Join(paths...)
	if len(paths) > 1 {
		_, err := createDir(paths[0 : len(paths)-1]...)
		if err != nil {
			return path, err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return path, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()
	return path, nil
}
