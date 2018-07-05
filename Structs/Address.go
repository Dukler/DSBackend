package Structs

type Address struct {
	Street string `json:"street"`
	Number int `json:"number"`
	Location string `json:"location"`
	Default bool `json:"default"`
}
