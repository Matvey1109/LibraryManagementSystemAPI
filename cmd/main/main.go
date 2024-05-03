package main

import (
	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/api"
	"github.com/Matvey1109/LibraryManagementSystemCore/pkg/logs"
)

func main() {
	defer logs.CloseLogFile()
	apiservise := api.NewAPIService()
	handler := api.RegisterAPIEndpoints(apiservise)
	api.StartServer(handler)
}
