package ui

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Components struct {
	ByIDs map[string]*Component `json:"byIds"`
	IDs   []string              `json:"ids"`
}

type Containers struct {
	ByIDs map[string]*Container `json:"byIds"`
	IDs   []string              `json:"ids"`
}

type LinkList struct {
	ByIDs map[string]*ListedLink `json:"byIds"`
	IDs   []string               `json:"ids"`
}

type ContentRoutes struct {
	ByIDs map[string]*ContentRoute `json:"byIds"`
	IDs   []string                 `json:"ids"`
}

type UI struct {
	Components     `json:"components"`
	Containers     `json:"containers"`
	LinkList       `json:"linkList"`
	ContentRoutes  `json:"contentRoutes"`
	ComponentsPool map[string]string `json:"componentsPool"`
}

//var UIState *UI

func NewUI() *UI {
	ui := new(UI)
	ui.Components.ByIDs = make(map[string]*Component)
	ui.Containers.ByIDs = make(map[string]*Container)
	ui.LinkList.ByIDs = make(map[string]*ListedLink)
	ui.ContentRoutes.ByIDs = make(map[string]*ContentRoute)
	ui.ComponentsPool = make(map[string]string)
	return ui
}

func (ui *UI) Unmarshal(data *map[string][]interface{}) {
	for _, component := range (*data)["Components"] {
		id := component.(map[string]interface{})["ID"].(string)
		lazyID := ""
		if component.(map[string]interface{})["LazyID"] != nil {
			lazyID = component.(map[string]interface{})["LazyID"].(string)
		}
		ui.Components.ByIDs[id] = new(Component)
		err := FillStruct(ui.Components.ByIDs[id], component.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		ui.Components.IDs = append(ui.Components.IDs, id)
		switch {
		case strings.Contains(lazyID, "Item"):
			ui.ComponentsPool[lazyID] = "Components/Items"
		case strings.Contains(lazyID, "Button"):
			ui.ComponentsPool[lazyID] = "Components/Buttons"
		case strings.Contains(lazyID, "Picker"):
			ui.ComponentsPool[lazyID] = "Components/Pickers"
		case lazyID == "":
			break
		default:
			ui.ComponentsPool[lazyID] = "Components"
		}
	}
	for _, container := range (*data)["Containers"] {
		id := container.(map[string]interface{})["ID"].(string)
		lazyID := ""
		if container.(map[string]interface{})["LazyID"] != nil {
			lazyID = container.(map[string]interface{})["LazyID"].(string)
		}
		ui.Containers.ByIDs[id] = new(Container)
		err := FillStruct(ui.Containers.ByIDs[id], container.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		ui.Containers.IDs = append(ui.Containers.IDs, id)
		if strings.Title(lazyID) == lazyID {
			ui.ComponentsPool[lazyID] = "Containers"
		}
	}
	for _, listedList := range (*data)["LinkList"] {
		id := listedList.(map[string]interface{})["ID"].(string)
		ui.LinkList.ByIDs[id] = new(ListedLink)
		err := FillStruct(ui.LinkList.ByIDs[id], listedList.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		ui.LinkList.IDs = append(ui.LinkList.IDs, id)
	}
	for _, contentRoute := range (*data)["ContentRoutes"] {
		id := contentRoute.(map[string]interface{})["ID"].(string)
		ui.ContentRoutes.ByIDs[id] = new(ContentRoute)
		err := FillStruct(ui.ContentRoutes.ByIDs[id], contentRoute.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		ui.ContentRoutes.IDs = append(ui.ContentRoutes.IDs, id)
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

	if structFieldType.Kind() != reflect.Interface && structFieldType != val.Type() {
		a := val.Type().String()
		b := structFieldType.String()
		log.Print(a, " ", b)
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
