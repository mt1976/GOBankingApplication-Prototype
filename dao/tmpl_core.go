package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/tmpl.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Tmpl_GetList() returns a list of all Tmpl records
func Tmpl_GetList() (int, []dm.Tmpl, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	count, tmplList, _, _ := tmpl_Fetch(tsql)
	
	return count, tmplList, nil
}



// Tmpl_GetByID() returns a single Tmpl record
func Tmpl_GetByID(id string) (int, dm.Tmpl, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	tsql = tsql + " WHERE " + dm.Tmpl_SQLSearchID + "='" + id + "'"
	_, _, tmplItem, _ := tmpl_Fetch(tsql)

	return 1, tmplItem, nil
}



// Tmpl_DeleteByID() deletes a single Tmpl record
func Tmpl_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Tmpl_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Tmpl_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Tmpl_Store() saves/stores a Tmpl record to the database
func Tmpl_Store(r dm.Tmpl,req *http.Request) error {

	err := tmpl_Save(r,Audit_User(req))

	return err
}

// Tmpl_StoreSystem() saves/stores a Tmpl record to the database
func Tmpl_StoreSystem(r dm.Tmpl) error {
	
	err := tmpl_Save(r,Audit_Host())

	return err
}

// tmpl_Save() saves/stores a Tmpl record to the database
func tmpl_Save(r dm.Tmpl,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = Tmpl_NewID(r)
	}

// If there are fields below, create the methods in dao\tmpl_impl.go










  r.ExtraField,err = adaptor.Tmpl_ExtraField_OnStore_impl (r.ExtraField,r,usr)

  r.ExtraField3,err = adaptor.Tmpl_ExtraField3_OnStore_impl (r.ExtraField3,r,usr)


	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("Tmpl",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Tmpl_SYSId, r.SYSId)
	ts = addData(ts, dm.Tmpl_FIELD1, r.FIELD1)
	ts = addData(ts, dm.Tmpl_FIELD2, r.FIELD2)
	ts = addData(ts, dm.Tmpl_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.Tmpl_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.Tmpl_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.Tmpl_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.Tmpl_SYSUpdatedHost, r.SYSUpdatedHost)
	ts = addData(ts, dm.Tmpl_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.Tmpl_ID, r.ID)
	
	
	
		
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Tmpl_Delete(r.ID)
	das.Execute(tsql)



	return err

}



// tmpl_Fetch read all Tmpl's
func tmpl_Fetch(tsql string) (int, []dm.Tmpl, dm.Tmpl, error) {

	var recItem dm.Tmpl
	var recList []dm.Tmpl

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Tmpl_SYSId, "0")
   recItem.FIELD1  = get_String(rec, dm.Tmpl_FIELD1, "N")
   recItem.FIELD2  = get_String(rec, dm.Tmpl_FIELD2, "")
   recItem.SYSCreated  = get_String(rec, dm.Tmpl_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Tmpl_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Tmpl_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Tmpl_SYSUpdated, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Tmpl_SYSUpdatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Tmpl_SYSUpdatedBy, "")
   recItem.ID  = get_String(rec, dm.Tmpl_ID, "")



// If there are fields below, create the methods in adaptor\Tmpl_impl.go










   recItem.ExtraField  = adaptor.Tmpl_ExtraField_OnFetch_impl (recItem)

   recItem.ExtraField3  = adaptor.Tmpl_ExtraField3_OnFetch_impl (recItem)

	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Tmpl_NewID(r dm.Tmpl) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

