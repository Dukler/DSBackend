package ui

type Component struct {
	ID            string                 `json:"id"`
	LazyID        string                 `json:"lazyID"`
	Actions       map[string]interface{} `json:"actions"`
	Container     map[string]interface{} `json:"container"`
	Value         interface{}            `json:"value"`
	Styles        map[string]interface{} `json:"styles"`
	ExtProperties map[string]interface{} `json:"extProperties"`
}
