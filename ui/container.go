package ui

//type DSContainer struct{
//	Containers map[string]Container	`json:"containers"`
//}

type Container struct {
	ID               string                 `json:"id"`
	LazyID           string                 `json:"lazyID"`
	Components       []interface{}          `json:"components"`
	ExtProperties    map[string]interface{} `json:"extProperties"`
	RenderComponents map[string]interface{} `json:"renderComponents"`
}
