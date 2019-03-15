package DSUI

type DSComponent struct {
	ID            	string `json:"id"`
	Caption       	string `json:"caption"`
	Type          	string `json:"type"`
	ComponentName   string `json:"componentName"`
	Name          	string `json:"name"`
	Suggested     	string `json:"suggested"`
	Action        	string `json:"action"`
	ContentFilter 	string `json:"contentFilter"`
	Container 		string `json:"container"`
	Components		[]*DSComponent `json:"components"`
	Value         	string `json:"value"`
}

func NewWidget(ID string, Name string, Caption string, Type string, ComponentName string, Value string) *DSComponent {
	w := new(DSComponent)
	w.ID = ID
	w.Caption = Caption
	w.Type = Type
	w.ComponentName = ComponentName
	w.Name = Name
	w.Value = Value
	return w
}
