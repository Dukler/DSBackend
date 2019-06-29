package routing

import . "duckstack.com/DSFramework/DSUI/common"

type ListedLink struct {
	ID      		string `json:"id"`
	Caption 		string `json:"caption"`
	Path    		string `json:"path"`
	Icon    		string `json:"icon"`
	ExtProperties
}
