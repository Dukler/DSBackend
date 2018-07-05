package Framework

import (
	"ServerApp/Structs"
)

type Client struct {
	DBObject
	Telephones []Structs.Telephone `json:"telephones"`
	Addresses []Structs.Address `json:"addresses"`
	Email string `json:"email"`
	Description string `json:"description"`
}