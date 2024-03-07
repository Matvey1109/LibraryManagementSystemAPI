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
	router.GET("/members/:memberID", apiservise.GetMemberHandler)
	router.POST("/members", apiservise.AddMemberHandler)
	router.PUT("/members/:memberID", apiservise.UpdateMemberHandler)
	router.DELETE("/members/:memberID", apiservise.DeleteMemberHandler)

	router.GET("/books", apiservise.GetAllBooksHandler)
	router.GET("/books/:bookID", apiservise.GetBookHandler)
	router.POST("/books", apiservise.AddBookHandler)
	router.PUT("/books/:bookID", apiservise.UpdateBookHandler)
	router.DELETE("/books/:bookID", apiservise.DeleteBookHandler)
	return router
}

func StartServer(router *httprouter.Router) {
	logs.LogWriter("", "", 0)
	_, port, _ := loadenv.LoadEnv()
	http.ListenAndServe(port, router)
}
