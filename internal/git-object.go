package internal

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
)

type GitObject interface {
	Serialize() []byte
	Deserialize(data []byte)
	Format() []byte
}

func (g *GitRepo) readObject(sha string) (GitObject, error) {
	path := g.repoPath("objects", sha[:2], sha[2:])
	fmt.Println("path is:", path)
	if !fileExists(path) {
		return nil, fmt.Errorf("object doesn't exist")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader, err := zlib.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	rawData := buf.Bytes()
	spacePos := bytes.Index(rawData, []byte(" "))
	nullBytePos := bytes.Index(rawData, []byte{byte(0x00)})
	format := rawData[:spacePos]
	size, err := strconv.Atoi(string(rawData[spacePos+1 : nullBytePos]))
	if err != nil {
		return nil, err
	}

	data := rawData[nullBytePos+1:]
	if size != len(data) {
		return nil, fmt.Errorf("malformed object")
	}
	var res GitObject
	switch string(format) {
	case "blob":
		res = &BlobObject{
			data: data,
		}
	default:
		return nil, fmt.Errorf("unsupported format: %v", string(format))
	}
	return res, nil
}

func (g *GitRepo) getWriteObject(obj GitObject) (finalData []byte, sha []byte) {
	data := obj.Serialize()
	finalData = append(obj.Format(), []byte(" ")...)
	finalData = append(finalData, []byte(strconv.Itoa(len(data)))...)
	finalData = append(finalData, []byte{byte(0x00)}...)
	finalData = append(finalData, data...)

	hasher := sha1.New()
	hasher.Write(finalData)
	sha = hasher.Sum(nil)
	return
}

type BlobObject struct {
	data []byte
}

func (o *BlobObject) Format() []byte {
	return []byte("blob")
}

func (o *BlobObject) Serialize() []byte {
	return o.data
}

func (o *BlobObject) Deserialize(data []byte) {
	o.data = data
}
