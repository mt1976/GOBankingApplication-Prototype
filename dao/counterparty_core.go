package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterparty.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:49:58
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

// Counterparty_GetList() returns a list of all Counterparty records
func Counterparty_GetList() (int, []dm.Counterparty, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Counterparty_SQLTable)
	count, counterpartyList, _, _ := counterparty_Fetch(tsql)
	
	return count, counterpartyList, nil
}



// Counterparty_GetByID() returns a single Counterparty record
func Counterparty_GetByID(id string) (int, dm.Counterparty, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Counterparty_SQLTable)
	tsql = tsql + " WHERE " + dm.Counterparty_SQLSearchID + "='" + id + "'"
	_, _, counterpartyItem, _ := counterparty_Fetch(tsql)

	return 1, counterpartyItem, nil
}



// Counterparty_DeleteByID() deletes a single Counterparty record
func Counterparty_Delete(id string) {


	adaptor.Counterparty_Delete_impl(id)
	
}


// Counterparty_Store() saves/stores a Counterparty record to the database
func Counterparty_Store(r dm.Counterparty,req *http.Request) error {

	err := counterparty_Save(r,Audit_User(req))

	return err
}

// Counterparty_StoreSystem() saves/stores a Counterparty record to the database
func Counterparty_StoreSystem(r dm.Counterparty) error {
	
	err := counterparty_Save(r,Audit_Host())

	return err
}

// counterparty_Save() saves/stores a Counterparty record to the database
func counterparty_Save(r dm.Counterparty,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = Counterparty_NewID(r)
	}

// If there are fields below, create the methods in dao\counterparty_impl.go




















	
logs.Storing("Counterparty",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Counterparty_impl.go file
	err1 := adaptor.Counterparty_Delete_impl(r.CompID)
	err2 := adaptor.Counterparty_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterparty_Fetch read all Counterparty's
func counterparty_Fetch(tsql string) (int, []dm.Counterparty, dm.Counterparty, error) {

	var recItem dm.Counterparty
	var recList []dm.Counterparty

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.NameCentre  = get_String(rec, dm.Counterparty_NameCentre, "")
   recItem.NameFirm  = get_String(rec, dm.Counterparty_NameFirm, "")
   recItem.FullName  = get_String(rec, dm.Counterparty_FullName, "")
   recItem.TelephoneNumber  = get_String(rec, dm.Counterparty_TelephoneNumber, "")
   recItem.EmailAddress  = get_String(rec, dm.Counterparty_EmailAddress, "")
   recItem.CustomerType  = get_String(rec, dm.Counterparty_CustomerType, "")
   recItem.AccountOfficer  = get_String(rec, dm.Counterparty_AccountOfficer, "")
   recItem.CountryCode  = get_String(rec, dm.Counterparty_CountryCode, "")
   recItem.SectorCode  = get_String(rec, dm.Counterparty_SectorCode, "")
   recItem.CpartyGroupName  = get_String(rec, dm.Counterparty_CpartyGroupName, "")
   recItem.Notes  = get_String(rec, dm.Counterparty_Notes, "")
   recItem.Owner  = get_String(rec, dm.Counterparty_Owner, "")
   recItem.Authorised  = get_Bool(rec, dm.Counterparty_Authorised, "True")
   recItem.NameFirmName  = get_String(rec, dm.Counterparty_NameFirmName, "")
   recItem.NameCentreName  = get_String(rec, dm.Counterparty_NameCentreName, "")
   recItem.CountryCodeName  = get_String(rec, dm.Counterparty_CountryCodeName, "")
   recItem.SectorCodeName  = get_String(rec, dm.Counterparty_SectorCodeName, "")
   recItem.CompID  = get_String(rec, dm.Counterparty_CompID, "")
// If there are fields below, create the methods in adaptor\Counterparty_impl.go



















	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Counterparty_NewID(r dm.Counterparty) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

