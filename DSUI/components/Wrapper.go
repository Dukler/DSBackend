package components

import . "duckstack.com/DSFramework/DSUI/common"

//type DSWrapper struct{
//	Wrappers map[string]Wrapper	`json:"wrappers"`
//}

type Wrapper struct{
	ID 			string			`json:"id"`
	LazyID		string 			`json:"lazyID"`
	ComponentIds
	ExtProperties
	RenderComponents
}

type RenderComponents struct {
	RenderComponents 	map[string]*interface{} `json:"renderComponents"`
}