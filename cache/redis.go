package cache

import (
	"encoding/json"
	"fmt"
	"time"
	"weather-api/configs"
	"weather-api/models"

	"github.com/redis/go-redis/v9"
)

// GetDatasFromCache récupère les données météorologiques du cache
func GetDatasFromCache(location string) (*models.WeatherData, error) {
	//Récupérer la clé pour ce lieu
	keyForAskedLocation := fmt.Sprintf("weather:%s", location)

	//Vérifier si la clé est présente dans le cache
	weatherDatasFromCache, err := configs.RedisClient.Get(configs.Ctx, keyForAskedLocation).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	//Désérialiser les données JSON du cache
	var locationWeatherDatas models.WeatherData

	if err := json.Unmarshal([]byte(weatherDatasFromCache), &locationWeatherDatas); err != nil {
		return nil, fmt.Errorf("erreur de désérialisation du cache: %v", err)
	}

	return &locationWeatherDatas, nil
}

// SaveDatasInCache stocke les données météo dans Redis sous format JSON avec une expiration
func SaveDatasInCache(location string, value *models.WeatherData, expiration time.Duration) error {
	key := fmt.Sprintf("weather:%s", location)

	weatherDatasJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde des données dans le cache: %v", err)
	}

	err = configs.RedisClient.Set(configs.Ctx, key, weatherDatasJSON, expiration).Err()
	if err != nil {
		return fmt.Errorf("erreur lors de l'enregistrement des données dans le cache: %v", err) 
	}

	return nil
}
