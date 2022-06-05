package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/simulatorswift.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : SimulatorSWIFT (simulatorswift)
// Endpoint 	        : SimulatorSWIFT (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 05/06/2022 at 12:56:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"

	
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// SimulatorSWIFT_GetList() returns a list of all SimulatorSWIFT records
func SimulatorSWIFT_GetList() (int, []dm.SimulatorSWIFT, error) {
	
	count, simulatorswiftList, _ := adaptor.SimulatorSWIFT_GetList_impl()
	
	return count, simulatorswiftList, nil
}



// SimulatorSWIFT_GetByID() returns a single SimulatorSWIFT record
func SimulatorSWIFT_GetByID(id string) (int, dm.SimulatorSWIFT, error) {


	 _, simulatorswiftItem, _ := adaptor.SimulatorSWIFT_GetByID_impl(id)
	
	return 1, simulatorswiftItem, nil
}



// SimulatorSWIFT_DeleteByID() deletes a single SimulatorSWIFT record
func SimulatorSWIFT_Delete(id string) {


	adaptor.SimulatorSWIFT_Delete_impl(id)
	
}


// SimulatorSWIFT_Store() saves/stores a SimulatorSWIFT record to the database
func SimulatorSWIFT_Store(r dm.SimulatorSWIFT,req *http.Request) error {

	err := simulatorswift_Save(r,Audit_User(req))

	return err
}

// SimulatorSWIFT_StoreSystem() saves/stores a SimulatorSWIFT record to the database
func SimulatorSWIFT_StoreSystem(r dm.SimulatorSWIFT) error {
	
	err := simulatorswift_Save(r,Audit_Host())

	return err
}

// simulatorswift_Save() saves/stores a SimulatorSWIFT record to the database
func simulatorswift_Save(r dm.SimulatorSWIFT,usr string) error {

    var err error



	logs.Storing("SimulatorSWIFT",fmt.Sprintf("%s", r))

	if len(r.ID) == 0 {
		r.ID = SimulatorSWIFT_NewID(r)
	}

// If there are fields below, create the methods in dao\simulatorswift_impl.go











  err = adaptor.SimulatorSWIFT_MessageFmt_Override_Store_impl (r)


	


// Please Create Functions Below in the adaptor/SimulatorSWIFT_impl.go file
	err1 := adaptor.SimulatorSWIFT_Delete_impl(r.ID)
	err2 := adaptor.SimulatorSWIFT_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// simulatorswift_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	

func SimulatorSWIFT_NewID(r dm.SimulatorSWIFT) string {
	
		// SimulatorSWIFT_NewID_impl should be specified in dao/SimulatorSWIFT_impl.go
		// to provide the implementation for the special case.
		// override should return id - override function should be defined as
		// SimulatorSWIFT_NewID_impl(r dm.SimulatorSWIFT) string {...}
		//
		id := adaptor.SimulatorSWIFT_NewID_impl(r)
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

