package ui

//type DSContainer struct{
//	Containers map[string]Container	`json:"containers"`
//}

type Container struct {
	ID                string                 `json:"id"`
	LazyID            string                 `json:"lazyID"`
	Standalones       []interface{}          `json:"standalones"`
	ExtProperties     map[string]interface{} `json:"extProperties"`
	RenderStandalones map[string]interface{} `json:"renderStandalones"`
}
