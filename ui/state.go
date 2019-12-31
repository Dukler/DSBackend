package ui

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Standalones struct {
	ByIDs map[string]*Standalone `json:"byIds"`
	IDs   []string               `json:"ids"`
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
	Standalones    `json:"standalones"`
	Containers     `json:"containers"`
	LinkList       `json:"linkList"`
	ContentRoutes  `json:"contentRoutes"`
	ComponentsPool map[string]string `json:"componentsPool"`
}

//var UIState *UI

func NewUI() *UI {
	ui := new(UI)
	ui.Standalones.ByIDs = make(map[string]*Standalone)
	ui.Containers.ByIDs = make(map[string]*Container)
	ui.LinkList.ByIDs = make(map[string]*ListedLink)
	ui.ContentRoutes.ByIDs = make(map[string]*ContentRoute)
	ui.ComponentsPool = make(map[string]string)
	return ui
}

func (ui *UI) Unmarshal(data *map[string][]interface{}) {
	for _, standalone := range (*data)["Standalones"] {
		id := standalone.(map[string]interface{})["ID"].(string)
		lazyID := ""
		if standalone.(map[string]interface{})["LazyID"] != nil {
			lazyID = standalone.(map[string]interface{})["LazyID"].(string)
		}
		ui.Standalones.ByIDs[id] = new(Standalone)
		err := FillStruct(ui.Standalones.ByIDs[id], standalone.(map[string]interface{}))
		if err != nil {
			log.Fatal(err)
		}
		ui.Standalones.IDs = append(ui.Standalones.IDs, id)
		switch {
		case strings.Contains(lazyID, "Item"):
			ui.ComponentsPool[lazyID] = "Standalone/Items"
		case strings.Contains(lazyID, "Button"):
			ui.ComponentsPool[lazyID] = "Standalone/Buttons"
		case strings.Contains(lazyID, "Picker"):
			ui.ComponentsPool[lazyID] = "Standalone/Pickers"
		case lazyID == "":
			break
		default:
			ui.ComponentsPool[lazyID] = "Standalone"
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
			ui.ComponentsPool[lazyID] = "Container"
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
