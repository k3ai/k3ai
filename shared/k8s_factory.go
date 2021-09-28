package shared


type factory struct {
	KubeConfig            string
	Context               string

}

// GetK3sKubeConfig returns the kubeconfig path
func GetK3sKubeConfig() string {
	return "/etc/rancher/k3s/k3s.yaml"
}

// GetK0sKubeConfig returns the kubeconfig path
func GetK0sKubeConfig() string {
	return "/var/lib/k0s/pki/admin.conf"
}