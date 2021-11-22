package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/account.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 18:11:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Account_Delete(id string) error {
	var er error

	message:= "Implement Account_Delete: " + id

	// Implement Account_Delete_Impl in account_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Account_Update(item dm.Account) error {
	var er error

	message:= "Implement Account_Update: " + item.SienaReference

	// Implement Account_Update_Impl in account_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Update_Impl(item)
	//

	logs.Success(message)
	return er
}