package controllers

import (
	"net/http"
	"time"
	"weather-api/cache"
	"weather-api/services"

	"github.com/gin-gonic/gin"
)

// LocationWeatherDatasHandler traite la demande des données météo pour une localisation donnée
func LocationWeatherDatasHandler(c *gin.Context) {
	location := c.Query("location")

	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "localisation requise"})
		return
	}

	//Vérifier le cache
	weatherDatasFromCache, err := cache.GetDatasFromCache(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erreur lors de la récupération des données du cache"})
		return
	}
	//Si les données sont présentes
	if weatherDatasFromCache != nil {
		c.JSON(http.StatusOK, weatherDatasFromCache)
		return
	}

	//Si elles sont absentes, appeler l'api de visual crossing
	weatherDatasFromVisualCrossing, err := services.FetchWeatherData(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erreur lors de la récupération des données de l'API Visual Crossing"})
		return
	}

	//Les save dans le cache
	err = cache.SaveDatasInCache(location, weatherDatasFromVisualCrossing, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//Répondre
	c.JSON(http.StatusOK, weatherDatasFromVisualCrossing)
}
