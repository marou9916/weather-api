package controllers

import (
	"net/http"
	"weather-api/services"

	"github.com/gin-gonic/gin"
)

func LocationWeatherHandler(c *gin.Context) {
	location := c.Query("location")

	weatherData, err := services.FetchWeatherData(location, "MY_API_KEY")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, weatherData)
	
}
