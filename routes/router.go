package routes

import (
	"weather-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.LocationWeatherHandler)

	return router
}
