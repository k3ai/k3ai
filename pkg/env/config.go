/*
TODO: Configuration initial sequence
BODY: We have to do the following actions:
		b. k3ai minimal set of tools to operate against kubernetes:
			- kubectl
			- helm
			- nerdctl (future)
			NOTE: specific cluster tools to be downloaded from the plugin directly
		c. k3ai plugin list to store in the database
			- We assume user is not authenticated so we will have to ask for a token (GH) that we have to store just for the operation.
				NOTE: to be checked if we need the token for the plugin installation (i.e.: bundles)

*/

package env

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"

	http "github.com/k3ai/pkg/http"
)

const (
	configPath    = "/.config/k3ai/"
	k3aiPath      = "/.k3ai"
	configUrl     = "https://raw.githubusercontent.com/k3ai/plugins/main/config.json"
	kubectlUrl    = "https://dl.k8s.io/release/v1.22.2/bin/linux/amd64/kubectl"
	kubectlSha256 = "https://dl.k8s.io/v1.22.2/bin/linux/amd64/kubectl.sha256" //nolint
	helmUrl       = "https://get.helm.sh/helm-v3.7.0-linux-amd64.tar.gz"
	nerdctl       = "" //nolint
	civoUrl       = "https://github.com/civo/cli/releases/download/v1.0.4/civo-1.0.4-linux-amd64.tar.gz"
)

/*
Check if a previous environment exist in both $HOME/.config/k3ai and $HOME/.k3ai
		a. If folders do not exist we have to create them
		b. If folder exist we try to read them, if error we will exit and inform the user
*/
func InitConfig(ch chan bool, msg string, sConfig string) {
	var homeDir, _ = os.UserHomeDir()
	if _, err := os.Stat(homeDir + configPath); !os.IsNotExist(err) {
		// if bForce {
		// 	os.RemoveAll(homeDir + configPath)
		// 	err := os.Mkdir(homeDir + configPath, 0755)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	Config()
		// }
		Config()
	} else if os.IsNotExist(err) {
		err := os.Mkdir(homeDir+configPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
		Config()
	} else {
		log.Fatal(err)

	}
	if _, err := os.Stat(homeDir + k3aiPath); !os.IsNotExist(err) {
		// if bForce {
		// 	os.RemoveAll(homeDir + k3aiPath)
		// 	err := os.Mkdir(homeDir + k3aiPath , 0755)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		kubectlConfig()
		civoConfig()
		helmConfig(ch)

	} else if os.IsNotExist(err) {
		err := os.Mkdir(homeDir+k3aiPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
		kubectlConfig()
		civoConfig()
		helmConfig(ch)
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		log.Fatal(err)

	}
	ch <- true
}

func Config() {
	var homeDir, _ = os.UserHomeDir()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(homeDir + "/.config/k3ai/")
	viper.AddConfigPath("~/.config/k3ai/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// err = viper.SafeWriteConfigAs(homeDir + "/.config/k3ai/config.yml")
			configData, err := http.Download(configUrl)
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile(homeDir+configPath+"/config.json", configData, 0775)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			log.Fatal(err)
		}
	}
}

func kubectlConfig() {
	homedir, _ := os.UserHomeDir()
	k3aiDir := homedir + "/.k3ai/.tools/"

	_, err := exec.Command("wget", kubectlUrl, "-P", k3aiDir).Output()
	if err != nil {
		log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
		os.Exit(0)
	}
	_, err = exec.Command("chmod", "+x", k3aiDir+"/kubectl").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func helmConfig(ch chan bool) {
	homedir, _ := os.UserHomeDir()
	k3aiDir := homedir + "/.k3ai/.tools/"

	_, err := exec.Command("wget", helmUrl, "-P", k3aiDir).Output()
	if err != nil {
		log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
		os.Exit(0)
	}
	_, err = exec.Command("tar", "-xvzf", k3aiDir+"/helm-v3.7.0-linux-amd64.tar.gz", "-C", k3aiDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "mv "+k3aiDir+"/linux-amd64/helm "+k3aiDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "rm "+k3aiDir+"/helm-v3.7.0-linux-amd64.tar.gz").Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "rm -r "+k3aiDir+"/linux-amd64").Output()
	if err != nil {
		log.Fatal(err)
	}
	ch <- true
}

func civoConfig() {
	homedir, _ := os.UserHomeDir()
	k3aiDir := homedir + "/.k3ai/.tools/"

	_, err := exec.Command("wget", civoUrl, "-P", k3aiDir).Output()
	if err != nil {
		log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
		os.Exit(0)
	}
	_, err = exec.Command("/bin/bash", "-c", "mkdir "+k3aiDir+"/civodir").Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("tar", "-xvzf", k3aiDir+"/civo-1.0.4-linux-amd64.tar.gz", "-C", k3aiDir+"/civodir").Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "mv "+k3aiDir+"civodir/civo "+k3aiDir).Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "rm "+"-r "+k3aiDir+"/civodir").Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("/bin/bash", "-c", "rm "+"-r "+k3aiDir+"/civo-1.0.4-linux-amd64.tar.gz").Output()
	if err != nil {
		log.Fatal(err)
	}
}
