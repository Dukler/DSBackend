package components

type Component struct {
	ID      		string         `json:"id"`
	LazyID  		string         `json:"lazyID"`
	Actions 		map[string]interface{}	`json:"actions"`
	Wrapper 		map[string]interface{} `json:"wrapper"`
	Value         	interface{} 	`json:"value"`
	Styles 			map[string]interface{} `json:"styles"`
	ExtProperties 	map[string]interface{} `json:"extProperties"`
}

