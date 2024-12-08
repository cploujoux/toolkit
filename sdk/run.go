package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) Run(
	ctx context.Context,
	workspaceName string,
	environment string,
	resourceType string,
	resourceName string,
	method string,
	path string,
	headers map[string]string,
	body string,
	reqEditors ...RequestEditorFn,
) (*http.Response, error) {
	var bodyReader io.Reader
	if body != "" {
		bodyReader = bytes.NewReader([]byte(body))
	}

	req, err := NewRunRequest(c.RunServer, method, path, headers, workspaceName, environment, resourceType, resourceName, bodyReader)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func NewRunRequest(
	RunServer string,
	method string,
	path string,
	headers map[string]string,
	workspaceName string,
	environment string,
	resourceType string,
	resourceName string,
	body io.Reader,
) (*http.Request, error) {
	var err error

	runURL, err := url.Parse(RunServer)
	if err != nil {
		return nil, err
	}
	if path != "" {
		path = fmt.Sprintf("%s/%ss/%s/%s", workspaceName, resourceType, resourceName, path)
	} else {
		path = fmt.Sprintf("%s/%ss/%s", workspaceName, resourceType, resourceName)
	}

	queryURL, err := runURL.Parse(path + "?environment=" + environment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/json"
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}
