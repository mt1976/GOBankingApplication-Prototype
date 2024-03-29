package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:52
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// ExternalMessage_GetList() returns a list of all ExternalMessage records
func ExternalMessage_GetList() (int, []dm.ExternalMessage, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.ExternalMessage_SQLTable)
	count, externalmessageList, _, _ := externalmessage_Fetch(tsql)

	return count, externalmessageList, nil
}

// ExternalMessage_GetByID() returns a single ExternalMessage record
func ExternalMessage_GetByID(id string) (int, dm.ExternalMessage, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.ExternalMessage_SQLTable)
	tsql = tsql + " WHERE " + dm.ExternalMessage_SQLSearchID + "='" + id + "'"
	_, _, externalmessageItem, _ := externalmessage_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, externalmessageItem, nil
}

// ExternalMessage_DeleteByID() deletes a single ExternalMessage record
func ExternalMessage_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.ExternalMessage_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.ExternalMessage_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// ExternalMessage_Store() saves/stores a ExternalMessage record to the database
func ExternalMessage_Store(r dm.ExternalMessage, req *http.Request) error {

	err, r := ExternalMessage_Validate(r)
	if err == nil {
		err = externalmessage_Save(r, Audit_User(req))
	} else {
		logs.Information("ExternalMessage_Store()", err.Error())
	}

	return err
}

// ExternalMessage_StoreSystem() saves/stores a ExternalMessage record to the database
func ExternalMessage_StoreSystem(r dm.ExternalMessage) error {

	err, r := ExternalMessage_Validate(r)
	if err == nil {
		err = externalmessage_Save(r, Audit_Host())
	} else {
		logs.Information("ExternalMessage_Store()", err.Error())
	}

	return err
}

// ExternalMessage_Validate() validates for saves/stores a ExternalMessage record to the database
func ExternalMessage_Validate(r dm.ExternalMessage) (error, dm.ExternalMessage) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// externalmessage_Save() saves/stores a ExternalMessage record to the database
func externalmessage_Save(r dm.ExternalMessage, usr string) error {

	var err error

	if len(r.MessageID) == 0 {
		r.MessageID = ExternalMessage_NewID(r)
	}

	// If there are fields below, create the methods in dao\externalmessage_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("ExternalMessage", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.ExternalMessage_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.ExternalMessage_MessageID_sql, r.MessageID)
	ts = addData(ts, dm.ExternalMessage_MessageFormat_sql, r.MessageFormat)
	ts = addData(ts, dm.ExternalMessage_MessageDeliveredTo_sql, r.MessageDeliveredTo)
	ts = addData(ts, dm.ExternalMessage_MessageBody_sql, r.MessageBody)
	ts = addData(ts, dm.ExternalMessage_MessageFilename_sql, r.MessageFilename)
	ts = addData(ts, dm.ExternalMessage_MessageLife_sql, r.MessageLife)
	ts = addData(ts, dm.ExternalMessage_MessageDate_sql, r.MessageDate)
	ts = addData(ts, dm.ExternalMessage_MessageTime_sql, r.MessageTime)
	ts = addData(ts, dm.ExternalMessage_MessageTimeoutAction_sql, r.MessageTimeoutAction)
	ts = addData(ts, dm.ExternalMessage_MessageACKNAK_sql, r.MessageACKNAK)
	ts = addData(ts, dm.ExternalMessage_ResponseID_sql, r.ResponseID)
	ts = addData(ts, dm.ExternalMessage_ResponseFilename_sql, r.ResponseFilename)
	ts = addData(ts, dm.ExternalMessage_ResponseBody_sql, r.ResponseBody)
	ts = addData(ts, dm.ExternalMessage_ResponseDate_sql, r.ResponseDate)
	ts = addData(ts, dm.ExternalMessage_ResponseTime_sql, r.ResponseTime)
	ts = addData(ts, dm.ExternalMessage_ResponseAdditionalInfo_sql, r.ResponseAdditionalInfo)
	ts = addData(ts, dm.ExternalMessage_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.ExternalMessage_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.ExternalMessage_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.ExternalMessage_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.ExternalMessage_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.ExternalMessage_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.ExternalMessage_MessageTimeout_sql, r.MessageTimeout)
	ts = addData(ts, dm.ExternalMessage_MessageEmitted_sql, r.MessageEmitted)
	ts = addData(ts, dm.ExternalMessage_ResponseRecieved_sql, r.ResponseRecieved)
	ts = addData(ts, dm.ExternalMessage_MessageClass_sql, r.MessageClass)
	ts = addData(ts, dm.ExternalMessage_AppID_sql, r.AppID)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.ExternalMessage_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	ExternalMessage_Delete(r.MessageID)
	das.Execute(tsql)

	return err

}

// externalmessage_Fetch read all ExternalMessage's
func externalmessage_Fetch(tsql string) (int, []dm.ExternalMessage, dm.ExternalMessage, error) {

	var recItem dm.ExternalMessage
	var recList []dm.ExternalMessage

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.ExternalMessage_SYSId_sql, "0")
		recItem.MessageID = get_String(rec, dm.ExternalMessage_MessageID_sql, "")
		recItem.MessageFormat = get_String(rec, dm.ExternalMessage_MessageFormat_sql, "")
		recItem.MessageDeliveredTo = get_String(rec, dm.ExternalMessage_MessageDeliveredTo_sql, "")
		recItem.MessageBody = get_String(rec, dm.ExternalMessage_MessageBody_sql, "")
		recItem.MessageFilename = get_String(rec, dm.ExternalMessage_MessageFilename_sql, "")
		recItem.MessageLife = get_String(rec, dm.ExternalMessage_MessageLife_sql, "")
		recItem.MessageDate = get_String(rec, dm.ExternalMessage_MessageDate_sql, "")
		recItem.MessageTime = get_String(rec, dm.ExternalMessage_MessageTime_sql, "")
		recItem.MessageTimeoutAction = get_String(rec, dm.ExternalMessage_MessageTimeoutAction_sql, "")
		recItem.MessageACKNAK = get_String(rec, dm.ExternalMessage_MessageACKNAK_sql, "")
		recItem.ResponseID = get_String(rec, dm.ExternalMessage_ResponseID_sql, "")
		recItem.ResponseFilename = get_String(rec, dm.ExternalMessage_ResponseFilename_sql, "")
		recItem.ResponseBody = get_String(rec, dm.ExternalMessage_ResponseBody_sql, "")
		recItem.ResponseDate = get_String(rec, dm.ExternalMessage_ResponseDate_sql, "")
		recItem.ResponseTime = get_String(rec, dm.ExternalMessage_ResponseTime_sql, "")
		recItem.ResponseAdditionalInfo = get_String(rec, dm.ExternalMessage_ResponseAdditionalInfo_sql, "")
		recItem.SYSCreated = get_String(rec, dm.ExternalMessage_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.ExternalMessage_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.ExternalMessage_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.ExternalMessage_SYSUpdated_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.ExternalMessage_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.ExternalMessage_SYSUpdatedHost_sql, "")
		recItem.MessageTimeout = get_String(rec, dm.ExternalMessage_MessageTimeout_sql, "")
		recItem.MessageEmitted = get_String(rec, dm.ExternalMessage_MessageEmitted_sql, "")
		recItem.ResponseRecieved = get_String(rec, dm.ExternalMessage_ResponseRecieved_sql, "")
		recItem.MessageClass = get_String(rec, dm.ExternalMessage_MessageClass_sql, "")
		recItem.AppID = get_String(rec, dm.ExternalMessage_AppID_sql, "")

		// If there are fields below, create the methods in adaptor\ExternalMessage_impl.go

		//
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func ExternalMessage_NewID(r dm.ExternalMessage) string {

	id := uuid.New().String()

	return id
}

// externalmessage_Fetch read all ExternalMessage's
func ExternalMessage_New() (int, []dm.ExternalMessage, dm.ExternalMessage, error) {

	var r = dm.ExternalMessage{}
	var rList []dm.ExternalMessage

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
