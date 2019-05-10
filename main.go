package main

import (
	"DuckstackBE/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var router *mux.Router

func main() {
	routerBehavior()
	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)

	port :=  os.Getenv("PORT")
	if os.Getenv("APP_ENV")== "Debug"{
		port = "8081"
	}


	log.Fatal(http.ListenAndServe(":"+ port, handler))
}

func routerBehavior() {
	router = mux.NewRouter()
	router.HandleFunc("/api/ui/update/{Screen}", controllers.UIEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	http.Handle("/", router)
}
