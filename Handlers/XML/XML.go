package XML


import (
	"encoding/xml"
	"fmt"
	. "ServerApp/Structs"
)

//var postedXML []byte

func XMLHandler(){

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func CreateXML() [] byte{
	// xml
	v := &Indoor{}
	v.Components = append(v.Components, Component{1,"EarthHumidity", 10.1,true})
	v.Components = append(v.Components, Component{2, "WaterPump", 1,true})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output

}