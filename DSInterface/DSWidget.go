package DSInterface

type DSWidget struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewWidget(ID string, Name string, Label string, Type string, Value string) *DSWidget {
	w := new(DSWidget)
	w.ID = ID
	w.Label = Label
	w.Type = Type
	w.Name = Name
	w.Value = Value
	return w
}
