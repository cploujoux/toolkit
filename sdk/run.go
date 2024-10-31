package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) Run(ctx context.Context, workspaceName string, environment string, modelName string, body string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	fmt.Println("run", workspaceName, environment, modelName, body)
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	fmt.Println("bodyReader", bodyReader)
	runUrl, err := url.Parse(c.RunServer)
	fmt.Println("server", runUrl)

	// url := "https://edge-gw.beamlit.net/" + workspaceName + "/models/" + modelName + "?environment=" + environment

	// payload := strings.NewReader(data)

	// req, _ := http.NewRequest("POST", url, payload)
	// provider.Intercept(ctx, req)

	// req.Header.Add("content-type", "application/json")

	// fmt.Printf("URL: %s\n", req.URL)
	// fmt.Printf("Method: %s\n", req.Method)
	// fmt.Println("Headers:")
	// for key, values := range req.Header {
	// 	for _, value := range values {
	// 		fmt.Printf("  %s: %s\n", key, value)
	// 	}
	// }
	// fmt.Println("Payload:", data)

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	// return res, nil
	return nil, nil
}
