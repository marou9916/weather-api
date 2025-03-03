package main

import (
	"weather-api/configs"
	"weather-api/routes"
)

func main() {
	router := routes.SetupRoutes()

	configs.InitRedis()

	router.Run(":8080")
}
