package repository

import "github.com/divya2703/covid-tracker-rest-api/entity"

type Repository interface {
	FindByName(stateName string) (*entity.StateReport, error)
	FindAll() ([]entity.StateReport, error)
}
