package Structs

type Appointment_type struct {
	ID int
	Name string
	Active bool
}

//func getAppointment_typeByID (id int,dbh *Helpers.DBHelper) (Appointment_type,error){
//	rows, err:= dbh.GetRowsByID(id,"appointment_type")
//	var dataJSON map[string]interface{}
//	var app Appointment_type
//	for rows.Next() {
//		var id int
//		var data string
//		byt := []byte(data)
//		if err := json.Unmarshal(byt, &dataJSON); err != nil {
//			panic(err)
//		}
//		err = rows.Scan(&id,&data)
//
//		//app = Appointment_type{DBObject.DBObject{id,dataJSON["name"].(string),dataJSON["active"].(bool)}}
//	}
//	return app, err
//}