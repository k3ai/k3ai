package git

import (
	"fmt"

	git "github.com/go-git/go-git/v5"

)


func Clone(cloneUrl string) error {
	repository := cloneUrl
	_, err := git.PlainClone("/home/alefesta/.k3ai/git/katib",false,&git.CloneOptions{
		URL: repository,
	})
	if err != nil && err.Error() != "repository already exists" {
			fmt.Printf("%v", err)
			return err
	}
	fmt.Println("Repository cloned")
	return nil
}