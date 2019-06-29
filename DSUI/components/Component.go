package components

import . "duckstack.com/DSFramework/DSUI/common"

type Component struct {
	ID      		string         `json:"id"`
	LazyID  		string         `json:"lazyID"`
	Actions 		Actions `json:"actions"`
	Wrapper 		map[string]*interface{} `json:"wrapper"`
	Value         	*interface{} 	`json:"value"`
	ExtProperties
	Styles
}

