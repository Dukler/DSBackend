package DSUI

type DSComponent struct {
	ID            	string 			`json:"id"`
	LazyID		   	string 			`json:"lazyID"`
	Actions        	DSActions 		`json:"actions"`
	Wrapper
	Value         	*interface{} 	`json:"value"`
	ExtProperties
	Styles
}

type Wrapper struct{
	Wrapper 	map[string]*interface{} `json:"wrapper"`
}