package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/customeronboarding.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CustomerOnboarding (customeronboarding)
// Endpoint 	        : CustomerOnboarding (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:40
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

// CustomerOnboarding_GetList() returns a list of all CustomerOnboarding records
func CustomerOnboarding_GetList() (int, []dm.CustomerOnboarding, error) {
	
	count, customeronboardingList, _ := adaptor.CustomerOnboarding_GetList_impl()
	
	return count, customeronboardingList, nil
}



// CustomerOnboarding_GetByID() returns a single CustomerOnboarding record
func CustomerOnboarding_GetByID(id string) (int, dm.CustomerOnboarding, error) {


	 _, customeronboardingItem, _ := adaptor.CustomerOnboarding_GetByID_impl(id)
	
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, customeronboardingItem, nil
}



// CustomerOnboarding_DeleteByID() deletes a single CustomerOnboarding record
func CustomerOnboarding_Delete(id string) {


	adaptor.CustomerOnboarding_Delete_impl(id)
	
}


// CustomerOnboarding_Store() saves/stores a CustomerOnboarding record to the database
func CustomerOnboarding_Store(r dm.CustomerOnboarding,req *http.Request) error {

	err, r := CustomerOnboarding_Validate(r)
	if err == nil {
		err = customeronboarding_Save(r, Audit_User(req))
	} else {
		logs.Information("CustomerOnboarding_Store()", err.Error())
	}

	return err
}

// CustomerOnboarding_StoreSystem() saves/stores a CustomerOnboarding record to the database
func CustomerOnboarding_StoreSystem(r dm.CustomerOnboarding) error {
	
	err, r := CustomerOnboarding_Validate(r)
	if err == nil {
		err = customeronboarding_Save(r, Audit_Host())
	} else {
		logs.Information("CustomerOnboarding_Store()", err.Error())
	}

	return err
}

// CustomerOnboarding_Validate() validates for saves/stores a CustomerOnboarding record to the database
func CustomerOnboarding_Validate(r dm.CustomerOnboarding) (error,dm.CustomerOnboarding) {
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

// customeronboarding_Save() saves/stores a CustomerOnboarding record to the database
func customeronboarding_Save(r dm.CustomerOnboarding,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = CustomerOnboarding_NewID(r)
	}

// If there are fields below, create the methods in dao\customeronboarding_impl.go















	
logs.Storing("CustomerOnboarding",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CustomerOnboarding_impl.go file
	err1 := adaptor.CustomerOnboarding_Delete_impl(r.ID)
	err2 := adaptor.CustomerOnboarding_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// customeronboarding_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func CustomerOnboarding_NewID(r dm.CustomerOnboarding) string {
	
			id := uuid.New().String()
	
	return id
}



// customeronboarding_Fetch read all CustomerOnboarding's
func CustomerOnboarding_New() (int, []dm.CustomerOnboarding, dm.CustomerOnboarding, error) {

	var r = dm.CustomerOnboarding{}
	var rList []dm.CustomerOnboarding
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}