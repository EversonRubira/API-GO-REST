package routes

import (
	controllers "API-GO-REST/Controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/id", controllers.TodasPersonalidades).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
