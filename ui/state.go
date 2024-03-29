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
	for _, c := range (*data)["Containers"] {
		container := c.(map[string]interface{});
		// log.Print(c)
		id := container["ID"].(string)
		lazyID := ""
		if container["LazyID"] != nil {
			lazyID = container["LazyID"].(string)
		}
		ui.Containers.ByIDs[id] = new(Container)
		err := FillStruct(ui.Containers.ByIDs[id], container)
		if err != nil {
			log.Fatal(err)
		}
		ui.Containers.IDs = append(ui.Containers.IDs, id)
		if strings.Title(lazyID) == lazyID {
			ui.ComponentsPool[lazyID] = "Container"
		}
		setSysInfo(container["ID"].(string),"containers",container["Components"].([]interface{}), ui);
	}
	for _, l := range (*data)["LinkList"] {
		listedList := l.(map[string]interface{})
		id := listedList["ID"].(string)
		ui.LinkList.ByIDs[id] = new(ListedLink)
		err := FillStruct(ui.LinkList.ByIDs[id], listedList)
		if err != nil {
			log.Fatal(err)
		}
		ui.LinkList.IDs = append(ui.LinkList.IDs, id)
	}
	for _, r := range (*data)["ContentRoutes"] {
		contentRoute := r.(map[string]interface{})
		id := contentRoute["ID"].(string)
		ui.ContentRoutes.ByIDs[id] = new(ContentRoute)
		err := FillStruct(ui.ContentRoutes.ByIDs[id], contentRoute)
		if err != nil {
			log.Fatal(err)
		}
		ui.ContentRoutes.IDs = append(ui.ContentRoutes.IDs, id)

		setSysInfo(contentRoute["ID"].(string),"contentRoutes",contentRoute["Components"].([]interface{}), ui);
	}
}

func setSysInfo (containerID string, storeReducer string, arr []interface{}, ui *UI){
	for _, c := range arr{
		if _, ok := ui.Standalones.ByIDs[c.(string)]; ok {
			comp := c.(string)
			ui.Standalones.ByIDs[comp].SystemInfo.TreePosition.StoreReducer = storeReducer;
			ui.Standalones.ByIDs[comp].SystemInfo.TreePosition.ID = containerID;
			// log.Print(comp);
		}
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
