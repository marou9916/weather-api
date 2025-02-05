package routes

import (
	"weather-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.LocationWeatherDatasHandler)

	return router
}
