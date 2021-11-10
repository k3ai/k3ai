package internal

var (
	// Version is the version of the CLI injected in compilation time
	Version = "dev"
)

type Options struct {
	Quiet      bool
	PAT        string
	Config     string
	All        bool
	Force      bool
	Deploy     bool
	Remove     bool
	Name       string
	Type       string
	Target     string
	Filter     string
	Deployed   bool
	Source     string
	Backend    string
	Extras     string
	Entrypoint string
}

type Env struct {
	K3AI_TOKEN string `yaml:"K3AI_TOKEN"`
}

type K3aiPlugin struct {
	Api      string `yaml:"api"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name       string `yaml:"name"`
		Desc       string `yaml:"desc"`
		Tag        string `yaml:"tag,omitempty"`
		Version    string `yaml:"version,omitempty"`
		PluginType string `yaml:"plugintype"`
	}

	Resources []string `yaml:"resources,omitempty"`
}

type K3aiInternalPlugin struct {
	Api      string `yaml:"api"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name       string `yaml:"name"`
		Desc       string `yaml:"desc"`
		Tag        string `yaml:"tag,omitempty"`
		Version    string `yaml:"version,omitempty"`
		PluginType string `yaml:"plugintype"`
	}

	InternalResources []string `yaml:"resources,omitempty"`
}

type K3aiRootPlugin struct {
	Api      string `yaml:"api"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name         string `yaml:"name"`
		Desc         string `yaml:"desc"`
		Tag          string `yaml:"tag,omitempty"`
		Version      string `yaml:"version,omitempty"`
		PluginType   string `yaml:"plugintype"`
		PluginStatus string `yaml:"status"`
	}

	Resources []string `yaml:"resources,omitempty"`
}

type K3aiConfig struct {
	Api     string `yaml:"api"`
	Kind    string `yaml:"kind"`
	Cluster struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}
	Plugin struct {
		Name   string `yaml:"name"`
		Target string `yaml:"target"`
	}
	Run struct {
		Source  string `yaml:"source"`
		Target  string `yaml:"target"`
		Backend string `yaml:"backend"`
	}
}

type AppPlugin struct {
	Api       string `yaml:"api"`
	Kind      string `yaml:"kind"`
	Resources []AppPluginResources
}

type AppPluginResources struct {
	Path       string `yaml:"path"`
	Args       string `yaml:"args,omitempty"`
	Kubecfg    string `yaml:"kubeconfig,omitempty"`
	PluginType string `yaml:"type"`
	Wait       bool   `yaml:"wait"`
	Remove     string `yaml:"remove,omitempty"`
}
