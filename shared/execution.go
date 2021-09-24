package shared

import (
	// "os"
	"bufio"
	"os/exec"
	// "io/ioutil"
	// config "github.com/alefesta/k3ai/config"
	"github.com/alefesta/k3ai/log"
)

func InitExec(pluginName string ,pluginEx string ,pluginArgs string,pluginKube string,pluginType string,pluginWait bool) {
	
	if pluginType == "shell" {
		cmd := exec.Command("/bin/bash","-c",pluginEx,pluginArgs)
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})
		scanner := bufio.NewScanner(r)
		go func() {

			// Read line by line and process it
			for scanner.Scan() {
				line := scanner.Text()
				log.Info(line)
			}
			done <- struct{}{}

		}()
		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			log.Error("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
			// os.Exit(0)	
		}
		<-done
		err = cmd.Wait()
		log.Error(err)
		if pluginWait {
			err = InitK8s(pluginKube, pluginName)
			if err != nil {
				log.Error(err)
			}
		}

	}
}

func InitRemove(pluginName string ,pluginEx string ,pluginArgs string,pluginKube string,pluginType string,pluginWait bool, pluginRemove string) {
	
	if pluginType == "shell" {
		cmd := exec.Command("/bin/bash","-c",pluginRemove,pluginArgs)
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})
		scanner := bufio.NewScanner(r)
		go func() {

			// Read line by line and process it
			for scanner.Scan() {
				line := scanner.Text()
				log.Info(line)
			}
			done <- struct{}{}

		}()
		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			log.Error("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")
			// os.Exit(0)	
		}
		<-done
		err = cmd.Wait()
		log.Error(err)
		// if pluginWait {
		// 	err = InitK8s(pluginKube, pluginName)
		// 	if err != nil {
		// 		log.Error(err)
		// 	}
		// }

	}


}