package sdk

type Credentials struct {
	APIKey       string `json:"api_key"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Workspace    string `json:"workspace"`
	ExpiresIn    int    `json:"expires_in"`
}

type DeviceLogin struct {
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
}

type DeviceLoginResponse struct {
	ClientID                string `json:"client_id"`
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
}

type DeviceLoginFinalizeRequest struct {
	GrantType  string `json:"grant_type"`
	ClientID   string `json:"client_id"`
	DeviceCode string `json:"device_code"`
}

type DeviceLoginFinalizeResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}
