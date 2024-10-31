package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Inference(ctx context.Context, workspaceName string, environment string, modelName string, data string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	if workspaceName == "" {
		workspaceName = CurrentWorkspace()
	}
	credentials := LoadCredentials(workspaceName)
	var provider AuthProvider
	if credentials.AccessToken != "" {
		provider = NewBearerTokenProvider(credentials, workspaceName)
	} else if credentials.APIKey != "" {
		provider = NewApiKeyProvider(credentials, workspaceName)
	} else {
		provider = NewPublicProvider()
	}

	url := "https://run.beamlit.dev/" + workspaceName + "/models/" + modelName + "?environment=" + environment

	payload := strings.NewReader(data)

	req, _ := http.NewRequest("POST", url, payload)
	provider.Intercept(ctx, req)

	req.Header.Add("content-type", "application/json")

	fmt.Printf("URL: %s\n", req.URL)
	fmt.Printf("Method: %s\n", req.Method)
	fmt.Println("Headers:")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	fmt.Println("Payload:", data)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return res, nil
}
