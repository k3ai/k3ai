package shared

import (
	"context"
	"github.com/alefesta/k3ai/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"

)


func InitK8s() error {
	kubeconfig := "/home/alefesta/.kube/config"
	config,err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Info(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	_, err = clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Error("Sorry I do not find any active cluster, is it one configured or did you run apply any K3ai infrastructure plugin first?")
		return err
	}
	return nil
}