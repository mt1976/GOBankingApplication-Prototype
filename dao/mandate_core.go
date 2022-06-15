package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/mandate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:58:56
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

// Mandate_GetList() returns a list of all Mandate records
func Mandate_GetList() (int, []dm.Mandate, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	count, mandateList, _, _ := mandate_Fetch(tsql)
	
	return count, mandateList, nil
}



// Mandate_GetByID() returns a single Mandate record
func Mandate_GetByID(id string) (int, dm.Mandate, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	tsql = tsql + " WHERE " + dm.Mandate_SQLSearchID + "='" + id + "'"
	_, _, mandateItem, _ := mandate_Fetch(tsql)

	return 1, mandateItem, nil
}



// Mandate_DeleteByID() deletes a single Mandate record
func Mandate_Delete(id string) {


	adaptor.Mandate_Delete_impl(id)
	
}


// Mandate_Store() saves/stores a Mandate record to the database
func Mandate_Store(r dm.Mandate,req *http.Request) error {

	err := mandate_Save(r,Audit_User(req))

	return err
}

// Mandate_StoreSystem() saves/stores a Mandate record to the database
func Mandate_StoreSystem(r dm.Mandate) error {
	
	err := mandate_Save(r,Audit_Host())

	return err
}

// mandate_Save() saves/stores a Mandate record to the database
func mandate_Save(r dm.Mandate,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = Mandate_NewID(r)
	}

// If there are fields below, create the methods in dao\mandate_impl.go





















	
logs.Storing("Mandate",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Mandate_impl.go file
	err1 := adaptor.Mandate_Delete_impl(r.CompID)
	err2 := adaptor.Mandate_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// mandate_Fetch read all Mandate's
func mandate_Fetch(tsql string) (int, []dm.Mandate, dm.Mandate, error) {

	var recItem dm.Mandate
	var recList []dm.Mandate

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.MandatedUserKeyCounterpartyFirm  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyFirm, "")
   recItem.MandatedUserKeyCounterpartyCentre  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyCentre, "")
   recItem.MandatedUserKeyUserName  = get_String(rec, dm.Mandate_MandatedUserKeyUserName, "")
   recItem.TelephoneNumber  = get_String(rec, dm.Mandate_TelephoneNumber, "")
   recItem.EmailAddress  = get_String(rec, dm.Mandate_EmailAddress, "")
   recItem.Active  = get_Bool(rec, dm.Mandate_Active, "True")
   recItem.FirstName  = get_String(rec, dm.Mandate_FirstName, "")
   recItem.Surname  = get_String(rec, dm.Mandate_Surname, "")
   recItem.DateOfBirth  = get_Time(rec, dm.Mandate_DateOfBirth, "")
   recItem.Postcode  = get_String(rec, dm.Mandate_Postcode, "")
   recItem.NationalIDNo  = get_String(rec, dm.Mandate_NationalIDNo, "")
   recItem.PassportNo  = get_String(rec, dm.Mandate_PassportNo, "")
   recItem.Country  = get_String(rec, dm.Mandate_Country, "")
   recItem.CountryName  = get_String(rec, dm.Mandate_CountryName, "")
   recItem.FirmName  = get_String(rec, dm.Mandate_FirmName, "")
   recItem.CentreName  = get_String(rec, dm.Mandate_CentreName, "")
   recItem.Notify  = get_Bool(rec, dm.Mandate_Notify, "True")
   recItem.SystemUser  = get_String(rec, dm.Mandate_SystemUser, "")
   recItem.CompID  = get_String(rec, dm.Mandate_CompID, "")
// If there are fields below, create the methods in adaptor\Mandate_impl.go




















	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Mandate_NewID(r dm.Mandate) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

