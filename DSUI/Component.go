package DSUI

type DSComponent struct {
	ID            	string 			`json:"id"`
	Caption       	string 			`json:"caption"`
	Type          	string 			`json:"type"`
	ClassName   	string 			`json:"className"`
	ClassType 		string			`json:"classType"`
	Name          	string 			`json:"name"`
	AutoComplete    string 			`json:"autoComplete"`
	Actions        	DSActions 		`json:"actions"`
	ContentFilter 	string 			`json:"contentFilter"`
	Wrapper 		string 			`json:"wrapper"`
	Value         	string 			`json:"value"`
}


