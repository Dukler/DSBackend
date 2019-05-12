package DSUI

import "strings"

func (dsui *UI) NormalizeData (data []byte) *NUI{
	var nUI = NewNUI()
	//var cmpIds 	[]string
	//var wrIds	[]string
	//var llIds 	[]string
	//var crIds 	[]string
	componentPool := make(map[string]string)
	dsui.Unmarshal(data)

	for _,component := range dsui.Components {
		nUI.Components.ByIDs[component.ID] = component
		nUI.Components.IDs = append(nUI.Components.IDs,component.ID)
		//cmpIds = append(cmpIds,component.ID)
		//componentPool = append(componentPool,component.ClassName)
		componentPool[component.ClassName]= "Components"
	}
	for _,wrapper := range dsui.Wrappers {
		nUI.Wrappers.ByIDs[wrapper.ID] = wrapper
		nUI.Wrappers.IDs = append(nUI.Wrappers.IDs,wrapper.ID)
		//wrIds = append(wrIds,wrapper.ID)
		if(strings.Title(wrapper.ClassName) == wrapper.ClassName){
			componentPool[wrapper.ClassName]= "Wrappers"
		}
	}
	for _,listedList := range dsui.LinkList {
		nUI.LinkList.ByIDs[listedList.ID] = listedList
		nUI.LinkList.IDs = append(nUI.LinkList.IDs,listedList.ID)
		//llIds = append(llIds,listedList.ID)
	}
	for _,contentRoute := range dsui.ContentRoutes {
		nUI.ContentRoutes.ByIDs[contentRoute.ID] = contentRoute
		nUI.ContentRoutes.IDs = append(nUI.ContentRoutes.IDs,contentRoute.ID)
		//crIds = append(crIds,contentRoute.ID)
	}

	//nUI.Components.IDs=cmpIds
	//nUI.Wrappers.IDs=wrIds
	//nUI.LinkList.IDs=llIds
	//nUI.ContentRoutes.IDs=crIds
	delete(componentPool, "")
	nUI.ComponentsPool=componentPool

	return nUI
}

type NUI struct{
	Components 		nComponents			`json:"components"`
	Wrappers		nWrappers			`json:"wrappers"`
	LinkList    	nLinkList			`json:"linkList"`
	ContentRoutes 	nContentRoute		`json:"contentRoutes"`
	ComponentsPool	map[string]string	`json:"componentsPool"`
}

type nComponents struct {
	ByIDs 			map[string]*DSComponent		`json:"byIds"`
	IDs 			[]string 					`json:"ids"`
}

type nWrappers struct {
	ByIDs			map[string]*DSWrapper		`json:"byIds"`
	IDs 			[]string 					`json:"ids"`
}

type nLinkList struct{
	ByIDs    		map[string]*DSListedLink	`json:"byIds"`
	IDs 			[]string 					`json:"ids"`
}

type nContentRoute struct {
	ByIDs 			map[string]*DSContentRoute	`json:"byIds"`
	IDs 			[]string 					`json:"ids"`
}

func  NewNUI() *NUI {
	nUI := new(NUI)
	nUI.Components.ByIDs = make(map[string]*DSComponent)
	nUI.Wrappers.ByIDs = make(map[string]*DSWrapper)
	nUI.LinkList.ByIDs = make(map[string]*DSListedLink)
	nUI.ContentRoutes.ByIDs = make(map[string]*DSContentRoute)
	nUI.ComponentsPool = make(map[string]string)
	return nUI
}