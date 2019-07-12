package main

import (
	"DSBackend/controllers"
	"DSBackend/data"
	"DSBackend/dsui"
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
	dsui.UIState = dsui.NewUI()
	router.HandleFunc("/api/ui/update/{Screen}", controllers.UIEndpoint).Methods("GET")
	router.HandleFunc("/api/do/{Action}", controllers.DoEndpoint).Methods("POST")
	http.Handle("/", router)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)

	port :=  os.Getenv("PORT")

	data.InitBucket("duckstackui")
	//REST
	log.Fatal(http.ListenAndServe(":"+ port, handler))
}

