package DSUI

type DSComponent struct {
	ID            	string 			`json:"id"`
	LazyID		   	string 			`json:"lazyID"`
	Actions        	DSActions 		`json:"actions"`
	Wrapper
	Value         	*interface{} 	`json:"value"`
	ValueProps		`json:"valueProps"`
	ExtProperties
	Styles
}

type Wrapper struct{
	Wrapper 	map[string]*interface{} `json:"wrapper"`
}
type ValueProps struct{
	IsHtml	bool	`json:"isHtml"`
}