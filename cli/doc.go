package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func (r *Operations) DocCmd() *cobra.Command {
	var format string
	var outputDir string

	docCmd := &cobra.Command{
		Use:    "docs",
		Short:  "Generate documentation for the CLI",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}

			switch format {
			case "markdown":
				return doc.GenMarkdownTree(rootCmd, outputDir)
			case "man":
				header := &doc.GenManHeader{
					Title:   "BEAMLIT",
					Section: "1",
				}
				return doc.GenManTree(rootCmd, header, outputDir)
			case "rst":
				return doc.GenReSTTree(rootCmd, outputDir)
			case "yaml":
				return doc.GenYamlTree(rootCmd, outputDir)
			default:
				return fmt.Errorf("unknown format %s", format)
			}
		},
	}

	docCmd.Flags().StringVarP(&format, "format", "f", "markdown", "Documentation format (markdown, man, rst, yaml)")
	docCmd.Flags().StringVarP(&outputDir, "output", "o", "./docs", "Output directory for documentation")

	return docCmd
}
