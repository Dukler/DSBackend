package API

import (
	"AppointmentServer/DSInterface"
	"AppointmentServer/Helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//type urlType struct{
//	index int
//	description string
//}

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

func PostServer() {

	routerBehavior()
	//CreateJSON()
	dbtest()
	log.Fatal(
		http.ListenAndServe(
			":8080", handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST"}),
				handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
			)(router)))

}

func routerBehavior() {
	router = mux.NewRouter()
	router.HandleFunc(api+"/{entity}", SaveEndpoint).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc(api+"/{entity}"+"/{id}", ObjectEndpoint).Methods("GET", "DELETE", "OPTIONS", "POST")
	http.Handle("/", router)
}

func dbtest() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	dbh, err = Helpers.NewDBHelper(psqlInfo)
	fmt.Println(dbh.GetEntityByID(1, "appointment"))
	fmt.Println(dbh.GetEntityByID(1, "client"))

	checkErr(err)
}

func SaveEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var entity string
	entity = vars["entity"]
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		/*var obj interface{}
		err:= decoder.Decode(&obj)
		if err != nil {
			panic(err)
		}*/
		//cli.Entity = "Client"
		dbh.Save(decoder, entity)
	case "OPTIONS":
		decoder := json.NewDecoder(req.Body)
		var cli DSInterface.Client
		err := decoder.Decode(&cli)
		if err != nil {
			panic(err)
		}
	}
}

func ObjectEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var entity string
	var id int
	var err error
	id, err = strconv.Atoi(vars["id"])
	entity = vars["entity"]
	switch req.Method {
	case "GET":
		appointment, err := dbh.GetAppointmentByID(id)
		json.NewEncoder(w).Encode(&appointment)
		if err != nil {
			// handle error
			fmt.Println(err)
			fmt.Println(entity)
			os.Exit(2)
		}
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var cli DSInterface.Client
		err := decoder.Decode(&cli)
		if err != nil {
			panic(err)
		}

		//case "DELETE":
		//	err := dbobj.DeleteByID(id)
		//	if err != nil {
		//		// handle error
		//		fmt.Println(err)
		//		os.Exit(2)
		//	}
		//case "OPTIONS":
		//	err := dbobj.DeleteByID(id)
		//	if err != nil {
		//		// handle error
		//		fmt.Println(err)
		//		os.Exit(2)
		//	}
	}
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

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
