package DSUI

import (
	"encoding/json"
	"fmt"
	"os"
)

type ComponentIds struct {
	ComponentIds  []string	`json:"components"`
}

type ExtProperties struct {
	ExtProperties 	map[string]*interface{} `json:"extProperties"`
}

type UI struct {
	Components 		[]*DSComponent		`json:"components"`
	Wrappers		[]*DSWrapper		`json:"wrappers"`
	LinkList    	[]*DSListedLink		`json:"linkList"`
	ContentRoutes 	[]*DSContentRoute	`json:"contentRoutes"`
}

func NewUI() *UI{
	ui := new(UI)
	return ui
}

func(dsui *UI) Unmarshal(data []byte){
	err := json.Unmarshal(data, &dsui)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}