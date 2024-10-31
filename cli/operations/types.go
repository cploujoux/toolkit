package operations

type ResultMetadata struct {
	Workspace string
	Name      string
}

type Result struct {
	Kind     string
	Metadata ResultMetadata
	Spec     interface{}
}
