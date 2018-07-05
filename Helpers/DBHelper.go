package Helpers

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"encoding/json"
	"log"
	"ServerApp/DSInterface"
	"ServerApp/Structs"
)

const (
	client_index = 0
	appType_index = 1

)

type DBHelper struct {
	db *sql.DB
	connection string
}

func NewDBHelper(connection string) (*DBHelper, error){
	dbh := new(DBHelper)
	dbh.connection = connection
	dbh.open()
	var err error
	return dbh, err
}

func (dbh *DBHelper) open(){
	var err error
	dbh.db, err = sql.Open("postgres",dbh.connection)
	checkErr(err)
}

func (dbh *DBHelper) GetJsonByID (id int, entity string) ([]byte,error){
	query := "SELECT data FROM " + entity + " where id = " + strconv.Itoa(id)
	rows, err := dbh.db.Query(query)
	var data []byte
	for rows.Next() {
		err = rows.Scan(&data)
	}
	return data, err
}

func (dbh *DBHelper) GetRowByID (id int, entity string) *sql.Row{
	query := "SELECT * FROM " + entity + " where id = " + strconv.Itoa(id)
	row := dbh.db.QueryRow(query)
	return row
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (dbh *DBHelper) Close (){
	defer dbh.db.Close()
}

func (dbh *DBHelper) Save(object interface{}){
	var err error
	var sqlStatement string
	var data []byte
	dbobj := object.(DSInterface.Appointment)
	switch dbobj.Entity {
	case "appointment" :
		app := object.(DSInterface.Appointment)
		if (dbh.entityExists(app.Client_id, "client")  && dbh.entityExists(app.Appointment_type_id, "appointment_type")) {
			data, err = json.Marshal(&object)
			sqlStatement = `INSERT INTO appointment (id, client_id, appointment_type_id, data)
VALUES (` + strconv.Itoa(app.ID) +` , '` + strconv.Itoa(app.Client_id) +`', '` + strconv.Itoa(app.Appointment_type_id) +`', '`+ string(data)  +`')`
			_, err = dbh.db.Exec(sqlStatement)
			if err != nil {
				panic(err)
			}
		}else {
			log.Print("client o tipo de turno no encontrado")
		}
	default:
		data, err = json.Marshal(&object)
		sqlStatement = `INSERT INTO ` + dbobj.Entity +` (id, data)
VALUES (` + strconv.Itoa(dbobj.ID) +`', '`+ string(data)  +`')`
		_, err = dbh.db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}
}


func (dbh *DBHelper) GetEntityByID (ID int, entitiy string) (interface{},error){
	//instance, err := dbh.GetJsonByID(ID,"client",dbh)
	//cli:= instance.(Client)
	//return cli, err
	switch entitiy {
	case "appointment":
		return dbh.GetAppointmentByID(ID)
	case "client":
		var cli DSInterface.Client
		data, err := dbh.GetJsonByID(ID,entitiy)
		if err := json.Unmarshal(data, &cli); err != nil {
			panic(err)
		}
		cli.ID = ID
		return cli, err
	case "appointment_type":
		var app_type Structs.Appointment_type
		data, err := dbh.GetJsonByID(ID,entitiy)
		if err := json.Unmarshal(data, &app_type); err != nil {
			panic(err)
		}
		app_type.ID = ID
		return app_type, err
	}
	return nil,nil
}

func (dbh *DBHelper) getLastID(entity string) int{
	query := "select coalesce(max(id),0) id from " + entity
	rows, err := dbh.db.Query(query)
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
	}
	if err != nil {
		panic(err)
	}
	return id
}

func (dbh *DBHelper) entityExists(id int,entity string) bool{
	query := "SELECT id FROM " + entity + " where id = " + strconv.Itoa(id)
	err := dbh.db.QueryRow(query).Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return false
	}
	return true
}

func (dbh *DBHelper) GetAppointmentByID (ID int) (DSInterface.Appointment,error){
	var app DSInterface.Appointment
	var data []byte
	var id, client_id,appointment_type_id int
	err := dbh.GetRowByID(ID,"appointment").Scan(&id,&client_id,&appointment_type_id, &data)
	if err := json.Unmarshal(data, &app); err != nil {
		panic(err)
	}
	app.ID = ID
	app.Appointment_type_id = appointment_type_id
	app.Client_id = client_id
	app.Entity = "appointment"

	return app, err
}

func (dbh *DBHelper) getJSON (ID int, entitiy string) (interface{},error){
	data, err := dbh.GetJsonByID(ID,entitiy)
	var inter interface{}
	if err := json.Unmarshal(data, &inter); err != nil {
		panic(err)
	}
	return inter, err
}
