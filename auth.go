package ultradns

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type OAuthResponse struct {
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func Authenticate(creds string) (OAuthResponse, error) {
	var OAuth OAuthResponse
	var credsJson Creds

	err := json.NewDecoder(strings.NewReader(string(creds))).Decode(&credsJson)
	if err != nil {
		return OAuth, err
	}

	data := url.Values{}
	data.Set("username", credsJson.Username)
	data.Add("password", credsJson.Password)
	data.Add("grant_type", "password")

	client := &http.Client{}
	resp, err := client.Post("https://restapi.ultradns.com/v2/authorization/token",
		"Content-Type: application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
	if err != nil {
		return OAuth, err
	}
	if resp.StatusCode != 200 {
		return OAuth, errors.New("Auth Failed")
	}

	err = json.NewDecoder(resp.Body).Decode(&OAuth)
	if err != nil {
		return OAuth, errors.New("Error decoding Response\n")
	}
	return OAuth, nil
}
