package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/firm.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:52:05
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

// Firm_GetList() returns a list of all Firm records
func Firm_GetList() (int, []dm.Firm, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	count, firmList, _, _ := firm_Fetch(tsql)
	
	return count, firmList, nil
}


// Firm_GetLookup() returns a lookup list of all Firm items in lookup format
func Firm_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, firmList, _ := Firm_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: firmList[i].FirmName, Name: firmList[i].FullName})
	}
	return returnList
}


// Firm_GetByID() returns a single Firm record
func Firm_GetByID(id string) (int, dm.Firm, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	tsql = tsql + " WHERE " + dm.Firm_SQLSearchID + "='" + id + "'"
	_, _, firmItem, _ := firm_Fetch(tsql)

	return 1, firmItem, nil
}

// Firm_GetByReverseLookup(id string) returns a single Firm record using the FullName field as key.
func Firm_GetByReverseLookup(id string) (int, dm.Firm, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	tsql = tsql + " WHERE FullName = '" + id + "'"

	_, _, firmItem, _ := firm_Fetch(tsql)
	
	return 1, firmItem, nil
}

// Firm_DeleteByID() deletes a single Firm record
func Firm_Delete(id string) {


	adaptor.Firm_Delete_impl(id)
	
}


// Firm_Store() saves/stores a Firm record to the database
func Firm_Store(r dm.Firm,req *http.Request) error {

	err := firm_Save(r,Audit_User(req))

	return err
}

// Firm_StoreSystem() saves/stores a Firm record to the database
func Firm_StoreSystem(r dm.Firm) error {
	
	err := firm_Save(r,Audit_Host())

	return err
}

// firm_Save() saves/stores a Firm record to the database
func firm_Save(r dm.Firm,usr string) error {

    var err error



	

	if len(r.FirmName) == 0 {
		r.FirmName = Firm_NewID(r)
	}

// If there are fields below, create the methods in dao\firm_impl.go






	
logs.Storing("Firm",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Firm_impl.go file
	err1 := adaptor.Firm_Delete_impl(r.FirmName)
	err2 := adaptor.Firm_Update_impl(r.FirmName,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// firm_Fetch read all Firm's
func firm_Fetch(tsql string) (int, []dm.Firm, dm.Firm, error) {

	var recItem dm.Firm
	var recList []dm.Firm

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.FirmName  = get_String(rec, dm.Firm_FirmName, "")
   recItem.FullName  = get_String(rec, dm.Firm_FullName, "")
   recItem.Country  = get_String(rec, dm.Firm_Country, "")
   recItem.Sector  = get_String(rec, dm.Firm_Sector, "")
// If there are fields below, create the methods in adaptor\Firm_impl.go





	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Firm_NewID(r dm.Firm) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

