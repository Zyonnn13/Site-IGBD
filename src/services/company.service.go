package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-igdb/src/models"
	"strings"
	"time"
)

func GetCompanies() ([]models.Company, int, error) {
	client := http.Client{Timeout: 5 * time.Second}
	url := "https://api.igdb.com/v4/companies"

	token, errToken := GetIGDBToken()
	if errToken != nil {
		return nil, http.StatusInternalServerError, errToken
	}

	queryBody := `
		fields name, description, logo.url, start_date;
		where id = (112, 70, 835, 104, 10100);
		sort name asc;
	`

	request, err := http.NewRequest("POST", url, strings.NewReader(queryBody))
	if err != nil {
		return nil, 0, fmt.Errorf("erreur création de la requête: %w", err)
	}

	request.Header.Add("Client-ID", ClientID)
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("l'API a retourné le code %d: %s", response.StatusCode, response.Status)
	}

	var companies []models.Company
	err = json.NewDecoder(response.Body).Decode(&companies)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("erreur décodage JSON: %w", err)
	}

	for i := range companies {
		if !strings.HasPrefix(companies[i].Logo.URL, "https:") {
			companies[i].Logo.URL = "https:" + companies[i].Logo.URL
		}
		companies[i].Logo.URL = strings.Replace(companies[i].Logo.URL, "t_thumb", "t_logo_med", 1)
	}

	return companies, response.StatusCode, nil
}
