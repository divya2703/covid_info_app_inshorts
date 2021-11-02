package cache

import "github.com/divya2703/covid-tracker-rest-api/entity"

type ICache interface {
	Set(key string, value *entity.StateReport)
	Get(key string) *entity.StateReport
}
