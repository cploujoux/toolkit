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
	"gopkg.in/yaml.v3"
)

func (r *Operations) ApplyCmd() *cobra.Command {
	var filePath string

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a configuration to a resource by file",
		Long:  "Apply a configuration to a resource by file",
		Example: `
			beamlit apply -f ./my-deployment.yaml
			# Or using stdin
			cat file.yaml | beamlit apply -f -
		`,
		Run: func(cmd *cobra.Command, args []string) {
			var reader io.Reader

			// Choisir la source (stdin ou fichier)
			if filePath == "-" {
				reader = os.Stdin
			} else {
				file, err := os.Open(filePath)
				if err != nil {
					fmt.Printf("Error opening file: %v\n", err)
					return
				}
				defer file.Close()
				reader = file
			}

			// Lire et parser les documents YAML
			decoder := yaml.NewDecoder(reader)
			var results []Result

			for {
				var result Result
				err := decoder.Decode(&result)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("Error decoding YAML: %v\n", err)
					return
				}
				results = append(results, result)
			}

			// À ce stade, results contient tous vos documents YAML
			for _, result := range results {
				for _, resource := range resources {
					if resource.Kind == result.Kind {
						resource.PutFn(resource.Kind, result.Metadata.Name, result.Spec)
					}
				}
			}
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to YAML file to apply")
	cmd.MarkFlagRequired("file")

	return cmd
}

func (resource Resource) PutFn(resourceName string, name string, spec interface{}) {
	ctx := context.Background()
	// Use reflect to call the function
	funcValue := reflect.ValueOf(resource.Put)
	if funcValue.Kind() != reflect.Func {
		fmt.Println("fn is not a valid function")
		os.Exit(1)
	}
	// Convert spec to the expected type using JSON marshaling/unmarshaling
	specJson, err := json.Marshal(spec)
	if err != nil {
		fmt.Printf("Error marshaling spec: %v\n", err)
		os.Exit(1)
	}

	// Create a new instance of the expected type
	destBody := reflect.New(resource.SpecType).Interface()
	if err := json.Unmarshal(specJson, destBody); err != nil {
		fmt.Printf("Error unmarshaling to target type: %v\n", err)
		os.Exit(1)
	}

	// Use the converted spec in the function call
	fnargs := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(name),
		reflect.ValueOf(destBody).Elem(),
	}

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
	if response.StatusCode >= 400 {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Printf("Resource %s:%s configured\n", resourceName, name)
}
