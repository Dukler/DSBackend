package DSInterface

import (
	"AppointmentServer/Structs"
	"time"
)

type Sender interface {
	Send(a Appointment) error // or whatever it needs to be
}

type Appointment struct {
	DBObject
	Date                time.Time `json:"date"`
	Client_id           int       `json:"-"`
	Appointment_type_id int       `json:"-"`
}
type Client struct {
	DBObject
	Telephones  []Structs.Telephone `json:"telephones"`
	Addresses   []Structs.Address   `json:"addresses"`
	Email       string              `json:"email"`
	Description string              `json:"description"`
}
