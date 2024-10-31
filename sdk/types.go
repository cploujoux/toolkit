package sdk

type Config struct {
	CurrentContext string            `yaml:"current_context"`
	Workspaces     []WorkspaceConfig `yaml:"workspaces"`
}
type WorkspaceConfig struct {
	Name        string      `yaml:"name"`
	Credentials Credentials `yaml:"credentials"`
}

type Credentials struct {
	APIKey       string `yaml:"api_key"`
	AccessToken  string `yaml:"access_token"`
	RefreshToken string `yaml:"refresh_token"`
	ExpiresIn    int    `yaml:"expires_in"`
	DeviceCode   string `yaml:"device_code"`
}
