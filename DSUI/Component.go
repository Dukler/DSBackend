package DSUI

type DSComponent struct {
	ID            	string 			`json:"id"`
	ClassName   	string 			`json:"className"`
	Actions        	DSActions 		`json:"actions"`
	Wrapper
	Value         	string 			`json:"value"`
	ExtProperties
}

type Wrapper struct{
	Wrapper 	map[string]*interface{} `json:"wrapper"`
}
