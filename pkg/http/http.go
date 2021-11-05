package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/k3ai/pkg/auth"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	internal "github.com/k3ai/internal"
	db "github.com/k3ai/pkg/db"
)

var (
	repoPaths   = []string{"apps", "infra", "bundles"}
	dataResults = internal.K3aiRootPlugin{}
)

//Download read the remote file and retrieve its content
func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func RetrievePlugins(token string, action string, ch chan bool) {
	var baseUser string
	var baseRepo string
	var homeDir, _ = os.UserHomeDir()
	k3aiConfig := homeDir + "/.config/k3ai/"
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(k3aiConfig)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if viper.GetBool("plugins.community") {
		repoPaths = append(repoPaths, "community")
	}
	if viper.GetString("plugins.base_url") != "" {
		baseRepo = viper.GetString("plugins.baseRepo")
		baseUser = viper.GetString("default.baseUser")
	} else {
		baseRepo = viper.GetString("default.baseRepo")
		baseUser = viper.GetString("default.baseUser")
	}
	if viper.GetString("default.cluster_types") != "all" {
		log.Print("clusters")
	}
	//let's find owner, base repo and add the sub-folders
	client, _, ctx := auth.GitHub(token)
	for i := 0; i < len(repoPaths); i++ {
		_, reposUrl, _, _ := client.Repositories.GetContents(ctx, baseUser, baseRepo, repoPaths[i], nil)
		for _, repoUrl := range reposUrl {
			if repoUrl.GetType() == "dir" && repoUrl.GetName() != "template" {
				subPath := repoUrl.GetPath()
				_, subContents, _, _ := client.Repositories.GetContents(ctx, baseUser, baseRepo, subPath, nil)
				for _, subContent := range subContents {
					if subContent.GetType() == "file" && subContent.GetName() == "k3ai.yaml" {
						url := subContent.GetDownloadURL()
						data, _ := Download(url)

						err := yaml.Unmarshal([]byte(data), &dataResults)
						if err != nil {
							fmt.Print(err)
						}
						plugins := [...]string{dataResults.Metadata.Name, dataResults.Metadata.Desc, dataResults.Kind, dataResults.Metadata.Tag, dataResults.Metadata.Version, subContent.GetDownloadURL(), dataResults.Metadata.PluginStatus}
						if action == "config" {
							err = db.InsertPlugins(plugins)

						} else if action == "update" {
							err = db.UpdatePlugins(plugins)

						}

						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
	ch <- true
}

func GetIP() string {
	var ipnet *net.IPNet
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ipnet.IP.String()
}
