package dsui

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Components struct {
	ByIDs 			map[string]*Component `json:"byIds"`
	IDs 			[]string                `json:"ids"`
}

type Wrappers struct {
	ByIDs			map[string]*Wrapper `json:"byIds"`
	IDs 			[]string              `json:"ids"`
}

type LinkList struct{
	ByIDs    		map[string]*ListedLink `json:"byIds"`
	IDs 			[]string                 `json:"ids"`
}

type ContentRoutes struct {
	ByIDs 			map[string]*ContentRoute `json:"byIds"`
	IDs 			[]string                   `json:"ids"`
}

type UI struct {
	Components      	`json:"components"`
	Wrappers		   	`json:"wrappers"`
	LinkList    	    `json:"linkList"`
	ContentRoutes 	 	`json:"contentRoutes"`
	ComponentsPool		map[string]string	`json:"componentsPool"`
}

var UIState *UI

func NewUI() *UI{
	ui := new(UI)
	ui.Components.ByIDs = make(map[string]*Component)
	ui.Wrappers.ByIDs = make(map[string]*Wrapper)
	ui.LinkList.ByIDs = make(map[string]*ListedLink)
	ui.ContentRoutes.ByIDs = make(map[string]*ContentRoute)
	ui.ComponentsPool = make(map[string]string)
	return ui
}

func (dsui *UI) Unmarshal (data *map[string][]interface{}) {
	for _,component := range (*data)["Components"] {
		id := component.(map[string]interface{})["ID"].(string)
		lazyID := ""
		if (component.(map[string]interface{})["LazyID"] != nil){
			lazyID = component.(map[string]interface{})["LazyID"].(string)
		}
		dsui.Components.ByIDs[id] = new(Component)
		err := FillStruct(dsui.Components.ByIDs[id], component.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		dsui.Components.IDs = append(dsui.Components.IDs,id)
		switch {
		case strings.Contains(lazyID,"Item"):
			dsui.ComponentsPool[lazyID] = "Components/Items"
		case strings.Contains(lazyID,"Button"):
			dsui.ComponentsPool[lazyID] = "Components/Buttons"
		case strings.Contains(lazyID,"Picker"):
			dsui.ComponentsPool[lazyID] = "Components/Pickers"
		case lazyID == "":
			break
		default:
			dsui.ComponentsPool[lazyID] = "Components"
		}
	}
	for _,wrapper := range (*data)["Wrappers"] {
		id := wrapper.(map[string]interface{})["ID"].(string)
		lazyID := ""
		if (wrapper.(map[string]interface{})["LazyID"] != nil){
			lazyID = wrapper.(map[string]interface{})["LazyID"].(string)
		}
		dsui.Wrappers.ByIDs[id] = new(Wrapper)
		err := FillStruct(dsui.Wrappers.ByIDs[id], wrapper.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		dsui.Wrappers.IDs = append(dsui.Wrappers.IDs,id)
		if strings.Title(lazyID) == lazyID {
			dsui.ComponentsPool[lazyID]= "Wrappers"
		}
	}
	for _,listedList := range (*data)["LinkList"] {
		id := listedList.(map[string]interface{})["ID"].(string)
		dsui.LinkList.ByIDs[id] = new(ListedLink)
		err := FillStruct(dsui.LinkList.ByIDs[id], listedList.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		dsui.LinkList.IDs = append(dsui.LinkList.IDs,id)
	}
	for _,contentRoute := range (*data)["ContentRoutes"] {
		id := contentRoute.(map[string]interface{})["ID"].(string)
		dsui.ContentRoutes.ByIDs[id] = new(ContentRoute)
		err := FillStruct(dsui.ContentRoutes.ByIDs[id], contentRoute.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		dsui.ContentRoutes.IDs = append(dsui.ContentRoutes.IDs,id)
	}
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		a := val.Type().String()
		b := structFieldType.String()
		log.Print(a,b)
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func FillStruct(s interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
