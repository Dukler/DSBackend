package DSUI



//type DSWrapper struct{
//	Wrappers map[string]Wrapper	`json:"wrappers"`
//}

type DSWrapper struct{
	ID string			`json:"id"`
	ClassName string	`json:"className"`
	ClassType string	`json:"classType"`
	ComponentIds
	LinkList []string	`json:"linkList"`
	Attributes
}

