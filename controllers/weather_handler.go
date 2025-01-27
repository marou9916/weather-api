package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"weather-api/cache"
	"weather-api/models"
	"weather-api/services"

	"github.com/gin-gonic/gin"
)

func LocationWeatherHandler(c *gin.Context) {
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La localisation est requise"})
		return
	}

	// Essayer de récupérer les données du cache
	weatherDatasFromCache, err := cache.GetDatasFromCache(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des données"})
		return
	}

	if weatherDatasFromCache != "" {
		var locationWeatherDatas models.WeatherData
		err := json.Unmarshal([]byte(weatherDatasFromCache), &locationWeatherDatas)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du formatage des données du cache"})
		} else {
			c.JSON(http.StatusOK, locationWeatherDatas)
		}
		return
	}

	// Si les données ne sont pas dans le cache, appeler l'API de Visual Crossing pour récupérer les données
	weatherDatasFromVisualCrossing, err := services.FetchWeatherData(location, "MY_API_KEY")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'appel à l'API Visual Crossing"})
		return
	}

	//Sauvegarder les données dans le cache
	weatherDatasJSON, _ := json.Marshal(&weatherDatasFromVisualCrossing)
	cache.SaveDatasInCache(location, string(weatherDatasJSON), 15*time.Minute)

	//Répondre au client en servant les données
	c.JSON(http.StatusOK, weatherDatasFromVisualCrossing)

}
