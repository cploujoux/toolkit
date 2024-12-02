package cli

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Entries : ListOperations, ListEnvironments, GetOperation, GetEnvironment
// Results [list, operations, list environments, get operation, get environment]
func formatOperationId(operationId string) []string {
	// Regular expression to match capital letters
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	// Find all matches and convert them to lowercase
	words := re.FindAllString(operationId, -1)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func getResults(filePath string, recursive bool) ([]Result, error) {
	var reader io.Reader
	var results []Result
	// Choisir la source (stdin ou fichier)
	if filePath == "-" {
		reader = os.Stdin
	} else {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("error getting file info: %v", err)
		}
		// If the path is a directory, read all files in the directory
		if fileInfo.IsDir() {
			if !recursive && strings.Contains(filePath, "/") {
				return nil, nil
			}
			return handleDirectory(filePath, recursive)
		}
		// Skip non-YAML files
		if !strings.HasSuffix(strings.ToLower(filePath), ".yml") && !strings.HasSuffix(strings.ToLower(filePath), ".yaml") {
			return nil, nil
		}
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()
		reader = file
	}
	// Read the entire content as a string first
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading content: %v", err)
	}

	// Replace environment variables in the content
	contentStr := string(content)
	re := regexp.MustCompile(`\$([A-Za-z0-9_]+)|\${([A-Za-z0-9_]+)}`)
	contentStr = re.ReplaceAllStringFunc(contentStr, func(match string) string {
		// Remove $, ${, and } to get the env var name
		varName := match
		varName = strings.TrimPrefix(varName, "$")
		varName = strings.TrimPrefix(varName, "{")
		varName = strings.TrimSuffix(varName, "}")

		if value, exists := os.LookupEnv(varName); exists {
			return value
		}
		return match // Keep original if env var not found
	})

	// Lire et parser les documents YAML
	decoder := yaml.NewDecoder(strings.NewReader(contentStr))
	for {
		var result Result
		err := decoder.Decode(&result)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error decoding YAML: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

func handleDirectory(filePath string, recursive bool) ([]Result, error) {
	var results []Result
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %v", filePath, err)
	}

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", filePath, file.Name())
		fileResults, err := getResults(path, recursive)
		if err != nil {
			fmt.Printf("error getting results for file %s: %v", path, err)
			continue
		}
		results = append(results, fileResults...)
	}
	return results, nil
}
