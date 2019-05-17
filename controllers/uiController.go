package controllers

import (
	"DuckstackBE/DSUI"
	. "DuckstackBE/cloudStorage"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var UIEndpoint = func (w http.ResponseWriter, req *http.Request) {
	ui := DSUI.NewUI()
	vars := mux.Vars(req)
	var screen string
	screen = vars["Screen"]


	switch req.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			w.Header().Add("Cache-Control", "max-age=86400")
			w.WriteHeader(http.StatusOK)
			data := getScreenJson(screen)
			nUI := ui.NormalizeData(data)
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
	var data []byte
	if (os.Getenv("APP_ENV")=="Debug"){
		s:= fmt.Sprintf("./Screens/%s.json",screen)
		path,_ := filepath.Abs(s)
		data,_ = ioutil.ReadFile(path)
	}else{
		s:= fmt.Sprintf("SPA/%s.json",screen)
		data = Firebase.ReadFile(s)
	}
	return data
}
