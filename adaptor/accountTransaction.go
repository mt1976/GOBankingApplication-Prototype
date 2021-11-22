package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/accounttransaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : AccountTransaction (accounttransaction)
// Endpoint 	        : AccountTransaction (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 19:43:22
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func AccountTransaction_Delete(id string) error {
	var er error

	message:= "Implement AccountTransaction_Delete: " + id

	// Implement AccountTransaction_Delete_Impl in accounttransaction_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := AccountTransaction_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func AccountTransaction_Update(item dm.AccountTransaction) error {
	var er error

	message:= "Implement AccountTransaction_Update: " + item.SienaReference

	// Implement AccountTransaction_Update_Impl in accounttransaction_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := AccountTransaction_Update_Impl(item)
	//

	logs.Success(message)
	return er
}