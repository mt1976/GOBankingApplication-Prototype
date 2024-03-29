package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/credentialsaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"

"github.com/google/uuid"

	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// CredentialsAction_GetList() returns a list of all CredentialsAction records
func CredentialsAction_GetList() (int, []dm.CredentialsAction, error) {
	
	count, credentialsactionList, _ := adaptor.CredentialsAction_GetList_impl()
	
	return count, credentialsactionList, nil
}



// CredentialsAction_GetByID() returns a single CredentialsAction record
func CredentialsAction_GetByID(id string) (int, dm.CredentialsAction, error) {


	 _, credentialsactionItem, _ := adaptor.CredentialsAction_GetByID_impl(id)
	
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, credentialsactionItem, nil
}



// CredentialsAction_DeleteByID() deletes a single CredentialsAction record
func CredentialsAction_Delete(id string) {


	adaptor.CredentialsAction_Delete_impl(id)
	
}


// CredentialsAction_Store() saves/stores a CredentialsAction record to the database
func CredentialsAction_Store(r dm.CredentialsAction,req *http.Request) error {

	err, r := CredentialsAction_Validate(r)
	if err == nil {
		err = credentialsaction_Save(r, Audit_User(req))
	} else {
		logs.Information("CredentialsAction_Store()", err.Error())
	}

	return err
}

// CredentialsAction_StoreSystem() saves/stores a CredentialsAction record to the database
func CredentialsAction_StoreSystem(r dm.CredentialsAction) error {
	
	err, r := CredentialsAction_Validate(r)
	if err == nil {
		err = credentialsaction_Save(r, Audit_Host())
	} else {
		logs.Information("CredentialsAction_Store()", err.Error())
	}

	return err
}

// CredentialsAction_Validate() validates for saves/stores a CredentialsAction record to the database
func CredentialsAction_Validate(r dm.CredentialsAction) (error,dm.CredentialsAction) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// credentialsaction_Save() saves/stores a CredentialsAction record to the database
func credentialsaction_Save(r dm.CredentialsAction,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = CredentialsAction_NewID(r)
	}

// If there are fields below, create the methods in dao\credentialsaction_impl.go






	
logs.Storing("CredentialsAction",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CredentialsAction_impl.go file
	err1 := adaptor.CredentialsAction_Delete_impl(r.ID)
	err2 := adaptor.CredentialsAction_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// credentialsaction_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func CredentialsAction_NewID(r dm.CredentialsAction) string {
	
			id := uuid.New().String()
	
	return id
}



// credentialsaction_Fetch read all CredentialsAction's
func CredentialsAction_New() (int, []dm.CredentialsAction, dm.CredentialsAction, error) {

	var r = dm.CredentialsAction{}
	var rList []dm.CredentialsAction
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}