package main

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

const cert = "C:/Users/iarwa/Workspace/Go/src/DuckstackBE/cert.pem"
const key = "C:/Users/iarwa/Workspace/Go/src/DuckstackBE/key.pem"
const build = "C:/Users/iarwa/Workspace/React/duckstackui/build"

const (
	host     = "192.168.0.7"
	port     = 5432
	user     = "dukler"
	password = "SpiderGoku7"
	dbname   = "test"
)



func main() {
	routerBehavior()
	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)

	//go func(){
	//	log.Fatal(http.ListenAndServe(":8082", http.HandlerFunc(redirectToHttps)))
	//}()

	//http
	log.Fatal(http.ListenAndServe(":8081", handler))
	//log.Fatal(http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", handler))





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


func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	print("dickballs")
	http.Redirect(w, r, "https://192.168.0.2:8081"+r.RequestURI, http.StatusMovedPermanently)
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
