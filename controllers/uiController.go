package controllers

import (
	"DuckstackBE/DSUI"
	. "DuckstackBE/cloudStorage"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
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
	s:= fmt.Sprintf("SPA/%s.json",screen)
	return Firebase.ReadFile(s)
}
