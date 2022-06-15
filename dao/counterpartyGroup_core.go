package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartygroup.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:56
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

// CounterpartyGroup_GetList() returns a list of all CounterpartyGroup records
func CounterpartyGroup_GetList() (int, []dm.CounterpartyGroup, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyGroup_SQLTable)
	count, counterpartygroupList, _, _ := counterpartygroup_Fetch(tsql)
	
	return count, counterpartygroupList, nil
}


// CounterpartyGroup_GetLookup() returns a lookup list of all CounterpartyGroup items in lookup format
func CounterpartyGroup_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, counterpartygroupList, _ := CounterpartyGroup_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: counterpartygroupList[i].Name, Name: counterpartygroupList[i].Name})
	}
	return returnList
}


// CounterpartyGroup_GetByID() returns a single CounterpartyGroup record
func CounterpartyGroup_GetByID(id string) (int, dm.CounterpartyGroup, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyGroup_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyGroup_SQLSearchID + "='" + id + "'"
	_, _, counterpartygroupItem, _ := counterpartygroup_Fetch(tsql)

	return 1, counterpartygroupItem, nil
}



// CounterpartyGroup_DeleteByID() deletes a single CounterpartyGroup record
func CounterpartyGroup_Delete(id string) {


	adaptor.CounterpartyGroup_Delete_impl(id)
	
}


// CounterpartyGroup_Store() saves/stores a CounterpartyGroup record to the database
func CounterpartyGroup_Store(r dm.CounterpartyGroup,req *http.Request) error {

	err := counterpartygroup_Save(r,Audit_User(req))

	return err
}

// CounterpartyGroup_StoreSystem() saves/stores a CounterpartyGroup record to the database
func CounterpartyGroup_StoreSystem(r dm.CounterpartyGroup) error {
	
	err := counterpartygroup_Save(r,Audit_Host())

	return err
}

// counterpartygroup_Save() saves/stores a CounterpartyGroup record to the database
func counterpartygroup_Save(r dm.CounterpartyGroup,usr string) error {

    var err error



	

	if len(r.Name) == 0 {
		r.Name = CounterpartyGroup_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartygroup_impl.go





	
logs.Storing("CounterpartyGroup",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyGroup_impl.go file
	err1 := adaptor.CounterpartyGroup_Delete_impl(r.Name)
	err2 := adaptor.CounterpartyGroup_Update_impl(r.Name,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartygroup_Fetch read all CounterpartyGroup's
func counterpartygroup_Fetch(tsql string) (int, []dm.CounterpartyGroup, dm.CounterpartyGroup, error) {

	var recItem dm.CounterpartyGroup
	var recList []dm.CounterpartyGroup

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.Name  = get_String(rec, dm.CounterpartyGroup_Name, "")
   recItem.CountryCode  = get_String(rec, dm.CounterpartyGroup_CountryCode, "")
   recItem.SuperGroup  = get_String(rec, dm.CounterpartyGroup_SuperGroup, "")
// If there are fields below, create the methods in adaptor\CounterpartyGroup_impl.go




	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func CounterpartyGroup_NewID(r dm.CounterpartyGroup) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

