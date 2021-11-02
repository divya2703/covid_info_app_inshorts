package service

import (
	"errors"

	"github.com/divya2703/covid-tracker-rest-api/entity"
	"github.com/divya2703/covid-tracker-rest-api/repository"
)

type Service interface {
	FindAll() ([]entity.StateReport, error)
	FindByName(stateName string) (*entity.StateReport, error)
	FindByGeoCoordinates(*entity.LocationInfo) (*entity.StateReport, error)
}

type service struct{}

var (
	repo repository.Repository
)

func NewService(repository repository.Repository) Service {
	repo = repository
	return &service{}
}

func (*service) FindAll() ([]entity.StateReport, error) {
	return repo.FindAll()
}

func (*service) FindByName(stateName string) (*entity.StateReport, error) {

	return repo.FindByName(stateName)
}

func (*service) FindByGeoCoordinates(locationInfo *entity.LocationInfo) (*entity.StateReport, error) {
	var stateReport entity.StateReport

	if locationInfo == nil || locationInfo.Address == nil || locationInfo.Address.State == "" {
		return &stateReport, errors.New("undefined location")
	}

	return repo.FindByName(locationInfo.Address.State)

}
