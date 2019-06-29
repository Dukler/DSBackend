package main

import (
	. "duckstack.com/DSFramework/cloudStorage"
	"duckstack.com/DSFramework/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	router *mux.Router
)


func main() {
	router = mux.NewRouter()
	router.HandleFunc("/api/ui/update/{Screen}", controllers.UIEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	http.Handle("/", router)


	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)

	port :=  os.Getenv("PORT")

	Firebase.InitBucket("duckstackui")

	//REST
	log.Fatal(http.ListenAndServe(":"+ port, handler))
}

