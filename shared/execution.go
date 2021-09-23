package shared

import (
	"os"
	"os/exec"
	// "io/ioutil"
	// config "github.com/alefesta/k3ai/config"
	"github.com/alefesta/k3ai/log"
)

func InitExec(pluginName string ,pluginEx string ,pluginArgs string,pluginKube string,pluginType string,pluginWait bool) {
	log.Info("Starting to install " + pluginName)
	if pluginType == "shell" {
		_,err := exec.Command("/bin/bash","-c",pluginEx,pluginArgs).Output()
		if err != nil {
			os.Exit(0)
			log.Error(err)
		} 
		err = InitK8s(pluginKube, pluginName)
			if err != nil {
				log.Error(err)
			}
		}

	}