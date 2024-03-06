package api

import (
	"app/pkg/loadenv"
	"app/pkg/logs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterAPIEndpoints(apiservise *APIService) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", apiservise.Index)

	router.GET("/members", apiservise.GetAllMembersHandler)
	router.GET("/members/:id", apiservise.GetMemberHandler)
	router.POST("/members", apiservise.AddMemberHandler)
	return router
}

func StartServer(router *httprouter.Router) {
	logs.LogWriter("", "", 0)
	_, port, _ := loadenv.LoadEnv()
	http.ListenAndServe(port, router)
}
