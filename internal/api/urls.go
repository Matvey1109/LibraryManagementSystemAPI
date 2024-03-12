package api

import (
	"app/internal/swagger"
	"app/pkg/loadenv"
	"app/pkg/logs"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func RegisterAPIEndpoints(apiservise *APIService) http.Handler {
	router := httprouter.New()
	router.GET("/", apiservise.Index)

	router.GET("/swagger", swagger.SwaggerHandler)

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

	// Serve static files
	fs := http.FileServer(http.Dir("internal/swagger/static"))
	router.GET("/styles.css", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fs.ServeHTTP(w, r)
	})
	router.GET("/app.js", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fs.ServeHTTP(w, r)
	})

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	return c.Handler(router)
}

func StartServer(handler http.Handler) {
	logs.LogWriter("", "", 0)
	_, port, _ := loadenv.LoadEnv()
	http.ListenAndServe(port, handler)
}
