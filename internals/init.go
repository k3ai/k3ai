package internals

import (
	"os"
	"time"

	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
	"github.com/spf13/viper"
	"github.com/google/go-github/v39/github"
	"github.com/briandowns/spinner"
	_ "github.com/mattn/go-sqlite3"
	data "github.com/k3ai/config"
	log "github.com/k3ai/log"
	utils "github.com/k3ai/shared"
	
)

const (
	repoOwner = "k3ai"
	repoRoot = "plugins"
	repoApps = "apps"
	repoComms = "community"
	repoInfra = "infra"
	repoBundle = "bundles"
	homeK3ai = ".k3ai"
	k3aiDb = "k3ai.db"
	k3aiFile = "k3ai.yaml"
)

var dataResults = data.K3ai{}
// var db *sql.DB
//Init initialize the k3ai tool. 
func Init(){
	icon := []string{"⛏️ "}
	s := spinner.New(icon, 100*time.Millisecond)
	// s.Color("green")
	s.Start()
	time.Sleep(500 * time.Millisecond)
	var action = "create"
	homeDir,_ := os.UserHomeDir()
	err := mkDir()
	_ = log.CheckErrors(err)
	log.Info("Initialize K3ai...")	
	// s.Prefix = "Initialize K3ai:"
	time.Sleep(500 * time.Millisecond)
	log.Info("Creating k3ai folder structure...")
	time.Sleep(500 * time.Millisecond)
	log.Info("Done | Created .k3ai folder at: " + homeDir + "/" + homeK3ai)
	time.Sleep(500 * time.Millisecond)
	log.Info("Setting up local database...")
	_,err = utils.DbCreate()
	_ = log.CheckErrors(err)
	time.Sleep(500 * time.Millisecond)
	log.Info("Done | K3ai DataBase created...")
	data.InitEnv()
	viper.AutomaticEnv()
	time.Sleep(500 * time.Millisecond)
	log.Info("Synchronizing plugin list...")
	err = pluginContent(action)
	log.CheckErrors(err)
	log.Info("Done | Plugins synchronized")
	s.Stop()
}

func Update(){
		var action = "update"
		log.Info("Updating K3ai plugin list...")
		time.Sleep(500 * time.Millisecond)
		err := pluginContent(action)
		_ = log.CheckErrors(err)
		log.Info("Done | Plugins synchronized")
}


//mkDir create a local directory under user home folder
func mkDir() error {
	homeDir,_ := os.UserHomeDir()
	if _, err := os.Stat(homeDir + "/" + homeK3ai); os.IsNotExist(err) {
		//Create a folder/directory at a full qualified path
		err := os.Mkdir(homeDir + "/" + homeK3ai, 0755)
		_ = log.CheckErrors(err)
	}
	return nil
}


//Read the current plugin details
 func pluginContent (action string) error {
	ctx,client,_ := utils.MainGitHub()
	var reposComms []*github.RepositoryContent
	// Let's retrieve the list of various plugins and store them as a
	_,reposApps,_,_:= client.Repositories.GetContents(ctx,repoOwner,repoRoot,repoApps,nil)
	_,reposInfra,_,_:= client.Repositories.GetContents(ctx,repoOwner,repoRoot,repoInfra,nil)
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.k3ai")  // call multiple times to add many search paths
	viper.ReadInConfig()
	if viper.GetBool("COMMUNITY") {
		_,reposComms,_,_= client.Repositories.GetContents(ctx,repoOwner,repoRoot,repoComms,nil)
	}

	_,reposBundles,_,_:= client.Repositories.GetContents(ctx,repoOwner,repoRoot,repoBundle,nil)

	for _,repoApp := range reposApps {
			if repoApp.GetType() == "dir"  && repoApp.GetName() != "template" {
				subRoot := repoApp.GetPath()
				_,subContents,_,_ := client.Repositories.GetContents(ctx,repoOwner,repoRoot,subRoot,nil)
				for _,subContent := range subContents {
					if subContent.GetType() == "file" && subContent.GetName() == k3aiFile {
						url := subContent.GetDownloadURL()
						data,_ := getContent(url)
					
						err := yaml.Unmarshal([]byte(data), &dataResults)
						if err != nil {
							log.Error(err)
						}
						if action == "create" {
							utils.FillPluginTables(&dataResults, subContent.GetDownloadURL(),"")
							} else if action == "update" {
								utils.UpdatePluginTables(&dataResults, subContent.GetDownloadURL(),"")
						}
						
					}
					
				}
			}
	 }

	 for _,repoInfra := range reposInfra {
		if repoInfra.GetType() == "dir" && repoInfra.GetName() != "template" {
			subRoot := repoInfra.GetPath()
			_,subContents,_,_ := client.Repositories.GetContents(ctx,repoOwner,repoRoot,subRoot,nil)
			for _,subContent := range subContents {
				if subContent.GetType() == "file" && subContent.GetName() == k3aiFile {
					url := subContent.GetDownloadURL()
					data,_ := getContent(url)
				
					err := yaml.Unmarshal([]byte(data), &dataResults)
					if err != nil {
						log.Error(err)
					}

					if action == "create" {
						utils.FillPluginTables(&dataResults, subContent.GetDownloadURL(),"")
						} else if action == "update" {
							utils.UpdatePluginTables(&dataResults, subContent.GetDownloadURL(),"")
					}
				}
				
			}
		}
 }

 for _,repoBundle := range reposBundles {
	if repoBundle.GetType() == "dir" && repoBundle.GetName() != "template"  {
		subRoot := repoBundle.GetPath()
		_,subContents,_,_ := client.Repositories.GetContents(ctx,repoOwner,repoRoot,subRoot,nil)
		for _,subContent := range subContents {
			if subContent.GetType() == "file" && subContent.GetName() == k3aiFile {
				url := subContent.GetDownloadURL()
				data,_ := getContent(url)
			
				err := yaml.Unmarshal([]byte(data), &dataResults)
				if err != nil {
					log.Error(err)
				}

				if action == "create" {
					utils.FillPluginTables(&dataResults, subContent.GetDownloadURL(),"")
				} else if action == "update" {
					utils.UpdatePluginTables(&dataResults, subContent.GetDownloadURL(),"")
				}
			}
			
		}
	}
}

for _,repoComm := range reposComms {
	if repoComm.GetType() == "dir" && repoComm.GetName() != "template"  {
		subRoot := repoComm.GetPath()
		_,subContents,_,_ := client.Repositories.GetContents(ctx,repoOwner,repoRoot,subRoot,nil)
		for _,subContent := range subContents {
			if subContent.GetType() == "file" && subContent.GetName() == k3aiFile {
				url := subContent.GetDownloadURL()
				data,_ := getContent(url)
			
				err := yaml.Unmarshal([]byte(data), &dataResults)
				if err != nil {
					log.Error(err)
				}

				if action == "create" {
					utils.FillPluginTables(&dataResults, subContent.GetDownloadURL(),"comm")
				} else if action == "update" {
					utils.UpdatePluginTables(&dataResults, subContent.GetDownloadURL(),"comm")
				}
			}
			
		}
	}
}
	 return nil
 }


//getContent read the remote file and retrieve its content
func getContent(url string) ([]byte, error) {
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