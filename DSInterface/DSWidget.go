package DSInterface

type DSWidget struct {
	ID            string `json:"id"`
	Caption       string `json:"caption"`
	Type          string `json:"type"`
	DSType        string `json:"dstype"`
	Name          string `json:"name"`
	Suggested     string `json:"suggested"`
	Action        string `json:"action"`
	ContentFilter string `json:"contentFilter"`
	Value         string `json:"value"`
}

func NewWidget(ID string, Name string, Caption string, Type string, DSType string, Value string) *DSWidget {
	w := new(DSWidget)
	w.ID = ID
	w.Caption = Caption
	w.Type = Type
	w.DSType = DSType
	w.Name = Name
	w.Value = Value
	return w
}
