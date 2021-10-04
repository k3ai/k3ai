package utils

import (
	"bufio"
	"os"
	"os/exec"

	log "github.com/k3ai/log"
)

const (
	kubectl = "https://dl.k8s.io/release/v1.22.2/bin/linux/amd64/kubectl"
	// kubectx = "https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubectx" // kubectx currently do not support kubectl binary not in PATH
	helm = "https://get.helm.sh/helm-v3.7.0-linux-amd64.tar.gz"

)

var tools = []string{kubectl,helm}

//GetTools download the tools from the constant and save them inside the .k3ai folder
func GetTools() {
	homedir,_ := os.UserHomeDir()
	k3aiDir := homedir + "/.k3ai/.utils/"
	if _, err := os.Stat(k3aiDir); os.IsNotExist(err)  {
		err := os.Mkdir(k3aiDir, 0755)
		_ = log.CheckErrors(err)
	}
	for _,tool :=range tools {
		cmd := exec.Command("wget",tool,"-P",k3aiDir)
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})
		scanner := bufio.NewScanner(r)
		go func() {
	
			// Read line by line and process it
			for scanner.Scan() {
				scanner.Text()
				// log.Info(line)
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
		cmd.Wait()
	}
	exec.Command("chmod","+x",k3aiDir + "/kubectl").Output()
	// exec.Command("chmod","+x",k3aiDir + "/kubectx").Output()
	_,err := exec.Command("tar","-xvzf",k3aiDir + "/helm-v3.7.0-linux-amd64.tar.gz","-C",k3aiDir).Output()
	if err != nil {
		log.Info(err)
	}
	_,err = exec.Command("/bin/bash","-c","mv " + k3aiDir + "/linux-amd64/helm " + k3aiDir).Output()
	if err != nil {
		log.Info(err)
	}
	_,err = exec.Command("/bin/bash","-c","rm " + k3aiDir + "/helm-v3.7.0-linux-amd64.tar.gz").Output()
	if err != nil {
		log.Info(err)
	}
	_,err = exec.Command("/bin/bash","-c","rm -r " + k3aiDir + "/linux-amd64").Output()
	if err != nil {
		log.Info(err)
	}

	




}