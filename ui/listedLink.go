package ui

type ListedLink struct {
	ID      		string `json:"id"`
	Caption 		string `json:"caption"`
	Path    		string `json:"path"`
	Icon    		string `json:"icon"`
	ExtProperties 	map[string]interface{} `json:"extProperties"`
}
