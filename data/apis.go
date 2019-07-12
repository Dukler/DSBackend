package data

import "os"

func GetApi(name string) string {
	var url string
	if (os.Getenv("APP_ENV") == "Debug"){
		url = "http://localhost:8082/"
	}else{
		switch name {
		case "login":
			url = "https://dsloginserver.herokuapp.com/"
			break
		default:
			break
		}
	}

	return url
}
