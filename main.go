package main

import (
	"DuckstackBE/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var router *mux.Router
var api = "/api"

func main() {
	routerBehavior()
	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)
	log.Fatal(http.ListenAndServe(":8081", handler))
}

func routerBehavior() {
	router = mux.NewRouter()
	router.HandleFunc(api+"/ui/update/{Screen}", controllers.UIEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	http.Handle("/", router)
}
