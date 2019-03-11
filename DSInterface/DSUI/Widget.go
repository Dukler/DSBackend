package DSUI

type DSWidget struct {
	ID            string `json:"id"`
	Caption       string `json:"caption"`
	Type          string `json:"type"`
	DSComponent   string `json:"dscomponent"`
	Name          string `json:"name"`
	Suggested     string `json:"suggested"`
	Action        string `json:"action"`
	ContentFilter string `json:"contentFilter"`
	Value         string `json:"value"`
}

func NewWidget(ID string, Name string, Caption string, Type string, DSComponent string, Value string) *DSWidget {
	w := new(DSWidget)
	w.ID = ID
	w.Caption = Caption
	w.Type = Type
	w.DSComponent = DSComponent
	w.Name = Name
	w.Value = Value
	return w
}
