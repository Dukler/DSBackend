package DSUI

type UI struct {
	Components 	[]*DSComponent `json:"Components"`
	LinkList    []*DSListedLink  `json:"LinkList"`
	Content 	[]*DSContentRoute `json:"Content"`
}

func (dsui *UI) AddComponent(widget *DSComponent) {
	dsui.Components = append(dsui.Components, widget)
}
func (dsui *UI) AddLink(menu *DSListedLink) {
	dsui.LinkList = append(dsui.LinkList, menu)
}
