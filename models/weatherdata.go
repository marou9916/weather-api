package models

type WeatherData struct {
	CurrentConditions struct {
		Temperature float64 `json:"temp"`
		Humidity    float64 `json:"humidity"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"currentConditions"`
}