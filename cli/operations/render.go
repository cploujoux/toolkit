package operations

import (
	"bytes"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func printTable(slices []interface{}) {
	// Print header with fixed width columns
	fmt.Printf("%-15s %-20s %-20s %-20s\n", "WORKSPACE", "NAME", "CREATED_AT", "UPDATED_AT")

	// Print each item in the array
	for _, item := range slices {
		// Convert item to map to access fields
		if itemMap, ok := item.(map[string]interface{}); ok {
			// Get the workspace field, default to "-" if not found
			workspace, _ := itemMap["workspace"].(string)
			if workspace == "" {
				workspace = "-"
			}
			// Get the name field, default to "-" if not found
			name, _ := itemMap["name"].(string)
			if name == "" {
				name = "-"
			}

			// Get the created_at field, default to "-" if not found
			createdAt, _ := itemMap["created_at"].(string)
			if createdAt == "" {
				createdAt = "-"
			} else {
				// Parse and format the date
				if parsedTime, err := time.Parse(time.RFC3339, createdAt); err == nil {
					createdAt = parsedTime.Format("2006-01-02 15:04:05")
				}
			}

			// Get the updated_at field, default to "-" if not found
			updatedAt, _ := itemMap["updated_at"].(string)
			if updatedAt == "" {
				updatedAt = "-"
			} else {
				// Parse and format the date
				if parsedTime, err := time.Parse(time.RFC3339, updatedAt); err == nil {
					updatedAt = parsedTime.Format("2006-01-02 15:04:05")
				}
			}

			// Print the fields with fixed width columns
			fmt.Printf("%-15s %-20s %-20s %-20s\n", workspace, name, createdAt, updatedAt)
		}
	}
}

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
