package data

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func InitBucket (bucketName string) {
	Firebase.InitBucket(bucketName)
}

func GetAppJson(name string) *map[string][]interface{} {
	data := make(map[string][]interface{})
	if os.Getenv("APP_ENV") == "NotDebug" {
		s := fmt.Sprintf("./Apps/%s", name)
		path, _ := filepath.Abs(s)
		err := filepath.Walk(path, mapData(&data))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//s := fmt.Sprintf("SPA/%s.json", name)
		//data = Firebase.ReadFile(s)
		s := fmt.Sprintf("%s/", name)

		a := Firebase.ReadFile(s)

		fmt.Print(a)
		log.Print(s,a)
	}
	return &data
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func mapData(data *map[string][]interface{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		var dataError error
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) == ".json" {
			//name := trimSuffix(info.Name(), ".json")
			aux := make(map[string][]interface{})
			var jsonData []byte
			jsonData, dataError = ioutil.ReadFile(path)
			err := json.Unmarshal(jsonData, &aux)

			for uiElement, components := range aux {
				(*data)[uiElement] = append((*data)[uiElement], components...)
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		return dataError
	}
}
