package JSON

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
	"ServerApp/Structs/DBObject"
	"ServerApp/Structs/Client"
)

func CreateJSON(){
	//v:=Appointment.Appointment{DBObject.DBObject{1,"EarthHumidity",true},time.Now()}
	//pagesJson, err := json.Marshal(v)
	//os.Stdout.Write(pagesJson)
	//if err != nil {
	//	log.Fatal("Cannot encode to JSON ", err)
	//}
}
func TestClientJSON() {
	var t []Client.Telephone
	var a []Client.Address
	t = append(t, Client.Telephone{1556960610,"movil",true})
	t = append(t,Client.Telephone{46390945,"casa",false})
	a = append(a, Client.Address{"marcos sastre",5790,"devoto",true})
	v:=Client.Client{DBObject.DBObject{1,"Dukler",true},t,a,"8amartin@gmail.com","god"}
	pagesJson, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	os.Stdout.Write(pagesJson)
}

func TestAppointmentJSON() {
	var t []Client.Telephone
	var a []Client.Address
	t = append(t, Client.Telephone{1556960610,"movil",true})
	t = append(t,Client.Telephone{46390945,"casa",false})
	a = append(a, Client.Address{"marcos sastre",5790,"devoto",true})
	v:=Client.Client{DBObject.DBObject{1,"Dukler",true},t,a,"8amartin@gmail.com","god"}
	pagesJson, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	os.Stdout.Write(pagesJson)
}

func LoadJSONfromFile(path string, class interface{}){
	v, err := os.Open(path)
	json.NewDecoder(v).Decode(&class)
	v.Close()
	if err != nil {
		fmt.Println(err)
	}
}