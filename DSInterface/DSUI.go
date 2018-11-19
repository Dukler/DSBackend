package DSInterface

type DSUI struct {
	Widgets []*DSWidget `json:"widgets"`
	Buttons []*DSButton `json:"buttons"`
}

func (dsui *DSUI) AddWidget(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *DSUI) AddButton(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
