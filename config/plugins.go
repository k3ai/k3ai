package config



type K3ai struct {
	Api string `yaml:"api"`
	Kind string `yaml:"kind"`
	Metadata struct {
					Name string `yaml:"name"`
					Desc string `yaml:"desc"`
					Tag	string `yaml:"tag,omitempty"`
					Version string `yaml:"version,omitempty"`
					PluginType string `yaml:"plugintype"`

	}

	Resources []string `yaml:"resources,omitempty"`

}

type Plugin struct {
	Api string `yaml:"api"`
	Kind string `yaml:"kind"`
	Resources []Resource
}

type Resource struct {
	Path string `yaml:"path"`
	Args string `yaml:"args,omitempty"`
	Kubecfg string `yaml:"kubeconfig,omitempty"`
	PluginType string `yaml:"type"`
	Wait bool `yaml:"wait"`
	Remove string `yaml:"remove,omitempty"`
}