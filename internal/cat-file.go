package internal

import "fmt"

func (g *GitRepo) CatFile(objectHash string) error {
	obj, err := g.readObject(objectHash)
	if err != nil {
		return err
	}
	data := obj.Serialize()
	fmt.Println("Object meta:", string(obj.Format()), len(data))
	fmt.Println("Object data:")
	fmt.Println(string(data))
	return nil
}
