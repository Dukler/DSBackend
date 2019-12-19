package data

import "os"

func GetApi(name string) string {
	var url string
	if os.Getenv("APP_ENV") == "Debug" {
		url = "http://localhost:8082/"
	} else {
		url = "https://dsappserver.herokuapp.com/"
	}
	//url = "https://dsappserver.herokuapp.com/"

	return url
}
