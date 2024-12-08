package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func (r *Operations) RunCmd() *cobra.Command {
	var data string
	var path string
	var method string
	var headerFlags []string
	var showHeaders bool
	var uploadFilePath string

	cmd := &cobra.Command{
		Use:     "run resource-type resource-name",
		Args:    cobra.ExactArgs(2),
		Short:   "Run a resource on beamlit",
		Example: "bl run agent my-agent --data '{\"inputs\": \"Hello, world!\"}'",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || len(args) == 1 {
				fmt.Println("Error: Resource type and name are required")
				os.Exit(1)
			}

			resourceType := args[0]
			resourceName := args[1]
			headers := make(map[string]string)

			// Parse header flags into map
			for _, header := range headerFlags {
				parts := strings.SplitN(header, ":", 2)
				if len(parts) != 2 {
					fmt.Printf("Error: Invalid header format '%s'. Must be 'Key: Value'\n", header)
					os.Exit(1)
				}
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				headers[key] = value
			}

			// Handle file upload if specified
			if uploadFilePath != "" {
				fileContent, err := os.ReadFile(uploadFilePath)
				if err != nil {
					fmt.Printf("Error reading file: %v\n", err)
					os.Exit(1)
				}
				data = string(fileContent)
			}

			res, err := client.Run(
				context.Background(),
				workspace,
				environment,
				resourceType,
				resourceName,
				method,
				path,
				headers,
				data,
			)
			if err != nil {
				fmt.Printf("Error making request: %v\n", err)
				os.Exit(1)
			}
			defer res.Body.Close()

			// Read response body
			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("Error reading response: %v\n", err)
				os.Exit(1)
			}
			// Only print status code if it's an error
			if res.StatusCode >= 400 {
				fmt.Printf("Response Status: %s\n", res.Status)
			}

			if showHeaders {
				fmt.Printf("Response Headers:\n")
				for key, values := range res.Header {
					for _, value := range values {
						fmt.Printf("  %s: %s\n", key, value)
					}
				}
			}

			// Try to pretty print JSON response
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
				fmt.Println(prettyJSON.String())
			} else {
				// If not JSON, print as string
				fmt.Println(string(body))
			}
		},
	}

	cmd.Flags().StringVar(&data, "data", "", "JSON body data for the inference request")
	cmd.Flags().StringVar(&path, "path", "", "path for the inference request")
	cmd.Flags().StringVar(&method, "method", "POST", "HTTP method for the inference request")
	cmd.Flags().StringVar(&uploadFilePath, "upload-file", "", "This transfers the specified local file to the remote URL")
	cmd.Flags().StringArrayVar(&headerFlags, "header", []string{}, "Request headers in 'Key: Value' format. Can be specified multiple times")
	cmd.Flags().BoolVar(&showHeaders, "show-headers", false, "Show response headers in output")
	return cmd
}
