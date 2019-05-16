package DSUI



//type DSWrapper struct{
//	Wrappers map[string]Wrapper	`json:"wrappers"`
//}

type DSWrapper struct{
	ID string			`json:"id"`
	ClassName string	`json:"className"`
	ComponentIds
	ExtProperties
}

