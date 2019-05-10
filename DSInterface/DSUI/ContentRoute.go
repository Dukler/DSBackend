package DSUI

type DSContentRoute struct {
	ID      	string 	`json:"id"`
	Path    	string 	`json:"path"`
	Caption 	string 	`json:"caption"`
	ClassName 	string 	`json:"className"`
	*ComponentIds
}
