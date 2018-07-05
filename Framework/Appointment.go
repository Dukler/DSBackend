package Framework

import (
	"time"
	"ServerApp/DSInterface"
)

type Appointment struct {
	DBObject
	Date time.Time `json:"date"`
	Client_id int `json:"-"`
	Appointment_type_id int `json:"-"`
}
func NewAppointment() *Appointment {
	return new(Appointment) // or whatever
}
func (app *Appointment) Send(a DSInterface.Appointment) error {
	var err error
	// send twitch message...

	// Need to send a Discord message?
	app.Send(a)
	return err
}


func getFullAppointment (){

}
func SetAppointmentByID()(){

}