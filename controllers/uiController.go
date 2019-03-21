package controllers

import (
	"DuckstackBE/DSInterface/DSUI"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func UIEndpoint(w http.ResponseWriter, req *http.Request) {
	var ui DSUI.UI
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
	err := json.Unmarshal(data, &ui)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	switch req.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, err := json.Marshal(ui)
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
	}

}

func getScreenJson(screen string) string {
	var filename = "C:/Users/iarwa/Workspace/Go/src/DuckstackBE/API/" + screen + ".json"
	return filename
}
