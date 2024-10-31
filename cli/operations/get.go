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
	"gopkg.in/yaml.v2"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get operations",
}

var outputFormat string

type GetCommand struct {
	List interface{}
	Get  interface{}
}

var refs = map[string]GetCommand{}

func GetRegister(ctx context.Context, operation []string, fn interface{}) {
	if len(operation) > 2 {
		return
	}

	name := operation[1]
	var plural string
	var singular string

	if operation[0] == "list" {
		singular = putToSingular(name)
		plural = name
	}
	if operation[0] == "get" {
		singular = name
		plural = putToPlural(name)
	}

	ref, ok := refs[plural]
	if !ok {
		ref = GetCommand{}
	}
	if operation[0] == "list" {
		ref.List = fn
	}
	if operation[0] == "get" {
		ref.Get = fn
	}
	refs[plural] = ref

	if ref.List != nil && ref.Get != nil {
		cmd := &cobra.Command{
			Use:     plural,
			Aliases: []string{singular},
			Short:   fmt.Sprintf("Execute %s operation", plural),
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					ListFunc(ctx, operation, ref.List)(cmd, args, singular)
					return
				}
				GetFunc(ctx, operation, ref.Get)(cmd, args, singular)
			},
		}
		cmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")
		GetCmd.AddCommand(cmd)
	}
}

func GetFunc(ctx context.Context, operation []string, fn interface{}) func(cmd *cobra.Command, args []string, singular string) {
	return func(cmd *cobra.Command, args []string, singular string) {
		// Use reflect to call the function
		funcValue := reflect.ValueOf(fn)
		if funcValue.Kind() != reflect.Func {
			fmt.Println("fn is not a valid function")
			os.Exit(1)
		}
		// Create a slice for the arguments
		fnargs := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(args[0])} // Add the context and the resource name

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
		var resource interface{}
		if err := json.Unmarshal(buf.Bytes(), &resource); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output(singular, []interface{}{resource})
	}
}

func ListFunc(ctx context.Context, operation []string, fn interface{}) func(cmd *cobra.Command, args []string, singular string) {
	return func(cmd *cobra.Command, args []string, singular string) {
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
		output(singular, slices)
	}
}

func output(singular string, slices []interface{}) {

	if outputFormat != "yaml" {
		printTable(slices)
		return
	}

	caser := cases.Title(language.English)
	formatted := []Result{}
	for _, slice := range slices {
		if sliceMap, ok := slice.(map[string]interface{}); ok {
			formatted = append(formatted, Result{
				Kind: caser.String(singular),
				Metadata: ResultMetadata{
					Workspace: sliceMap["workspace"].(string),
					Name:      sliceMap["name"].(string),
				},
				Spec: slice,
			})
		}
	}

	// Convert each object to YAML and add separators
	var yamlData []byte
	for _, result := range formatted {
		data, err := yaml.Marshal(result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		yamlData = append(yamlData, []byte("---\n")...)
		yamlData = append(yamlData, data...)
		yamlData = append(yamlData, []byte("\n")...)
	}

	// Print the YAML with colored keys and values
	printColoredYAML(yamlData)
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
