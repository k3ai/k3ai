package shared

import (

	"time"
	"context"
	"path/filepath"

	"github.com/pkg/errors"
	log "github.com/k3ai/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// auth providers
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"


)


// NewClientConfig returns a new Kubernetes client config set for a context
func NewClientConfig(configPath string, contextName string) clientcmd.ClientConfig {
	configPathList := filepath.SplitList(configPath)
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{}
	if len(configPathList) <= 1 {
		configLoadingRules.ExplicitPath = configPath
	} else {
		configLoadingRules.Precedence = configPathList
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		configLoadingRules,
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		},
	)
}

// NewClientSet returns a new Kubernetes client for a client config
func NewClientSet(clientConfig clientcmd.ClientConfig) (*kubernetes.Clientset, error) {
	c, err := clientConfig.ClientConfig()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get client config")
	}

	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create clientset")
	}

	return clientset, nil
}

// GetK3sClientSet returns the kubernetes Clientset
func GetK3sClientSet() (*kubernetes.Clientset, error) {
	clientConfig := NewClientConfig(GetK3sKubeConfig(), "default")
	return NewClientSet(clientConfig)

}

// GetK0sClientSet returns the kubernetes Clientset
func GetK0sClientSet() (*kubernetes.Clientset, error) {
	clientConfig := NewClientConfig(GetK0sKubeConfig(), "default")
	return NewClientSet(clientConfig)

}
func InitK8s(kubeconfig string,pluginName string) error {
	clientset,_ := GetK3sClientSet()


	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{} )
	for i:=0 ; i < len(pods.Items); i++ {
		if pods.Items[i].Status.Phase != "Running" {
			// log.Info("Waiting for " + pods.Items[i].Name + " to start...")
			time.Sleep(5 * time.Second)
		} 

	}
	if err != nil {
		log.Error(err)
	}	

return nil
}

