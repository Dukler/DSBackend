package do

import (
	"bytes"
	"net/http"
)


func Post (url string, jsonStr []byte) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	return resp
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Body:", string(body))
}