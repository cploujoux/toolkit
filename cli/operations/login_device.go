package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tmp-moon/toolkit/sdk"
)

func (r *Operations) DeviceModeLogin(workspace string) {
	url := r.BaseURL + "/login/device"

	payload := sdk.DeviceLogin{
		ClientID: "moon",
		Scope:    "offline_access",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var deviceLoginResponse sdk.DeviceLoginResponse
	if err := json.Unmarshal(body, &deviceLoginResponse); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Please visit the following URL to finish login: %s\n", deviceLoginResponse.VerificationURIComplete)

	r.DeviceModeLoginFinalize(deviceLoginResponse.DeviceCode, workspace)
}

func (r *Operations) DeviceModeLoginFinalize(userCode string, workspace string) {
	time.Sleep(3 * time.Second)
	url := r.BaseURL + "/oauth/token"

	payload := sdk.DeviceLoginFinalizeRequest{
		GrantType:  "urn:ietf:params:oauth:grant-type:device_code",
		ClientID:   "beamlit",
		DeviceCode: userCode,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var finalizeResponse sdk.DeviceLoginFinalizeResponse
	if err := json.Unmarshal(body, &finalizeResponse); err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		r.DeviceModeLoginFinalize(userCode, workspace)
	}

	creds := sdk.Credentials{
		AccessToken:  finalizeResponse.AccessToken,
		RefreshToken: finalizeResponse.RefreshToken,
		ExpiresIn:    finalizeResponse.ExpiresIn,
		Workspace:    workspace,
	}

	jsonData, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		panic(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	credentialsDir := filepath.Join(homeDir, ".beamlit")
	credentialsFile := filepath.Join(credentialsDir, "credentials.json")
	if err := os.WriteFile(credentialsFile, jsonData, 0600); err != nil {
		fmt.Printf("Error writing credentials file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully logged in")
}
