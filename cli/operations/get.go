package operations

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get operations",
}

var outputFormat string

func GetRegister(ctx context.Context, operation []string, fn interface{}) {
	use := operation[1]
	cmd := &cobra.Command{
		Use:   use,
		Short: fmt.Sprintf("Execute %s operation", operation[1]),
		Run:   GetFunc(ctx, operation, fn),
	}
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")

	GetCmd.AddCommand(cmd)

	singular := putToSingular(use)
	if singular != use {
		cmd := &cobra.Command{
			Use:   singular,
			Short: fmt.Sprintf("Execute %s operation", operation[1]),
			Run:   GetFunc(ctx, operation, fn),
		}
		cmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")

		GetCmd.AddCommand(cmd)
	}
}

func GetFunc(ctx context.Context, operation []string, fn interface{}) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Use reflect to call the function
		funcValue := reflect.ValueOf(fn)
		if funcValue.Kind() != reflect.Func {
			fmt.Println("fn is not a valid function")
			os.Exit(1)
		}
		// Create a slice for the arguments
		fnargs := []reflect.Value{reflect.ValueOf(ctx)} // Add the context

		// Call the function with the arguments
		results := funcValue.Call(fnargs)

		// Handle the results based on your needs
		if len(results) <= 1 {
			return
		}

		if err, ok := results[1].Interface().(error); ok && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Check if the first result is a pointer to http.Response
		response, ok := results[0].Interface().(*http.Response)
		if !ok {
			fmt.Println("the result is not a pointer to http.Response")
			os.Exit(1)
		}
		// Read the content of http.Response.Body
		defer response.Body.Close() // Ensure to close the ReadCloser
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, response.Body); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Check if the content is an array or an object

		var slices []interface{}
		if err := json.Unmarshal(buf.Bytes(), &slices); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Check the output format

		if outputFormat != "yaml" {
			printTable(slices)
			return
		}

		caser := cases.Title(language.English)
		formatted := []Result{}
		for _, slice := range slices {
			if sliceMap, ok := slice.(map[string]interface{}); ok {
				formatted = append(formatted, Result{
					Kind: caser.String(putToSingular(operation[1])),
					Metadata: ResultMetadata{
						Workspace: sliceMap["workspace"].(string),
						Name:      sliceMap["name"].(string),
					},
					Spec: slice,
				})
			}
		}

		// Convert the dynamic object to YAML
		yamlData, err := yaml.Marshal(formatted)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Print the YAML with colored keys and values
		printColoredYAML(yamlData)
	}
}

type ResultMetadata struct {
	Workspace string
	Name      string
}

type Result struct {
	Kind     string
	Metadata ResultMetadata
	Spec     interface{}
}
