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
// package main

// import (
// 	"fmt"
// 	"log"
// 	"weather-api/services"
// )

// func main() {
// 	location := "Rouen" // Remplace par la ville que tu veux tester
// 	weatherData, err := services.FetchWeatherData(location)
// 	if err != nil {
// 		log.Fatalf("Erreur : %v", err)
// 	}

// 	// Affichage amélioré avec un format clair
// 	fmt.Printf("🌤️ Météo pour %s :\n", location)
// 	fmt.Printf("🌡️ Température : %.1f°C\n", weatherData.CurrentConditions.Temperature)
// 	fmt.Printf("💧 Humidité : %.1f%%\n", weatherData.CurrentConditions.Humidity)
// 	fmt.Printf("💨 Vitesse du vent : %.1f km/h\n", weatherData.CurrentConditions.WindSpeed)
// }
