package DSUI

type DSActions struct {
	Actions map[string][]Action	`json:"actions"`
}

type Action struct {
	Action string `json:"action"`
	Params map[string][]string	`json:"params"`
}