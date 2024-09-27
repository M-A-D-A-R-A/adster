package service

import (
	"microservices/src/entity"
	"microservices/src/input"
	"microservices/src/repository"
)

type TargetService interface {
	GetForcast(forecastData input.ForecastRequest) (entity.ForecastData, error)
	GetAnalytics() ([]entity.ForecastData, error)
	// sendForecastRequest(forecastData input.ForecastRequest) ([]entity.ForecastData, error)
}

type targetService struct {
	repository repository.TargetRepository
}

func NewTargetService(repository repository.TargetRepository) *targetService {
	return &targetService{repository}
}

func (s *targetService) GetForcast(forecastData input.ForecastRequest) (entity.ForecastData, error) {
	result, err := s.repository.GetForcast(forecastData)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *targetService) GetAnalytics() ([]entity.ForecastData, error) {
	result, err := s.repository.GetAnalytics()
	if err != nil {
		return nil, err
	}

	return result, nil
}

