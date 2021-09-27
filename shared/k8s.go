package shared

import (
	"github.com/alefesta/k3ai/log"
	"fmt"
	"time"
	"context"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/janeczku/go-spinner"

)


func InitK8s(kubeconfig string, pluginName string) error {
	log.Info("Cluster up, waiting for remaining services to start...")
	time.Sleep(1 * time.Second)
	s := spinner.StartNew("Checking, please wait...")
	s.Start()
	s.SetSpeed(100 * time.Millisecond)
	s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
	if kubeconfig == "" {
		kubeconfig = "/home/alefesta/.kube/config"
	}
	if pluginName == "k3s" {
		kubeconfig = "/etc/rancher/k3s/k3s.yaml"
	}
	
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(40 * time.Second)
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	time.Sleep(5 * time.Second) // something more productive here
	if pods.Items[0].Status.Phase == "Running"{
		fmt.Println("Done")
		
	}
	s.Stop()
	log.Info("Services up and running...")
return nil
}


