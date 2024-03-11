package main

import (
	"app/internal/api"
	"app/pkg/logs"
)

func main() {
	defer logs.CloseLogFile()
	apiservise := api.NewAPIService()
	handler := api.RegisterAPIEndpoints(apiservise)
	api.StartServer(handler)
}
