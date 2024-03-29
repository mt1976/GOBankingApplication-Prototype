package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/inbox.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:53
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

// Inbox_GetList() returns a list of all Inbox records
func Inbox_GetList() (int, []dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	count, inboxList, _, _ := inbox_Fetch(tsql)

	return count, inboxList, nil
}

// Inbox_GetByID() returns a single Inbox record
func Inbox_GetByID(id string) (int, dm.Inbox, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " WHERE " + dm.Inbox_SQLSearchID + "='" + id + "'"
	_, _, inboxItem, _ := inbox_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, inboxItem, nil
}

// Inbox_DeleteByID() deletes a single Inbox record
func Inbox_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.Inbox_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Inbox_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Inbox_Store() saves/stores a Inbox record to the database
func Inbox_Store(r dm.Inbox, req *http.Request) error {

	err, r := Inbox_Validate(r)
	if err == nil {
		err = inbox_Save(r, Audit_User(req))
	} else {
		logs.Information("Inbox_Store()", err.Error())
	}

	return err
}

// Inbox_StoreSystem() saves/stores a Inbox record to the database
func Inbox_StoreSystem(r dm.Inbox) error {

	err, r := Inbox_Validate(r)
	if err == nil {
		err = inbox_Save(r, Audit_Host())
	} else {
		logs.Information("Inbox_Store()", err.Error())
	}

	return err
}

// Inbox_Validate() validates for saves/stores a Inbox record to the database
func Inbox_Validate(r dm.Inbox) (error, dm.Inbox) {
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

// inbox_Save() saves/stores a Inbox record to the database
func inbox_Save(r dm.Inbox, usr string) error {

	var err error

	if len(r.MailId) == 0 {
		r.MailId = Inbox_NewID(r)
	}

	// If there are fields below, create the methods in dao\inbox_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Inbox", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Inbox_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Inbox_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Inbox_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Inbox_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Inbox_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Inbox_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Inbox_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Inbox_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Inbox_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Inbox_MailId_sql, r.MailId)
	ts = addData(ts, dm.Inbox_MailTo_sql, r.MailTo)
	ts = addData(ts, dm.Inbox_MailFrom_sql, r.MailFrom)
	ts = addData(ts, dm.Inbox_MailSource_sql, r.MailSource)
	ts = addData(ts, dm.Inbox_MailSent_sql, r.MailSent)
	ts = addData(ts, dm.Inbox_MailUnread_sql, r.MailUnread)
	ts = addData(ts, dm.Inbox_MailSubject_sql, r.MailSubject)
	ts = addData(ts, dm.Inbox_MailContent_sql, r.MailContent)
	ts = addData(ts, dm.Inbox_MailAcknowledged_sql, r.MailAcknowledged)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.Inbox_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Inbox_Delete(r.MailId)
	das.Execute(tsql)

	return err

}

// inbox_Fetch read all Inbox's
func inbox_Fetch(tsql string) (int, []dm.Inbox, dm.Inbox, error) {

	var recItem dm.Inbox
	var recList []dm.Inbox

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Inbox_SYSId_sql, "0")
		recItem.SYSCreated = get_String(rec, dm.Inbox_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.Inbox_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.Inbox_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Inbox_SYSUpdated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Inbox_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Inbox_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Inbox_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Inbox_SYSUpdatedHost_sql, "")
		recItem.MailId = get_String(rec, dm.Inbox_MailId_sql, "")
		recItem.MailTo = get_String(rec, dm.Inbox_MailTo_sql, "")
		recItem.MailFrom = get_String(rec, dm.Inbox_MailFrom_sql, "")
		recItem.MailSource = get_String(rec, dm.Inbox_MailSource_sql, "")
		recItem.MailSent = get_String(rec, dm.Inbox_MailSent_sql, "")
		recItem.MailUnread = get_String(rec, dm.Inbox_MailUnread_sql, "")
		recItem.MailSubject = get_String(rec, dm.Inbox_MailSubject_sql, "")
		recItem.MailContent = get_String(rec, dm.Inbox_MailContent_sql, "")
		recItem.MailAcknowledged = get_String(rec, dm.Inbox_MailAcknowledged_sql, "")

		// If there are fields below, create the methods in adaptor\Inbox_impl.go

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

func Inbox_NewID(r dm.Inbox) string {

	id := uuid.New().String()

	return id
}

// inbox_Fetch read all Inbox's
func Inbox_New() (int, []dm.Inbox, dm.Inbox, error) {

	var r = dm.Inbox{}
	var rList []dm.Inbox

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
