package DSUI

type UI struct {
	Widgets []*DSWidget  `json:"Widgets"`
	Menu    []*DSMenu    `json:"NavMenu"`
	Content []*DSContent `json:"Content"`
}

func (dsui *UI) AddWidget(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *UI) AddButton(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *UI) AddMenu(menu *DSMenu) {
	dsui.Menu = append(dsui.Menu, menu)
}
