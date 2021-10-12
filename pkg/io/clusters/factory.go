package clusters

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	http "github.com/k3ai/pkg/http"
	internal "github.com/k3ai/internal"
)

const (
	k3aiKube ="/.k3ai/.tools/kubectl"
	k3aiHelm = "/.k3ai/.tools/helm"
	lnxApp = "/bin/bash"
	
)
var (
	appPlugin = internal.AppPlugin{}
	kubeconfig *string
) 


func Deployment (name string, ctype string) (status bool, err error) {

		appPlugin := http.InfrastructureDeployment(ctype)
		if appPlugin.Resources[0].PluginType == "shell" {
			_,err := exec.Command("/bin/bash","-c",appPlugin.Resources[0].Path).Output()
			if err != nil {
				os.Exit(0)
				log.Print(err)
			} 
		}


	// clientSet, kubeStr := Client(name, ctype)
	// WaitForDeployment(clientSet)
	status = true

	return status,nil
}

func WaitForDeployment(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		
}

func Client (name string,ctype string) (clientset  *kubernetes.Clientset, kubeStr []byte) { 
	var cPath string
	if ctype == "k3s" {
		cPath ="/etc/rancher/k3s/k3s.yaml"
	} else {
		cPath = homedir.HomeDir() + "/.kube/config"
	}
	if home := homedir.HomeDir(); home != "" {	
		out,_ := os.Create(homedir.HomeDir() + "/.k3ai/" + name +".config")
		in,_ := os.Open(cPath)
	
		
		_, err := io.Copy(out,in)
		if err != nil {
			log.Print(err)
		}
		out.Close()
		

		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".k3ai","johnny_cool.config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	 
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, kubeStr
}