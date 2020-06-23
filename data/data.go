package data

import (
	"DSBackend/do"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

func InitBucket(bucketName string) {
	Firebase.InitBucket(bucketName)
}

func readData(data map[string]interface{}) []byte {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonStr
}

func decodeBody(body io.ReadCloser) map[string]interface{} {
	decoder := json.NewDecoder(body)
	var b map[string]interface{}
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}
	return b
}

func cleanURL(origin *string) {
	*origin = strings.TrimPrefix(*origin, "https")
	*origin = strings.TrimPrefix(*origin, "http")
	*origin = strings.TrimPrefix(*origin, "://")
	*origin = strings.TrimPrefix(*origin, "www.")
}

func GetApp(req *http.Request) *map[string][]interface{} {
	data := make(map[string]interface{})
	// url := fmt.Sprintf("%sapi/host/app",  GetApi("login"))
	url := fmt.Sprintf("%sapi/host/app", GetApi(""))
	domain := req.Header.Get("origin")
	cleanURL(&domain)

	data["domain"] = domain
	data["token"] = req.Header.Get("Authorization")

	jsonStr := readData(data)
	response := do.Post(url, jsonStr)
	body := decodeBody(response.Body)
	appName := body["appName"].(string)

	s := fmt.Sprintf("./Apps/%s", appName)
	path, _ := filepath.Abs(s)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(), "setup.config") {
			p := filepath.ToSlash(fmt.Sprintf(path+"/%s", f.Name()))
			aux := make(map[string]map[string]interface{})
			var jsonData []byte
			jsonData, err := ioutil.ReadFile(p)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(jsonData, &aux)
			if err != nil {
				log.Fatal(err)
			}
			//status, err := strconv.Atoi(response.Header.Get("status"))
			// if aux["credentials"]["login"] == true && body["loggedIn"] != true {
			// 	path,err = filepath.Abs(fmt.Sprintf("./Apps/%s", aux["login"]["appName"]))
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// }
		}
	}

	app := getAppJson(path)

	return app
}

func getAppJson(path string) *map[string][]interface{} {
	data := make(map[string][]interface{})
	err := filepath.Walk(path, mapData(&data))
	if err != nil {
		log.Fatal(err)
	}
	//if os.Getenv("APP_ENV") != "Debug" {
	//	s := fmt.Sprintf("./Apps/%s", name)
	//	path, _ := filepath.Abs(s)
	//	err := filepath.Walk(path, mapData(&data))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//} else {
	//	//s := fmt.Sprintf("SPA/%s.json", name)
	//	//data = Firebase.ReadFile(s)
	//	s := fmt.Sprintf("%s/", name)
	//
	//	a := Firebase.ReadFile(s)
	//
	//	fmt.Print(a)
	//	log.Print(s,a)
	//}
	return &data
}

func mapData(data *map[string][]interface{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		var dataError error
		if err != nil {
			log.Fatal(err)
		}
		fileExt := filepath.Ext(path)
		fileName := filepath.Base(path)
		fileNameNoExt := strings.TrimSuffix(fileName, fileExt)
		//log.Print(test)
		if fileExt == ".json" && !strings.Contains(fileNameNoExt, ".") {
			//name := trimSuffix(info.Name(), ".json")
			aux := make(map[string][]interface{})
			var jsonData []byte
			jsonData, dataError = ioutil.ReadFile(path)
			err := json.Unmarshal(jsonData, &aux)

			for uiElement, standalones := range aux {
				(*data)[uiElement] = append((*data)[uiElement], standalones...)
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
		return dataError
	}
}
