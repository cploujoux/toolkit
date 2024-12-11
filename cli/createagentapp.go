package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"text/template"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type CreateAgentAppOptions struct {
	Directory          string
	ProjectName        string
	ProjectDescription string
	Model              string
	Template           string
	Author             string
	License            string
	Features           []string
	Ruff               string
}

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
			huh.NewInput().
				Title("Beamlit Model").
				Description("Model to use for your agent app").
				Value(&model),
			huh.NewInput().
				Title("Template").
				Description("Template to use for your agent app").
				Value(&template),
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

func (r *Operations) CreateAgentAppCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create-agent-app [directory]",
		Args:  cobra.MaximumNArgs(2),
		Short: "Create a new beamlit agent app",
		Long:  "Create a new beamlit agent app",
		Example: `
			bl create-agent-app
		`,
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
			spinner.New().
				Title("Creating your beamlit agent app...").
				Action(func() {
					err = createAgentApp(opts)

				}).
				Run()

			if err != nil {
				fmt.Println("Error creating agent app", err)
				os.RemoveAll(opts.Directory)
				return
			}
			fmt.Printf(`Your beamlit agent app has been created. Start working on it:
cd %s
bl serve --local
`, opts.Directory)
		},
	}
	return cmd
}
