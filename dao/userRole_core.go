package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/userrole.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : UserRole (userrole)
// Endpoint 	        : UserRole (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:59
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// UserRole_GetList() returns a list of all UserRole records
func UserRole_GetList() (int, []dm.UserRole, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.UserRole_SQLTable)
	count, userroleList, _, _ := userrole_Fetch(tsql)

	return count, userroleList, nil
}

// UserRole_GetLookup() returns a lookup list of all UserRole items in lookup format
func UserRole_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, userroleList, _ := UserRole_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: userroleList[i].Id, Name: userroleList[i].Name})
	}
	return returnList
}

// UserRole_GetByID() returns a single UserRole record
func UserRole_GetByID(id string) (int, dm.UserRole, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.UserRole_SQLTable)
	tsql = tsql + " WHERE " + dm.UserRole_SQLSearchID + "='" + id + "'"
	_, _, userroleItem, _ := userrole_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, userroleItem, nil
}

// UserRole_DeleteByID() deletes a single UserRole record
func UserRole_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.UserRole_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.UserRole_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// UserRole_Store() saves/stores a UserRole record to the database
func UserRole_Store(r dm.UserRole, req *http.Request) error {

	err, r := UserRole_Validate(r)
	if err == nil {
		err = userrole_Save(r, Audit_User(req))
	} else {
		logs.Information("UserRole_Store()", err.Error())
	}

	return err
}

// UserRole_StoreSystem() saves/stores a UserRole record to the database
func UserRole_StoreSystem(r dm.UserRole) error {

	err, r := UserRole_Validate(r)
	if err == nil {
		err = userrole_Save(r, Audit_Host())
	} else {
		logs.Information("UserRole_Store()", err.Error())
	}

	return err
}

// UserRole_Validate() validates for saves/stores a UserRole record to the database
func UserRole_Validate(r dm.UserRole) (error, dm.UserRole) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// userrole_Save() saves/stores a UserRole record to the database
func userrole_Save(r dm.UserRole, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = UserRole_NewID(r)
	}

	// If there are fields below, create the methods in dao\userrole_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("UserRole", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.UserRole_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.UserRole_Id_sql, r.Id)
	ts = addData(ts, dm.UserRole_Name_sql, r.Name)
	ts = addData(ts, dm.UserRole_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.UserRole_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.UserRole_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.UserRole_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.UserRole_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.UserRole_SYSCreated_sql, r.SYSCreated)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.UserRole_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	UserRole_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// userrole_Fetch read all UserRole's
func userrole_Fetch(tsql string) (int, []dm.UserRole, dm.UserRole, error) {

	var recItem dm.UserRole
	var recList []dm.UserRole

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.UserRole_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.UserRole_Id_sql, "")
		recItem.Name = get_String(rec, dm.UserRole_Name_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.UserRole_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.UserRole_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.UserRole_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.UserRole_SYSUpdatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.UserRole_SYSUpdated_sql, "")
		recItem.SYSCreated = get_String(rec, dm.UserRole_SYSCreated_sql, "")

		// If there are fields below, create the methods in adaptor\UserRole_impl.go

		//
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func UserRole_NewID(r dm.UserRole) string {

	id := uuid.New().String()

	return id
}

// userrole_Fetch read all UserRole's
func UserRole_New() (int, []dm.UserRole, dm.UserRole, error) {

	var r = dm.UserRole{}
	var rList []dm.UserRole

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
