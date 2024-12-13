package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"slices"
	"text/template"

	"github.com/beamlit/toolkit/sdk"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// CreateAgentAppOptions contains all the configuration options needed to create a new agent app.
type CreateAgentAppOptions struct {
	Directory          string   // Target directory for the new agent app
	ProjectName        string   // Name of the project
	ProjectDescription string   // Description of the project
	Model              string   // Selected AI model
	Template           string   // Template to use for the project
	Author             string   // Author of the project
	License            string   // License type (mit, apache, gpl)
	Features           []string // Additional features to include
}

// getTheme returns a custom theme configuration for the CLI interface using the Dracula color scheme.
// It customizes various UI elements like buttons, text inputs, and selection indicators.
func getTheme() *huh.Theme {
	t := huh.ThemeBase()
	var (
		background = lipgloss.AdaptiveColor{Dark: "#282a36"}
		selection  = lipgloss.AdaptiveColor{Dark: "#44475a"}
		foreground = lipgloss.AdaptiveColor{Dark: "#f8f8f2"}
		comment    = lipgloss.AdaptiveColor{Dark: "#6272a4"}
		green      = lipgloss.AdaptiveColor{Dark: "#50fa7b"}
		orange     = lipgloss.AdaptiveColor{Dark: "#fd7b35"}
		red        = lipgloss.AdaptiveColor{Dark: "#ff5555"}
		yellow     = lipgloss.AdaptiveColor{Dark: "#f1fa8c"}
	)

	t.Focused.Base = t.Focused.Base.BorderForeground(selection)
	t.Focused.Title = t.Focused.Title.Foreground(orange)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(orange)
	t.Focused.Description = t.Focused.Description.Foreground(comment)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(red)
	t.Focused.Directory = t.Focused.Directory.Foreground(orange)
	t.Focused.File = t.Focused.File.Foreground(foreground)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(red)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(yellow)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(yellow)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(yellow)
	t.Focused.Option = t.Focused.Option.Foreground(foreground)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(yellow)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(green)
	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.Foreground(green).SetString("[✓] ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(foreground)
	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.Foreground(comment)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(yellow).Background(orange).Bold(true)
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(foreground).Background(background)

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(yellow)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(comment)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(yellow)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

// retrieveModels fetches and returns a list of available model deployments from the API.
// It filters the models to only include supported runtime types (openai, anthropic, mistral, etc.).
// Returns an error if the API calls fail or if there are parsing issues.
func retrieveModels() ([]sdk.ModelDeployment, error) {
	var modelDeployments []sdk.ModelDeployment
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
		res, err := client.GetModelDeployment(ctx, *model.Name, "production")
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var modelDeployment sdk.ModelDeployment
		err = json.Unmarshal(body, &modelDeployment)
		if err != nil {
			return nil, err
		}
		if modelDeployment.Runtime != nil {
			runtimeType := *modelDeployment.Runtime.Type
			supportedRuntimes := []string{"openai", "anthropic", "mistral", "cohere", "xai", "vertex", "bedrock"}
			if slices.Contains(supportedRuntimes, runtimeType) {
				modelDeployments = append(modelDeployments, modelDeployment)
			}
		}
	}
	return modelDeployments, nil
}

// promptCreateAgentApp displays an interactive form to collect user input for creating a new agent app.
// It prompts for project name, model selection, template, author, license, and additional features.
// Takes a directory string parameter and returns a CreateAgentAppOptions struct with the user's selections.
func promptCreateAgentApp(directory string) CreateAgentAppOptions {
	var (
		projectName string
		author      string
		license     string
		model       string
		template    string
		features    []string
	)
	projectName = directory
	currentUser, err := user.Current()
	if err == nil {
		author = currentUser.Username
	}
	template = "empty"
	model = "gpt-4o-mini"
	license = "mit"
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Description("Name of your agent app").
				Value(&projectName),
			huh.NewSelect[string]().
				Title("Beamlit Model").
				Description("Model to use for your agent app").
				Height(5).
				OptionsFunc(func() []huh.Option[string] {
					options := []huh.Option[string]{}
					models, err := retrieveModels()
					if err != nil {
						return options
					}
					for _, model := range models {
						options = append(options, huh.NewOption(*model.Model, *model.Model))
					}
					return options
				}, &model).
				Value(&model),
			huh.NewInput().
				Title("Template").
				Description("Template to use for your agent app").
				Value(&template),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Author").
				Value(&author),
			huh.NewSelect[string]().
				Title("License").
				Options(
					huh.NewOption("MIT", "mit"),
					huh.NewOption("Apache 2.0", "apache"),
					huh.NewOption("GPL 3.0", "gpl"),
				).
				Value(&license),
			huh.NewMultiSelect[string]().
				Title("Features").
				Options(
					huh.NewOption("Intialized README", "readme"),
					huh.NewOption("Github action", "github-action"),
					huh.NewOption("Ruff", "ruff"),
				).
				Value(&features),
		),
	)
	form.WithTheme(getTheme())
	err = form.Run()
	if err != nil {
		fmt.Println("Cancel create beamlit agent app")
		os.Exit(0)
	}
	return CreateAgentAppOptions{
		Directory:   directory,
		ProjectName: projectName,
		Author:      author,
		License:     license,
		Model:       model,
		Template:    template,
		Features:    features,
	}
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

	templateDir := filepath.Join(cloneDir, "agents", opts.Template)
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

		// // Skip files based on config
		// if !te.shouldProcessFile(rel) {
		// 	return nil
		// }

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
		return tmpl.Execute(out, opts)
	})
	if err != nil {
		return err
	}
	// Remove templates directory after processing
	if err := os.RemoveAll(cloneDir); err != nil {
		return fmt.Errorf("failed to remove templates directory: %w", err)
	}
	// Run uv sync to install dependencies
	uvSyncCmd := exec.Command("uv", "sync")
	uvSyncCmd.Dir = opts.Directory
	if err := uvSyncCmd.Run(); err != nil {
		return fmt.Errorf("failed to run uv sync: %w", err)
	}
	return nil
}

// CreateAgentAppCmd returns a cobra.Command that implements the 'create-agent-app' CLI command.
// The command creates a new Beamlit agent app in the specified directory after collecting
// necessary configuration through an interactive prompt.
// Usage: bl create-agent-app directory
func (r *Operations) CreateAgentAppCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create-agent-app directory",
		Args:    cobra.MaximumNArgs(2),
		Aliases: []string{"ca", "caa"},
		Short:   "Create a new beamlit agent app",
		Long:    "Create a new beamlit agent app",
		Example: `bl create-agent-app`,
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
				Title("Creating your beamlit agent app...").
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
			fmt.Printf(`Your beamlit agent app has been created. Start working on it:
cd %s && source .venv/bin/activate;
bl serve --local;
`, opts.Directory)
		},
	}
	return cmd
}
