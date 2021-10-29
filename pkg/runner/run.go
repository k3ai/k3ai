package runner

import "log"

var template = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: k3ai-executor
spec:
  selector:
    matchLabels:
      app: k3ai-executor
  replicas: 1
  template:
    metadata:
      labels:
        app: k3ai-executor
    spec:
      containers:
      - name: k3ai-executor
        image: ghcr.io/k3ai/k3ai-executor:latest
		command: ["/bin/sleep","3650d"]
`



func Loader() {
	log.Printf(template)
}