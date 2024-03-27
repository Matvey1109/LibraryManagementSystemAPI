package main

import (
	"github.com/Matvey1109/LibraryManagementSystemAPI/internal/console"
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
