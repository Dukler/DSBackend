package actions

import (
	"encoding/json"
	"fmt"
	d "DSBackend/data"
	"DSBackend/do"
	"io"
)

//type user struct {
//	Email string `json:"email"`
//	Password string `json:"password"`
//}

func readData (data map[string]interface{}) []byte{
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonStr
}

func decodeBody(body io.ReadCloser) map[string]interface{}{
	decoder := json.NewDecoder(body)
	var b map[string]interface{}
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}
	return b
}

func Login (data map[string]interface{}) *interface{}{
	//usr := new(user)
	//usr.Email = data["email"].(string)
	//usr.Password = data["password"].(string)
	url := fmt.Sprintf("%sapi/user/login",  d.GetApi("login"))
	jsonStr := readData(data)
	response := do.Post(url,jsonStr)
	body := decodeBody(response.Body)
	token := body["user"].(map[string]interface{})["token"]
	//log.Print(token)
	return &token
}

func CreateUser (data map[string]interface{}) *interface{}{
	url := fmt.Sprintf("%sapi/user/new",  d.GetApi("login"))
	jsonStr := readData(data)
	response := do.Post(url,jsonStr)
	var body interface{}
	body = decodeBody(response.Body)
	//token := body["user"].(map[string]ui{})["token"].(string)
	//log.Print(token)
	return &body
}