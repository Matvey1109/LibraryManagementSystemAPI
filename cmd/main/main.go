package main

import (
	"app/internal/console"
)

func main() {
	// defer logs.CloseLogFile()
	// apiservise := api.NewAPIService()
	// handler := api.RegisterAPIEndpoints(apiservise)
	// api.StartServer(handler)
	console.TestMember()
	console.TestBook()
	console.TestBorrowing()
}
