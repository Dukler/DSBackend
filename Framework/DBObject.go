package Framework

type DBObject struct{
	ID int `json:"-"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Entity string `json:"-"`
}

func (dbobj *DBObject) New() interface{}{
	switch dbobj.Entity {
	case "appointment":
		var app Appointment
		return app
	}
	return nil
}
//func getJSONBYID(ID int, entity string, dbh *Helpers.DBHelper)	([]byte,error){
//	rows, err:= dbh.GetRowsByID(ID,entity)
//	var data []byte
//	for rows.Next() {
//		err = rows.Scan(&data)
//	}
//	return data, err
//}

//func getBYID(ID int, entity string, dbh *Helpers.DBHelper)	(interface{} ,error){
//	data, err:=  dbh.GetJsonByID(ID,entity)
//	var instance interface{}
//	if err := json.Unmarshal(data, &instance); err != nil {
//		panic(err)
//	}
//	//instance.(DBObject).ID = ID
//	var asd = DBObject{instance}
//	asd.ID = ID
//	return instance, err
//}