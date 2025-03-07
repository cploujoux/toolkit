package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/beamlit/toolkit/sdk"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type IgnoreFile struct {
	File string
	Skip string
}

type IgnoreDir struct {
	Folder string
	Skip   string
}

// CreateAgentAppOptions contains all the configuration options needed to create a new agent app.
type CreateAgentAppOptions struct {
	Directory          string             // Target directory for the new agent app
	ProjectName        string             // Name of the project
	ProjectDescription string             // Description of the project
	Language           string             // Language to use for the project
	Template           string             // Template to use for the project
	Author             string             // Author of the project
	TemplateOptions    map[string]*string // Options for the template
	IgnoreFiles        map[string]IgnoreFile
	IgnoreDirs         map[string]IgnoreDir
}

type TemplateConfig struct {
	Variables []struct {
		Name        string  `yaml:"name"`
		Label       *string `yaml:"label"`
		Type        string  `yaml:"type"`
		Description string  `yaml:"description"`
		File        string  `yaml:"file"`
		Skip        string  `yaml:"skip"`
		Folder      string  `yaml:"folder"`
		Options     []struct {
			Label  string `yaml:"label"`
			Value  string `yaml:"value"`
			Name   string `yaml:"name"`
			File   string `yaml:"file"`
			Skip   string `yaml:"skip"`
			Folder string `yaml:"folder"`
		} `yaml:"options"`
	} `yaml:"variables"`
}

type GithubTreeResponse struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

type GithubContentResponse struct {
	Content string `json:"content"`
}

// retrieveModels fetches and returns a list of available model deployments from the API.
// It filters the models to only include supported runtime types (openai, anthropic, mistral, etc.).
// Returns an error if the API calls fail or if there are parsing issues.
func retrieveModels(modelType string) ([]sdk.Model, error) {
	var modelDeployments []sdk.Model
	ctx := context.Background()
	res, err := client.ListModels(ctx)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var models []sdk.Model
	err = json.Unmarshal(body, &models)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		if model.Spec.Runtime != nil {
			runtimeType := *model.Spec.Runtime.Type
			modelName := *model.Spec.Runtime.Model
			if modelType == "model" {
				supportedRuntimes := []string{"openai", "anthropic", "mistral", "cohere", "xai", "vertex", "bedrock", "azure-ai-inference", "azure-marketplace", "gemini"}
				if slices.Contains(supportedRuntimes, runtimeType) && !strings.Contains(modelName, "realtime") {
					modelDeployments = append(modelDeployments, model)
				}
			} else if modelType == "realtime-model" {
				supportedRuntimes := []string{"openai", "azure-ai-inference", "azure-marketplace"}
				if slices.Contains(supportedRuntimes, runtimeType) && strings.Contains(modelName, "realtime") {
					modelDeployments = append(modelDeployments, model)
				}
			}
		}
	}
	return modelDeployments, nil
}

func getBranch() string {
	if os.Getenv("BL_ENV") == "dev" {
		return "develop"
	}
	return "main"
}

// retrieveTemplates retrieves the list of available templates from the templates repository.
// It fetches the repository's tree structure and extracts the paths of all directories.
// Returns a list of template names or an error if the retrieval fails.
func retrieveTemplates() ([]string, map[string][]string, error) {
	var scriptErr error
	languages := []string{}
	templates := map[string][]string{}
	spinnerErr := spinner.New().
		Title("Retrieving templates...").
		Action(func() {
			url := fmt.Sprintf("https://api.github.com/repos/beamlit/templates/git/trees/%s?recursive=1", getBranch())
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				scriptErr = err
				return
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				scriptErr = err
				return
			}

			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				scriptErr = err
				return
			}
			var treeResponse GithubTreeResponse
			err = json.Unmarshal(body, &treeResponse)
			if err != nil {
				scriptErr = err
				return
			}
			for _, tree := range treeResponse.Tree {
				if strings.HasPrefix(tree.Path, "agents/") && len(strings.Split(tree.Path, "/")) == 3 {
					language := strings.Split(tree.Path, "/")[1]
					if !slices.Contains(languages, language) {
						languages = append(languages, language)
					}
					if _, ok := templates[language]; !ok {
						templates[language] = []string{}
					}
					templates[language] = append(templates[language], strings.Split(tree.Path, "/")[2])
				}
			}
		}).
		Run()
	if spinnerErr != nil {
		return nil, nil, spinnerErr
	}
	if scriptErr != nil {
		return nil, nil, scriptErr
	}
	return languages, templates, nil
}

func retrieveTemplateConfig(language string, template string) (*TemplateConfig, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/beamlit/templates/refs/heads/%s/agents/%s/%s/template.yaml", getBranch(), language, template)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var templateConfig TemplateConfig
	err = yaml.Unmarshal(body, &templateConfig)
	if err != nil {
		return nil, err
	}
	return &templateConfig, nil
}

func promptTemplateConfig(agentAppOptions *CreateAgentAppOptions) {
	templateConfig, err := retrieveTemplateConfig(agentAppOptions.Language, agentAppOptions.Template)
	if err != nil {
		fmt.Println("Could not retrieve template configuration")
		os.Exit(0)
	}
	fields := []huh.Field{}
	values := map[string]*string{}
	array_values := map[string]*[]string{}
	mapped_values := map[string]string{}
	for _, variable := range templateConfig.Variables {
		var value string
		var array_value []string

		title := variable.Name
		if variable.Label != nil {
			title = *variable.Label
		}
		if variable.Type == "select" {
			values[variable.Name] = &value
			options := []huh.Option[string]{}
			if variable.File != "" {
				agentAppOptions.IgnoreFiles[variable.Name] = IgnoreFile{File: variable.File, Skip: variable.Skip}
			}
			if variable.Folder != "" {
				agentAppOptions.IgnoreDirs[variable.Name] = IgnoreDir{Folder: variable.Folder, Skip: variable.Skip}
			}
			for _, option := range variable.Options {
				options = append(options, huh.NewOption(option.Label, option.Value))
			}
			input := huh.NewSelect[string]().
				Title(title).
				Description(variable.Description).
				Options(options...).
				Value(&value)
			fields = append(fields, input)
		} else if variable.Type == "input" {
			values[variable.Name] = &value
			input := huh.NewInput().
				Title(title).
				Description(variable.Description).
				Value(&value)
			fields = append(fields, input)
		} else if variable.Type == "multiselect" {
			array_values[variable.Name] = &array_value
			options := []huh.Option[string]{}
			for _, option := range variable.Options {
				mapped_values[option.Value] = option.Name
				if option.File != "" {
					agentAppOptions.IgnoreFiles[option.Name] = IgnoreFile{File: option.File, Skip: option.Skip}
				}
				if option.Folder != "" {
					agentAppOptions.IgnoreDirs[option.Name] = IgnoreDir{Folder: option.Folder, Skip: option.Skip}
				}
				options = append(options, huh.NewOption(option.Label, option.Value))
			}
			input := huh.NewMultiSelect[string]().
				Title(title).
				Description(variable.Description).
				Options(options...).
				Value(&array_value)
			fields = append(fields, input)
		} else if variable.Type == "model" || variable.Type == "realtime-model" {
			values[variable.Name] = &value
			models, err := retrieveModels(variable.Type)
			if err == nil {
				if len(models) == 0 {
					value = "None"
				} else if len(models) == 1 {
					value = *models[0].Metadata.Name
				} else {
					options := []huh.Option[string]{}
					for _, model := range models {
						options = append(options, huh.NewOption(*model.Metadata.Name, *model.Metadata.Name))
					}
					options = append(options, huh.NewOption("None", ""))
					input := huh.NewSelect[string]().
						Title(title).
						Description(variable.Description).
						Height(5).
						Options(options...).
						Value(&value)
					fields = append(fields, input)
				}
			}
		}
	}

	if len(fields) > 0 {
		formTemplates := huh.NewForm(
			huh.NewGroup(fields...),
		)
		formTemplates.WithTheme(GetHuhTheme())
		err = formTemplates.Run()
		if err != nil {
			fmt.Println("Cancel create blaxel agent app")
			os.Exit(0)
		}
	}
	agentAppOptions.TemplateOptions = values
	for _, array_value := range array_values {
		for _, value := range *array_value {
			k := mapped_values[value]
			agentAppOptions.TemplateOptions[k] = &value
		}
	}
}

// promptCreateAgentApp displays an interactive form to collect user input for creating a new agent app.
// It prompts for project name, model selection, template, author, license, and additional features.
// Takes a directory string parameter and returns a CreateAgentAppOptions struct with the user's selections.
func promptCreateAgentApp(directory string) CreateAgentAppOptions {
	agentAppOptions := CreateAgentAppOptions{
		ProjectName: directory,
		Directory:   directory,
		IgnoreFiles: map[string]IgnoreFile{},
		IgnoreDirs:  map[string]IgnoreDir{},
	}
	currentUser, err := user.Current()
	if err == nil {
		agentAppOptions.Author = currentUser.Username
	} else {
		agentAppOptions.Author = "blaxel"
	}
	languages, templates, err := retrieveTemplates()
	if err != nil {
		fmt.Println("Could not retrieve templates")
		os.Exit(0)
	}
	languagesOptions := []huh.Option[string]{}
	for _, language := range languages {
		languagesOptions = append(languagesOptions, huh.NewOption(language, language))
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Description("Name of your agent app").
				Value(&agentAppOptions.ProjectName),
			huh.NewSelect[string]().
				Title("Language").
				Description("Language to use for your agent app").
				Height(5).
				Options(languagesOptions...).
				Value(&agentAppOptions.Language),
			huh.NewSelect[string]().
				Title("Template").
				Description("Template to use for your agent app").
				Height(5).
				OptionsFunc(func() []huh.Option[string] {
					templates := templates[agentAppOptions.Language]
					if len(templates) == 0 {
						return []huh.Option[string]{}
					}
					options := []huh.Option[string]{}
					for _, template := range templates {
						key := regexp.MustCompile(`^\d+-`).ReplaceAllString(template, "")
						options = append(options, huh.NewOption(key, template))
					}
					return options
				}, &agentAppOptions).
				Value(&agentAppOptions.Template),
		),
	)
	form.WithTheme(GetHuhTheme())
	err = form.Run()
	if err != nil {
		fmt.Println("Cancel create blaxel agent app")
		os.Exit(0)
	}
	promptTemplateConfig(&agentAppOptions)

	return agentAppOptions
}

func installPythonDependencies(directory string) error {
	uvSyncCmd := exec.Command("uv", "sync", "--refresh")
	uvSyncCmd.Dir = directory

	// Capture both stdout and stderr
	output, err := uvSyncCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run uv sync: %w\nOutput: %s", err, string(output))
	}
	return nil
}

func installTypescriptDependencies(directory string) error {
	npmInstallCmd := exec.Command("npm", "install")
	npmInstallCmd.Dir = directory

	// Capture both stdout and stderr
	output, err := npmInstallCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run npm install: %w\nOutput: %s", err, string(output))
	}
	return nil
}

// createAgentApp handles the actual creation of the agent app based on the provided options.
// It performs the following steps:
// 1. Creates the project directory
// 2. Clones the templates repository
// 3. Processes template files
// 4. Installs dependencies using uv sync
// Returns an error if any step fails.
func createAgentApp(opts CreateAgentAppOptions) error {
	// Create project directory
	if err := os.MkdirAll(opts.Directory, 0755); err != nil {
		return err
	}

	// Clone templates repository
	cloneDir := filepath.Join(opts.Directory, "templates")
	cloneDirCmd := exec.Command("git", "clone", "https://github.com/beamlit/templates.git", cloneDir)
	if err := cloneDirCmd.Run(); err != nil {
		return fmt.Errorf("failed to clone templates repository: %w", err)
	}

	templateOptions := map[string]string{
		"ProjectName":        opts.ProjectName,
		"ProjectDescription": opts.ProjectDescription,
		"Author":             opts.Author,
		"Workspace":          workspace,
	}
	for key, value := range opts.TemplateOptions {
		templateOptions[key] = *value
	}

	// Initialize ignore files and folders
	ignoreFiles := []string{"template.yaml"}
	ignoreFolders := []string{}
	for key, ignoreFile := range opts.IgnoreFiles {
		value, ok := templateOptions[key]
		if ok {
			if ignoreFile.Skip == value {
				ignoreFiles = append(ignoreFiles, ignoreFile.File)
			}
		} else {
			if ignoreFile.Skip == "" {
				ignoreFiles = append(ignoreFiles, ignoreFile.File)
			}
		}
	}
	for key, ignoreDir := range opts.IgnoreDirs {
		value, ok := templateOptions[key]
		if ok {
			if ignoreDir.Skip == value {
				ignoreFolders = append(ignoreFolders, ignoreDir.Folder)
			}
		} else {
			if ignoreDir.Skip == "" {
				ignoreFolders = append(ignoreFolders, ignoreDir.Folder)
			}
		}
	}
	templateDir := filepath.Join(cloneDir, "agents", opts.Language, opts.Template)
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Calculate relative path
		rel, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		// Skip files based on config
		for _, ignoreFile := range ignoreFiles {
			if strings.HasSuffix(rel, ignoreFile) {
				return nil
			}
		}

		// Skip folders based on config
		for _, ignoreFolder := range ignoreFolders {
			if strings.HasPrefix(rel, ignoreFolder) {
				return nil
			}
		}

		// Process template
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}

		// Create output file
		outPath := filepath.Join(opts.Directory, rel)
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		out, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Execute template
		return tmpl.Execute(out, templateOptions)
	})
	if err != nil {
		return err
	}
	// Remove templates directory after processing
	if err := os.RemoveAll(cloneDir); err != nil {
		return fmt.Errorf("failed to remove templates directory: %w", err)
	}

	// Install dependencies based on language
	switch opts.Language {
	case "python":
		if err := installPythonDependencies(opts.Directory); err != nil {
			return err
		}
	case "typescript":
		if err := installTypescriptDependencies(opts.Directory); err != nil {
			return err
		}
	}
	return nil
}

// CreateAgentAppCmd returns a cobra.Command that implements the 'create-agent-app' CLI command.
// The command creates a new Blaxel agent app in the specified directory after collecting
// necessary configuration through an interactive prompt.
// Usage: bl create-agent-app directory
func (r *Operations) CreateAgentAppCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create-agent-app directory",
		Args:    cobra.MaximumNArgs(2),
		Aliases: []string{"ca", "caa"},
		Short:   "Create a new blaxel agent app",
		Long:    "Create a new blaxel agent app",
		Example: `bl create-agent-app my-agent-app`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				fmt.Println("Please provide a directory name")
				return
			}
			// Check if directory already exists
			if _, err := os.Stat(args[0]); !os.IsNotExist(err) {
				fmt.Printf("Error: %s already exists\n", args[0])
				return
			}
			opts := promptCreateAgentApp(args[0])

			var err error
			spinnerErr := spinner.New().
				Title("Creating your blaxel agent app...").
				Action(func() {
					err = createAgentApp(opts)
				}).
				Run()
			if spinnerErr != nil {
				fmt.Println("Error creating agent app", spinnerErr)
				return
			}
			if err != nil {
				fmt.Println("Error creating agent app", err)
				os.RemoveAll(opts.Directory)
				return
			}
			res, err := client.ListModels(context.Background())
			if err != nil {
				return
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return
			}

			var models []sdk.Model
			err = json.Unmarshal(body, &models)
			if err != nil {
				return
			}
			fmt.Printf(`Your blaxel agent app has been created. Start working on it:
cd %s;
bl serve --hotreload;
`, opts.Directory)
		},
	}
	return cmd
}
