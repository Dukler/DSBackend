package controllers

import (
	"DSBackend/actions"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

var DoEndpoint = func (w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		//vars := mux.Vars(req)
		action := chi.URLParam(req, "Action")
		decoder := json.NewDecoder(req.Body)
		var payload map[string]interface{}
		err := decoder.Decode(&payload)
		if err != nil {
			panic(err)
		}
		answer := actions.Dispatch(action,payload)
		w.WriteHeader(http.StatusOK)
		response, err := json.Marshal(answer)
		_, responseErr := w.Write(response)
		if responseErr != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}

}

