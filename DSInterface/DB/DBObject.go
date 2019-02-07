package DSdbhelper

type DBObject struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	//Entity string `json:"-"`
}

//func (dbobj *DBObject) New() interface{}{
//	switch dbobj.Entity {
//	case "appointment":
//		var app Appointment
//		return app
//	}
//	return nil
//}
