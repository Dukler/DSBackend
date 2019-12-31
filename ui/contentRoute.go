package ui

type ContentRoute struct {
	ID          string        `json:"id"`
	Path        string        `json:"path"`
	Caption     string        `json:"caption"`
	ClassName   string        `json:"className"`
	Standalones []interface{} `json:"standalones"`
}
