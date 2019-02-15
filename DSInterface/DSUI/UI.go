package DSUI

type UI struct {
	Widgets []*DSWidget  `json:"Widgets"`
	LinkList    []*DSLinkList    `json:"LinkList"`
	Content []*DSContent `json:"Content"`
}

func (dsui *UI) AddWidget(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *UI) AddButton(widget *DSWidget) {
	dsui.Widgets = append(dsui.Widgets, widget)
}
func (dsui *UI) AddMenu(menu *DSLinkList) {
	dsui.LinkList = append(dsui.LinkList, menu)
}
