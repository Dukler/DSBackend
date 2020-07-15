package data

import "os"

func GetApi(name string) string {
	var url string
	if os.Getenv("ENV") == "HEROKU" {
		url = "https://dsappserver.herokuapp.com/"
	} else {
		url = "http://localhost:8082/"
	}
	//url = "https://dsappserver.herokuapp.com/"

	return url
}
