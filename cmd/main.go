package main

import "weather-api/routes"

func main() {
	router := routes.SetupRouter()

	router.Run(":8080")
}