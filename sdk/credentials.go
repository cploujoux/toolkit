package sdk

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func loadConfig() Config {
	config := Config{}
	homeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(homeDir, ".beamlit", "config.yaml")
		if data, err := os.ReadFile(configPath); err == nil {
			if err := yaml.Unmarshal(data, &config); err != nil {
				// Invalid YAML, use empty credentials
				config = Config{}
			}
		}
	}
	return config
}

func saveConfig(config Config) {
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	configDir := filepath.Join(homeDir, ".beamlit")
	configFile := filepath.Join(configDir, "config.yaml")
	if err := os.WriteFile(configFile, yamlData, 0600); err != nil {
		fmt.Printf("Error writing credentials file: %v\n", err)
		os.Exit(1)
	}
}

func CurrentContext() string {
	config := loadConfig()
	return config.CurrentContext
}

func SetCurrentContext(workspaceName string) {
	config := loadConfig()
	config.CurrentContext = workspaceName
	saveConfig(config)
}

func LoadCredentials(workspaceName string) Credentials {
	config := loadConfig()
	for _, workspace := range config.Workspaces {
		if workspace.Name == workspaceName {
			return workspace.Credentials
		}
	}
	return Credentials{}
}

func SaveCredentials(workspaceName string, credentials Credentials) {
	config := loadConfig()
	found := false
	for i, workspace := range config.Workspaces {
		if workspace.Name == workspaceName {
			config.Workspaces[i].Credentials = credentials
			found = true
			break
		}
	}
	if !found {
		config.Workspaces = append(config.Workspaces, WorkspaceConfig{Name: workspaceName, Credentials: credentials})
	}
	config.CurrentContext = workspaceName
	saveConfig(config)
}

func ClearCredentials(workspaceName string) {
	config := loadConfig()
	for i, workspace := range config.Workspaces {
		if workspace.Name == workspaceName {
			config.Workspaces = append(config.Workspaces[:i], config.Workspaces[i+1:]...)
			break
		}
	}
	if config.CurrentContext == workspaceName {
		config.CurrentContext = ""
	}
	saveConfig(config)
}
