package controllers

import (
	"DuckstackBE/DSUI"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var UIEndpoint = func (w http.ResponseWriter, req *http.Request) {
	ui := DSUI.NewUI()
	vars := mux.Vars(req)
	var screen string
	screen = vars["Screen"]
	data := getScreenJson(screen)

	nUI := ui.NormalizeData(data)

	switch req.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response, err := json.Marshal(nUI)
			w.Write(response)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

		case "POST":
			decoder := json.NewDecoder(req.Body)
			var test interface{}
			var uiState []byte
			err := decoder.Decode(&test)
			uiState, err = json.Marshal(&test)
			if err != nil {
				panic(err)
			}
			fmt.Println(uiState)

		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
	}

}

func getScreenJson(screen string) []byte {
	config := &firebase.Config{
		StorageBucket: "duckstackui.appspot.com",
	}
	path,_:=filepath.Abs("./tokens/firebase.json")
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	//test leeer bucket
	rc, err:= bucket.Object("SPA/Home.json").NewReader(context.Background())
	if err != nil {
		log.Fatalln(err)

	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatalln(err)

	}

	return slurp
}
