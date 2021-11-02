package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/divya2703/covid-tracker-rest-api/db"
	"github.com/divya2703/covid-tracker-rest-api/entity"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

type redisCache struct {
	host     string
	password string
	expires  time.Duration
}

func NewRedisCache() ICache {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	config := db.GetConfiguration()
	log.Print("Setting host and password for redis")
	return &redisCache{
		host:     config.RedisConnectionString,
		password: config.RedisConnectionPassword,
		expires:  time.Duration(config.RedisTTL),
	}

}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.password,
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
	//log.Print("Setting expiry of keys to " + cache.expires.String() + " minutes")
	client.Set(key, json, 300*time.Second)
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
