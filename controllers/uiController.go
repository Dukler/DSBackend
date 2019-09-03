package controllers

import (
	"DSBackend/data"
	"DSBackend/ui"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var UIEndpoint = func (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			//w.Header().Add("Cache-Control", "max-age=86400")
			w.WriteHeader(http.StatusOK)
			var answer *map[string][]interface{}
			answer = data.GetApp(req)
			responseUI := ui.NewUI()
			responseUI.Unmarshal(answer)
			response, err := json.Marshal(responseUI)
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
	}

}
