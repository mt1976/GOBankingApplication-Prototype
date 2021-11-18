package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Country_Delete(id string) error {
	var er error

	message:= "Implement Country_Delete: " + id

	// Implement Country_Delete_Impl in country_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Country_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Country_Update(item dm.Country) error {
	var er error

	message:= "Implement Country_Update: " + item.Code

	// Implement Country_Update_Impl in country_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Country_Update_Impl(item)
	//

	logs.Success(message)
	return er
}