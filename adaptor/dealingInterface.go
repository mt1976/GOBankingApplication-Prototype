package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/dealinginterface.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : DealingInterface (dealinginterface)
// Endpoint 	        : DealingInterface (Name)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 01/12/2021 at 20:36:39
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func DealingInterface_Delete(id string) error {
	var er error

	message:= "Implement DealingInterface_Delete: " + id

	// Implement DealingInterface_Delete_Impl in dealinginterface_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealingInterface_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func DealingInterface_Update(item dm.DealingInterface) error {
	var er error

	message:= "Implement DealingInterface_Update: " + item.Name

	// Implement DealingInterface_Update_Impl in dealinginterface_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealingInterface_Update_Impl(item)
	//

	logs.Success(message)
	return er
}