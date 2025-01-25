package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WeatherData structure pour mapper la réponse JSON de l'API
type WeatherData struct {
	CurrentConditions struct {
		Temperature float64 `json:"temp"`
		Humidity    float64 `json:"humidity"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"currentConditions"`
}

// FetchWeatherData récupère les données météo d'un lieu donné.
func FetchWeatherData(location, apiKey string) (*WeatherData, error) {
	// Construire l'URL pour l'appel API.
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?key=%s", location, apiKey)

	//Faire la requête HTTP
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	//Lire la réponse JSON
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	//Analyser le JSON en WeatherData
	var weatherData WeatherData

	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing du JSON : %v", err)
	}

	return &weatherData, nil
}
