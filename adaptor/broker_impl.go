package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/broker.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:13
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Broker_Delete_Impl(id string, usr string) error {
	var er error

	message := "Implement Broker_Delete: " + id

	// Implement Broker_Delete_Impl in broker_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Broker_Update_Impl(item dm.Broker, usr string) error {
	var er error

	message := "Implement Broker_Update: " + item.Code

	// Implement Broker_Update_Impl in broker_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Update_Impl(item)
	//

	logs.Success(message)
	return er
}