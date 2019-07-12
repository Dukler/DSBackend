package controllers

import (
	"encoding/json"
	"fmt"
	. "duckstack.com/DSBackend/dsui/state"
	"duckstack.com/DSBackend/data"
	"net/http"
	"os"
)

var UIEndpoint = func (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			//w.Header().Add("Cache-Control", "max-age=86400")
			w.WriteHeader(http.StatusOK)
			data :=  data.GetAppJson("mockApp")
			UIState.Unmarshal(data)

			response, err := json.Marshal(UIState)
			//response = append(response,response...)
			_, responseErr := w.Write(response)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			if responseErr != nil {
				fmt.Println(err)
				os.Exit(2)
			}

		//case "POST":
		//	decoder := json.NewDecoder(req.Body)
		//	var test interface{}
		//	var uiState []byte
		//	err := decoder.Decode(&test)
		//	uiState, err = json.Marshal(&test)
		//	if err != nil {
		//		panic(err)
		//	}
		//	fmt.Println(uiState)
		//
		//case "OPTIONS":
		//	w.WriteHeader(http.StatusOK)
	}

}
