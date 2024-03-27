package api

import (
	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/repositories"
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
