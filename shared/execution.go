package shared

import (
	// "os"
	"os/exec"
	// "io/ioutil"
	// config "github.com/alefesta/k3ai/config"
	"github.com/alefesta/k3ai/log"
)

func InitExec(pluginName string ,pluginEx string ,pluginArgs string,pluginKube string,pluginType string,pluginWait bool) {
	log.Info("Starting to install " + pluginName)
	if pluginType == "shell" {
		_,err := exec.Command("/bin/bash","-c",pluginEx,pluginArgs).Output()
		if err == nil {
			log.Error(err)
		} else {
			InitK8s(pluginKube, pluginName)
		}

		}

	}