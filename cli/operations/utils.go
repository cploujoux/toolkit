package operations

import (
	"regexp"
	"strings"
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

func putToSingular(use string) string {
	// Common irregular plurals
	irregulars := map[string]string{
		"children": "child",
		"people":   "person",
		"mice":     "mouse",
		"teeth":    "tooth",
	}

	// Check for irregular plurals first
	if singular, ok := irregulars[use]; ok {
		return singular
	}

	// Handle regular plural rules
	if len(use) < 2 {
		return use
	}

	// Handle words ending in 'ies' (e.g., policies -> policy)
	if strings.HasSuffix(use, "ies") {
		return use[:len(use)-3] + "y"
	}

	// Handle words ending in 'es' (e.g., boxes -> box)
	if strings.HasSuffix(use, "es") {
		// Special cases like 'analyses' -> 'analysis'
		if strings.HasSuffix(use, "ses") && len(use) > 4 {
			return use[:len(use)-2]
		}
		return use[:len(use)-2]
	}

	// Handle regular plural 's'
	if use[len(use)-1] == 's' {
		return use[:len(use)-1]
	}

	return use
}
