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

func SaveDatasInCache(apiKey string, value string, expiration time.Duration) {

}
