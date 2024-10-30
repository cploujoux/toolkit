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
	"gopkg.in/yaml.v2"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete operations",
}

func DeleteRegister(ctx context.Context, operation []string, fn interface{}) {
	cmd := &cobra.Command{
		Use:   operation[1],
		Short: fmt.Sprintf("Execute %s operation", operation[1]),
		Run: func(cmd *cobra.Command, args []string) {
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

			// Convert the dynamic object to YAML
			yamlData, err := yaml.Marshal(resource)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print the YAML with colored keys and values
			printColoredYAML(yamlData)
		},
	}
	DeleteCmd.AddCommand(cmd)
}
