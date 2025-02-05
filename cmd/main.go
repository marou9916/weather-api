package main

import "weather-api/routes"

func main() {
	router := routes.SetupRoutes()

	router.Run(":8080")
}