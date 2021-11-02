package cache

import (
	"encoding/json"
	"time"

	"github.com/divya2703/covid-tracker-rest-api/entity"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	expires time.Duration
}

func NewRedisCache(host string) ICache {
	return &redisCache{
		host: host,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "DJgPUVnk46QOrozQK31iDSahumyUn5JW",
		DB:       0,
	})
}

func (cache *redisCache) Set(key string, stateReport *entity.StateReport) {
	client := cache.getClient()

	// serialize Post object to JSON
	json, err := json.Marshal(stateReport)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, 100*time.Second)
}

func (cache *redisCache) Get(key string) *entity.StateReport {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	stateReport := entity.StateReport{}
	err = json.Unmarshal([]byte(val), &stateReport)
	if err != nil {
		panic(err)
	}

	return &stateReport
}
