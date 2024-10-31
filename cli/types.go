package cli

type ResultMetadata struct {
	Workspace string
	Name      string
}

type Result struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   ResultMetadata
	Spec       interface{}
}
