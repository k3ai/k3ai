package shared

import (
	// "github.com/alefesta/k3ai/log"
	"fmt"
	"time"
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

)


func InitK8s(kubeconfig string, pluginName string) error {
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
pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
if err != nil {
	panic(err.Error())
}
//# @to-do	Check if K8s is up and running
//# @body	Wait to see if the cluster is available
// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

// // Examples for error handling:
// // - Use helper functions like e.g. errors.IsNotFound()
// // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
// namespace := "kube-system"
// pod := "core-dns"
// _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
// if errors.IsNotFound(err) {
// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
// 		pod, namespace, statusError.ErrStatus.Message)
// } else if err != nil {
// 	panic(err.Error())
// } else {
// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
// }

// time.Sleep(10 * time.Second)
return nil
}
