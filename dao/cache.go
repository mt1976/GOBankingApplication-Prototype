package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/cache.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Cache (cache)
// Endpoint 	        : Cache (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:00
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	
)

// Cache_GetList() returns a list of all Cache records
func Cache_GetList() (int, []dm.Cache, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Cache_SQLTable)
	count, cacheList, _, _ := cache_Fetch(tsql)
	return count, cacheList, nil
}

// Cache_GetByID() returns a single Cache record
func Cache_GetByID(id string) (int, dm.Cache, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Cache_SQLTable)
	tsql = tsql + " WHERE " + dm.Cache_SQLSearchID + "='" + id + "'"

	_, _, cacheItem, _ := cache_Fetch(tsql)
	return 1, cacheItem, nil
}



// Cache_DeleteByID() deletes a single Cache record
func Cache_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Cache_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Cache_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Cache_Store() saves/stores a Cache record to the database
func Cache_Store(r dm.Cache) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Cache_NewID(r)
	}




//Deal with the if its Application or null add this bit, otherwise dont.
		//fmt.Println(credentialStore)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, Audit_User())
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",Audit_User())
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

	ts = addData(ts, dm.Cache_SYSId, r.SYSId)
	ts = addData(ts, dm.Cache_Id, r.Id)
	ts = addData(ts, dm.Cache_Object, r.Object)
	ts = addData(ts, dm.Cache_Field, r.Field)
	ts = addData(ts, dm.Cache_Value, r.Value)
	ts = addData(ts, dm.Cache_Expiry, r.Expiry)
	ts = addData(ts, dm.Cache_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.Cache_SYSWho, r.SYSWho)
	ts = addData(ts, dm.Cache_SYSHost, r.SYSHost)
	ts = addData(ts, dm.Cache_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.Cache_Source, r.Source)
	ts = addData(ts, dm.Cache_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.Cache_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.Cache_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.Cache_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Cache_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Cache_Delete(r.Id)
	das.Execute(tsql)


	return nil
}

// cache_Fetch read all employees
func cache_Fetch(tsql string) (int, []dm.Cache, dm.Cache, error) {

	var recItem dm.Cache
	var recList []dm.Cache

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Cache_Id,"")
   recItem.SYSId  = get_Int(rec, dm.Cache_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Cache_Id, "")
   recItem.Object  = get_String(rec, dm.Cache_Object, "")
   recItem.Field  = get_String(rec, dm.Cache_Field, "")
   recItem.Value  = get_String(rec, dm.Cache_Value, "")
   recItem.Expiry  = get_String(rec, dm.Cache_Expiry, "")
   recItem.SYSCreated  = get_String(rec, dm.Cache_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Cache_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Cache_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Cache_SYSUpdated, "")
   recItem.Source  = get_String(rec, dm.Cache_Source, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Cache_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Cache_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Cache_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Cache_SYSUpdatedHost, "")
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Cache_NewID(r dm.Cache) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------
