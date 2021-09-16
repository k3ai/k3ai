package config



type K3ai struct {
// api: k3ai.in/v1alpha1
// kind: Application
// metadata:
//   name: dummy
//	 desc: "..."
//   version: v0.0.0
//   type: kustomize
// resources:
// # Pre Requisites
// - ../common/cert-manager/
// - ../common/istio/
// # Application logic
// - "http://a.b.c"
// - "http://d.e.f"
// # Post Requisites
// - overlays/ingress.yaml
// - overlays/rbac.yaml
	Api string `yaml:"api"`
	Kind string `yaml:"kind"`
	Metadata struct {
					Name string `yaml:"name"`
					Desc string `yaml:"desc"`
					Tag	string `yaml:"tag,omitempty"`
					Version string `yaml:"version"`
					Type string `yaml:"type"`
	}

	Resources []string `yaml:"resources,omitempty"`

}