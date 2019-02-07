package DSUI

type DSUI struct {
	Widgets []*DSWidget  `json:"Widgets"`
	Menu    []*DSMenu    `json:"NavMenu"`
	Content []*DSContent `json:"Content"`
}

func (dsui *DSUI) AddWidget(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *DSUI) AddButton(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *DSUI) AddMenu(menu *DSMenu) {
	dsui.Menu = append(dsui.Menu, menu)
}
