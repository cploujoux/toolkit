package cli

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
)

func (r *Operations) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a resource",
	}
	var outputFormat string
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format. One of: yaml")
	for _, resource := range resources {
		subcmd := &cobra.Command{
			Use:     resource.Short,
			Aliases: []string{resource.Singular, resource.Plural},
			Short:   fmt.Sprintf("Get a %s resource", resource.Kind),
			Run: func(cmd *cobra.Command, args []string) {
				outputFormat := cmd.Flag("output").Value.String()
				if len(args) == 0 {
					resource.ListFn(outputFormat)
					return
				}
				if len(args) == 1 {
					resource.GetFn(args[0], outputFormat)
				}
			},
		}
		subcmd.Flags().StringP("output", "o", "", "Output format. One of: yaml")
		cmd.AddCommand(subcmd)
	}

	return cmd
}

func (resource Resource) GetFn(name string, outputFormat string) {
	ctx := context.Background()
	// Use reflect to call the function
	funcValue := reflect.ValueOf(resource.Get)
	if funcValue.Kind() != reflect.Func {
		fmt.Println("fn is not a valid function")
		os.Exit(1)
	}
	// Create a slice for the arguments
	fnargs := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(name)} // Add the context and the resource name

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
	var res interface{}
	if err := json.Unmarshal(buf.Bytes(), &res); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output(resource, []interface{}{res}, outputFormat)
}

func (resource Resource) ListFn(outputFormat string) {
	ctx := context.Background()
	// Use reflect to call the function
	funcValue := reflect.ValueOf(resource.List)
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
	output(resource, slices, outputFormat)
}
