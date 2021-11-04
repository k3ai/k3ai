package git

import (
	"os"
	"fmt"
	"strings"

	git "github.com/go-git/go-git/v5"
)


func Clone(cloneUrl string, name string) error {
	homePath,_ := os.UserHomeDir()
	repository := cloneUrl
	name = strings.ToLower(name)
	_, err := git.PlainClone(homePath+"/.k3ai/git/",false,&git.CloneOptions{
		URL: repository,
	})
	if err != nil && err.Error() != "repository already exists" {
			fmt.Printf("%v", err)
			return err
	}
	// fmt.Println("Repository cloned")
	return nil
}