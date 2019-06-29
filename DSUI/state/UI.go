package state

import (
	. "duckstack.com/DSFramework/DSUI/components"
	. "duckstack.com/DSFramework/DSUI/routing"
	"encoding/json"
	"fmt"
	"os"
)

type UI struct {
	Components 		[]*Component      	`json:"components"`
	Wrappers		[]*Wrapper    		`json:"wrappers"`
	LinkList    	[]*ListedLink    	`json:"linkList"`
	ContentRoutes 	[]*ContentRoute 	`json:"contentRoutes"`
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
