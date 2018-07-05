package Structs

import (
	"encoding/xml"
)

type Indoor struct {
	XMLName xml.Name `xml:"Indoor"`
	//Version string   `xml:"version,attr"`
	Components     []Component `xml:"Component"`
}

type Component struct {
	ID   int `xml:"ID"`
	Name string `xml:"Name"`
	Value float32 `xml:"Value"`
	Active bool `xml:"Active"`
}
