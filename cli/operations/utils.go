package operations

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func printColoredYAML(yamlData []byte) {
	lines := bytes.Split(yamlData, []byte("\n"))
	keyColor := color.New(color.FgBlue).SprintFunc()
	stringValueColor := color.New(color.FgGreen).SprintFunc()
	numberValueColor := color.New(color.FgYellow).SprintFunc()

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// Split the line into key and value
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) < 2 {
			fmt.Println(string(line))
			continue
		}
		key := parts[0]
		value := parts[1]

		// Determine the type of value and color it accordingly
		var coloredValue string
		if bytes.HasPrefix(value, []byte(" ")) {
			value = bytes.TrimSpace(value)
			if len(value) > 0 && (value[0] == '"' || value[0] == '\'') {
				coloredValue = stringValueColor(string(value))
			} else if _, err := fmt.Sscanf(string(value), "%f", new(float64)); err == nil {
				coloredValue = numberValueColor(string(value))
			} else {
				coloredValue = string(value)
			}
		}

		// Print the colored key and value
		fmt.Printf("%s: %s\n", keyColor(string(key)), coloredValue)
	}
}

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
