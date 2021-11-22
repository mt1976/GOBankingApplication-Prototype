package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/firm.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:03
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Firm_Delete(id string) error {
	var er error

	message:= "Implement Firm_Delete: " + id

	// Implement Firm_Delete_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Firm_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Firm_Update(item dm.Firm) error {
	var er error

	message:= "Implement Firm_Update: " + item.FirmName

	// Implement Firm_Update_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Firm_Update_Impl(item)
	//

	logs.Success(message)
	return er
}