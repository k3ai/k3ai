package internals

import (
	"fmt"
	"io/ioutil"
	"net/http"
	auth "github.com/alefesta/k3ai/shared"


)
func Init(){
	ctx,client,_ := auth.MainGitHub()

	// Let's read the folder structure of the base folder
	_,repos,apiRate,_:= client.Repositories.GetContents(ctx,"alefesta","k3ai-plugins","apps",nil)
   	fmt.Println(apiRate)
	for _,repo := range repos {
		if repo.GetType() == "dir" {
			repoName := "apps/" + repo.GetName()
			_,subRepos,_,_:= client.Repositories.GetContents(ctx,"kf5i","k3ai-plugins",repoName,nil)
			for _,subRepo := range subRepos {	
				if subRepo.GetType() == "dir" {
					subRepoName := repoName + "/" + subRepo.GetName()
					_, pluginContents,_,_ := client.Repositories.GetContents(ctx,"kf5i","k3ai-plugins",subRepoName,nil)
					for _,pluginContent := range pluginContents {
						url := pluginContent.GetDownloadURL()
						data,err := getContent(url)
						if err == nil {
							fmt.Println(string(data))
						}
					}

				}

				// fmt.Println(subRepo.GetName())
			}
		}
	}
}

func getContent(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}