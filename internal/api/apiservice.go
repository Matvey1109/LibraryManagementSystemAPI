package api

import (
	"app/internal/repositories"
)

type APIService struct{}

var (
	instance   *APIService
	repository = repositories.ExportRepository
)

func NewAPIService() *APIService {
	if instance == nil {
		instance = &APIService{}
	}
	return instance
}
