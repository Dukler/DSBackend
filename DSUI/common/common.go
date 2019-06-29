package common


type ComponentIds struct {
	ComponentIds  []string	`json:"components"`
}

type ExtProperties struct {
	ExtProperties 	map[string]*interface{} `json:"extProperties"`
}

type Styles struct {
	Name	string 	 `json:"name"`
	Styles 	map[string]*interface{} `json:"styles"`
}
