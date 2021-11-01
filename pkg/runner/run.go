package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
  "bufio"

	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	factory "github.com/k3ai/pkg/io/execution"
  // loader "github.com/k3ai/pkg/loader"
)

var template =`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: k3ai-executor
  labels:
    app: k3ai
spec:
  replicas: 1
  selector:
    matchLabels:
      app: k3ai
  template:
    metadata:
      labels:
        app: k3ai
    spec:
      containers:
      - name: k3ai
        image: ghcr.io/k3ai/k3ai-executor:latest
        command: ["/bin/sleep", "3650d"]
EOF
`
var k3aiKube ="/.tools/kubectl"


func Loader(source string,target string,backend string, extras string) error {
 var execTemplate ="\". ./run.sh -b " + backend + " -s " + source + "\" "
 execTemplate = execTemplate + `
EOF
`
log.Printf(execTemplate)

  name,ctype:= db.ListClusterByName(target)
  out := factory.Client(name,ctype)
	home,_ := os.UserHomeDir()
	shellPath := home + "/.k3ai"
  color.InProgress()
  fmt.Println(" 🚀 Working on the installation...")	
  outcome,err := exec.Command("/bin/bash","-c", "cat <<EOF | " + shellPath + k3aiKube + " apply  --kubeconfig="+ out +" -f - " + template ).Output()
  if err != nil {
    log.Println(err)
  }
  _,_ = exec.Command("/bin/bash","-c", shellPath + k3aiKube + " wait --for=condition=Ready pods --all --all-namespaces  --kubeconfig="+ out).Output()
  fmt.Println(string(outcome))



  if backend == "mlflow" {
    cmd:= exec.Command("/bin/bash","-c",  "cat <<EOF | " + shellPath + k3aiKube + "  --kubeconfig="+ out + " exec  deployment/k3ai-executor -- bash -c " + execTemplate )
    log.Println(cmd)
    cmd.Dir = shellPath
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
    done := make(chan struct{})

		scanner := bufio.NewScanner(r)

    // loader.Working(msg)
    go func() {
			// Read line by line and process it
      msg := "⏳	Working, please wait..."
      fmt.Printf("\r %v", msg)
			fmt.Println(" ")
			for scanner.Scan() {
				scanner.Text()
				color.Disable()
			}
			done <- struct{}{}
		}()
		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			log.Println("Something went wrong... did you check all the prerequisites to run this plugin? If so try to re-run the k3ai command...")	
		}
		<-done
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
  }
  _,_ = exec.Command("/bin/bash","-c", "cat <<EOF | " + shellPath + k3aiKube + " delete  --kubeconfig="+ out +" -f - " + template ).Output()
  if err != nil {
    log.Println(err)
  }
  return nil
}