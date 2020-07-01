package ui

type Standalone struct {
	ID            string                 `json:"id"`
	LazyID        string                 `json:"lazyID"`
	Actions       []interface{} 		 `json:"actions"`
	Container     map[string]interface{} `json:"container"`
	Value         interface{}            `json:"value"`
	Styles        map[string]interface{} `json:"styles"`
	Params     	  map[string]interface{} `json:"params"`
	ExtProperties map[string]interface{} `json:"extProperties"`
	Hooks		  interface{}          	 `json:"hooks"`
	SystemInfo	  system 				 `json:"systemInfo"`
}

type system struct{
	TreePosition tree `json:"treePosition"`
	SelfDestruct bool `json:"selfDestruct"`
}

type tree struct {
	StoreReducer string `json:"storeReducer"`
	ID string `json:"id"`
}