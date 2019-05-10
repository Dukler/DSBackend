package controllers

import (
	"DuckstackBE/DSInterface/DSUI"
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
	if _, err := os.Stat(getScreenJson(screen)); os.IsNotExist(err) {
		return
	}
	data, e := ioutil.ReadFile(getScreenJson(screen))
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}

	ui.Unmarshal(data)

	nUI := ui.NormalizeData()

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

func getScreenJson(screen string) string {
	//var filename = "C:/Users/iarwa/Workspace/Go/src/DuckstackBE/Screens/" + screen + ".json"
	filename, err := filepath.Abs("./Screens/" + screen + ".json")
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return filename
}
