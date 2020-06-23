package ui

//type DSContainer struct{
//	Containers map[string]Container	`json:"containers"`
//}

type Container struct {
	ID                string                 `json:"id"`
	LazyID            string                 `json:"lazyID"`
	Components 		  []interface{}          `json:"components"`
	Params     	  	  map[string]interface{} `json:"params"`
	ExtProperties     map[string]interface{} `json:"extProperties"`
	RenderStandalones map[string]interface{} `json:"renderStandalones"`
	Hooks			  interface{}          	 `json:"hooks"`
}
