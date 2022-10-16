package internal

import (
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"os"
)

func (g *GitRepo) HashOject(filepath string, write bool, objType string) error {
	fmt.Println("Hashing object with filepath:", filepath, "write:", write, "objType:", objType)

	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var obj GitObject
	switch objType {
	case "blob":
		obj = &BlobObject{
			data: fileData,
		}
	default:
		return fmt.Errorf("invalid object type")
	}
	data, sha := g.getWriteObject(obj)
	shaHexStr := hex.EncodeToString(sha)
	if write {
		objFolder := shaHexStr[:2]
		objFileName := shaHexStr[2:]
		objFilepath := g.repoPath("objects", objFolder, objFileName)
		_, err = createFile(objFilepath)
		if err != nil {
			return err
		}
		file, err := os.OpenFile(objFilepath, os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		writer := zlib.NewWriter(file)
		_, err = writer.Write(data)
		if err != nil {
			return err
		}
		defer writer.Close()

		fmt.Println("Written object:", shaHexStr)
	} else {
		fmt.Println("Object hash:", shaHexStr)
		fmt.Println("Data:", string(data))
	}
	return nil
}
