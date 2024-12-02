package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

type ErrorModel struct {
	Error string   `json:"error"`
	Code  int      `json:"code"`
	Stack []string `json:"stack"`
}

func ErrorHandler(kind string, name string, body string) {
	var error ErrorModel
	if err := json.Unmarshal([]byte(body), &error); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Afficher l'erreur et le code
	fmt.Printf("Resource %s:%s: %s (Code: %d)\n", kind, name, error.Error, error.Code)

	// Afficher le stack trace seulement s'il existe
	if verbose && len(error.Stack) > 0 {
		fmt.Println("Stack trace:")
		for _, line := range error.Stack {
			fmt.Printf("  %s\n", line)
		}
	}
}
