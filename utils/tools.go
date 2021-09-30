package utils

const (
	kubectl = "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
	kubectx = "https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubectx"
	helm = "https://get.helm.sh/helm-v3.7.0-linux-amd64.tar.gz"

)

//TODO : Get K3ai tools
//Body: We need to download the basic tools in the k3ai folder
// this way we will not have to relay on user enviroment
// ideally user may run them as "standlone" with `k3ai <tool>` i.e. (`k3ai kubectl get po -n default`)