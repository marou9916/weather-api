package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	models "weather-api/models"

	"github.com/joho/godotenv"
)

func FetchWeatherData(location string) (*models.WeatherData, error) {
	//Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Impossible de charger le fichier .env, utilisation des variables d'env système")
	}

	apiKey := os.Getenv("VISUAL_CROSSING_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}

	//Construire l'url
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", location, apiKey)

	//Faire la requête http
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	//Lire la réponse JSON
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	//Analyser le JSON
	var weatherData models.WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing du JSON : %v", err)
	}

	return &weatherData, nil
}
