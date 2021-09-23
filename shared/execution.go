package shared

import (
	// "os"
	"os/exec"
	// "io/ioutil"
	// config "github.com/alefesta/k3ai/config"
	"github.com/alefesta/k3ai/log"
)

func InitExec(pluginName string ,pluginEx string ,pluginArgs string,pluginKube string,pluginType string,pluginWait bool) {
	
	if pluginType == "shell" {
		_,err := exec.Command("/bin/bash","-c",pluginEx,pluginArgs).Output()
		if err != nil {
			log.Error("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
			// os.Exit(0)
			
		}
		if pluginWait {
			err = InitK8s(pluginKube, pluginName)
			if err != nil {
				log.Error(err)
			}
			}
		}


	}