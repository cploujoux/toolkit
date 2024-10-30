package beamlit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func deviceModeLogin(workspace string) {
	url := BASE_URL + "/login/device"

	payload := DeviceLogin{
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

	var deviceLoginResponse DeviceLoginResponse
	if err := json.Unmarshal(body, &deviceLoginResponse); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Please visit the following URL to finish login: %s\n", deviceLoginResponse.VerificationURIComplete)

	deviceModeLoginFinalize(deviceLoginResponse.DeviceCode, workspace)
}

func deviceModeLoginFinalize(userCode string, workspace string) {
	time.Sleep(3 * time.Second)
	url := BASE_URL + "/oauth/token"

	payload := DeviceLoginFinalizeRequest{
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

	var finalizeResponse DeviceLoginFinalizeResponse
	if err := json.Unmarshal(body, &finalizeResponse); err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		deviceModeLoginFinalize(userCode, workspace)
	}

	creds := Credentials{
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
