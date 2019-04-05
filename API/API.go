package API

import (
	"DuckstackBE/Helpers"
	"DuckstackBE/app"
	"DuckstackBE/controllers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var router *mux.Router
var api = "/api"
var dbh = new(Helpers.DBHelper)

//var electronicFilePath = "C:/Users/iarwa/go/src/ServerApp/Handlers/JSON/Electronics.JSON"

const (
	host     = "192.168.0.7"
	port     = 5432
	user     = "dukler"
	password = "SpiderGoku7"
	dbname   = "test"
)

func RESTApi() {

	routerBehavior()
	log.Fatal(
		http.ListenAndServe(
			":8080", handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
				handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
			)(router)))

}

func routerBehavior() {
	router = mux.NewRouter()
	router.Use(app.JwtAuthentication)
	//router.HandleFunc(api+"/Login", LogInEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	//router.HandleFunc(api+"/save/{entity}/{id}", SaveEndpoint).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc(api+"/ui/update/{Screen}", controllers.UIEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	http.Handle("/", router)
}


func testTimeString() {
	t1, err := time.Parse(
		time.RFC3339,
		"2018-03-13T00:17:41+00:00")
	fmt.Print(t1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
