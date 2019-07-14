package ui

//type DSWrapper struct{
//	Wrappers map[string]Wrapper	`json:"wrappers"`
//}

type Wrapper struct{
	ID 					string			`json:"id"`
	LazyID				string 			`json:"lazyID"`
	Components  		[]interface{}		`json:"components"`
	ExtProperties 		map[string]interface{} `json:"extProperties"`
	RenderComponents 	map[string]interface{} `json:"renderComponents"`
}

