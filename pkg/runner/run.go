package runner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	color "github.com/k3ai/pkg/color"
	db "github.com/k3ai/pkg/db"
	factory "github.com/k3ai/pkg/io/execution"
)

var template = `
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
        image: ghcr.io/k3ai/k3ai-executor:dev
        command: ["/bin/sleep", "3650d"]
EOF
`
var k3aiKube = "/.tools/kubectl"

func Loader(source string, target string, backend string, extras string, entrypoint string) error {
	var execTemplate string
	if entrypoint == "" {
		execTemplate = "\"/opt/k3ai-executor -b " + backend + " -s " + source + "\" "
		execTemplate = execTemplate + "\nEOF"
	}

	var execTemplateKfp = "\"/opt/k3ai-executor -b " + backend + " -s " + source + " -e " + entrypoint + "\" "
	execTemplateKfp = execTemplateKfp + `
EOF
`

	name, ctype := db.ListClusterByName(target)
	out := factory.Client(name, ctype)
	home, _ := os.UserHomeDir()
	shellPath := home + "/.k3ai"
	outcome, err := exec.Command("/bin/bash", "-c", "cat <<EOF | "+shellPath+k3aiKube+" apply  --kubeconfig="+out+" -f - "+template).Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" wait --for=condition=Ready pods --all -n default  --kubeconfig="+out).Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(outcome))

	time.Sleep(10 * time.Second)

	if backend == "mlflow" {
		_, err := exec.Command("/bin/bash", "-c", "cat <<EOF | "+shellPath+k3aiKube+"  --kubeconfig="+out+" exec  svc/minio-service -- bash -c \" mkdir /data/mlflow \"").Output()
		if err != nil {
		log.Println(err)
	}
	}
	_,err = exec.Command("/bin/bash", "-c", shellPath+k3aiKube+" wait --for=condition=Ready pods --all -n default  --kubeconfig="+out).Output()
	if err != nil {
		log.Println(err)
	}

	if backend == "mlflow" {
		cmd := exec.Command("/bin/bash", "-c", "cat <<EOF | "+shellPath+k3aiKube+"  --kubeconfig="+out+" exec  deployment/k3ai-executor -- bash -c "+execTemplate)
		if err != nil {
			log.Fatal(err)
		}
		cmd.Dir = shellPath
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})

		scanner := bufio.NewScanner(r)

		// loader.Working(msg)
		go func() {
			// Read line by line and process it
			msg := "🧪	Working, please wait..."
			fmt.Printf("\r %v", msg)
			fmt.Println(" ")
			for scanner.Scan() {
				line := scanner.Text()
				color.Disable()
				log.Println(line)
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
	if backend == "kfp" {
		cmd := exec.Command("/bin/bash", "-c", "cat <<EOF | "+shellPath+k3aiKube+"  --kubeconfig="+out+" exec  deployment/k3ai-executor -- bash -c "+execTemplateKfp)
		cmd.Dir = shellPath
		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})

		scanner := bufio.NewScanner(r)

		// loader.Working(msg)
		go func() {
			// Read line by line and process it
			msg := "🧪	Working, please wait..."
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
	return nil
}
