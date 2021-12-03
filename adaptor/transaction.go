package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/transaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Transaction (transaction)
// Endpoint 	        : Transaction (Ref)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:58
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Transaction_Delete(id string) error {
	var er error

	message:= "Implement Transaction_Delete: " + id

	// Implement Transaction_Delete_Impl in transaction_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Transaction_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Transaction_Update(item dm.Transaction) error {
	var er error

	message:= "Implement Transaction_Update: " + item.SienaReference

	// Implement Transaction_Update_Impl in transaction_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Transaction_Update_Impl(item)
	//

	logs.Success(message)
	return er
}