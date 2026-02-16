package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ClientID     = "8kzdgpohvkbo08vyj14ieec7byhukn"
	ClientSecret = "fwatf4pv1l19b10r1qsju9dfyqxpvj"
)

type IGDBTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetIGDBToken() (string, error) {
	data := url.Values{}
	data.Set("client_id", ClientID)
	data.Set("client_secret", ClientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, "https://id.twitch.tv/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("erreur initialisation requête token - %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur envoi requête token - %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur réponse token - code : %d", resp.StatusCode)
	}

	var tokenResp IGDBTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("erreur décodage token - %s", err.Error())
	}
	return tokenResp.AccessToken, nil
}
