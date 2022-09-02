package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ypgyan/go-rest-api/controllers"
	"github.com/ypgyan/go-rest-api/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)

	router.HandleFunc("/api", controllers.Home).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	router.HandleFunc("/api/personalities/{id}", controllers.FindPersonality).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	router.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
