package DSInterface

type DSUI struct {
	Widgets []*DSWidget `json:"widgets"`
}

func (dsui *DSUI) AddWidget(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
