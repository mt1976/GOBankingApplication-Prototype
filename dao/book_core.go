package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/book.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Book (book)
// Endpoint 	        : Book (Book)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Book_GetList() returns a list of all Book records
func Book_GetList() (int, []dm.Book, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Book_SQLTable)
	count, bookList, _, _ := book_Fetch(tsql)
	
	return count, bookList, nil
}


// Book_GetLookup() returns a lookup list of all Book items in lookup format
func Book_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, bookList, _ := Book_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: bookList[i].BookName, Name: bookList[i].FullName})
	}
	return returnList
}


// Book_GetByID() returns a single Book record
func Book_GetByID(id string) (int, dm.Book, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Book_SQLTable)
	tsql = tsql + " WHERE " + dm.Book_SQLSearchID + "='" + id + "'"
	_, _, bookItem, _ := book_Fetch(tsql)

	return 1, bookItem, nil
}

// Book_GetByReverseLookup(id string) returns a single Book record using the FullName field as key.
func Book_GetByReverseLookup(id string) (int, dm.Book, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Book_SQLTable)
	tsql = tsql + " WHERE FullName = '" + id + "'"

	_, _, bookItem, _ := book_Fetch(tsql)
	
	return 1, bookItem, nil
}

// Book_DeleteByID() deletes a single Book record
func Book_Delete(id string) {


	adaptor.Book_Delete_impl(id)
	
}


// Book_Store() saves/stores a Book record to the database
func Book_Store(r dm.Book,req *http.Request) error {

	err := book_Save(r,Audit_User(req))

	return err
}

// Book_StoreSystem() saves/stores a Book record to the database
func Book_StoreSystem(r dm.Book) error {
	
	err := book_Save(r,Audit_Host())

	return err
}

// book_Save() saves/stores a Book record to the database
func book_Save(r dm.Book,usr string) error {

    var err error



	

	if len(r.BookName) == 0 {
		r.BookName = Book_NewID(r)
	}

// If there are fields below, create the methods in dao\book_impl.go











	
logs.Storing("Book",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Book_impl.go file
	err1 := adaptor.Book_Delete_impl(r.BookName)
	err2 := adaptor.Book_Update_impl(r.BookName,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// book_Fetch read all Book's
func book_Fetch(tsql string) (int, []dm.Book, dm.Book, error) {

	var recItem dm.Book
	var recList []dm.Book

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.BookName  = get_String(rec, dm.Book_BookName, "")
   recItem.FullName  = get_String(rec, dm.Book_FullName, "")
   recItem.PLManage  = get_String(rec, dm.Book_PLManage, "")
   recItem.PLTransfer  = get_String(rec, dm.Book_PLTransfer, "")
   recItem.DerivePL  = get_Bool(rec, dm.Book_DerivePL, "True")
   recItem.CostOfCarry  = get_Bool(rec, dm.Book_CostOfCarry, "True")
   recItem.CostOfFunding  = get_Bool(rec, dm.Book_CostOfFunding, "True")
   recItem.LotAllocationMethod  = get_String(rec, dm.Book_LotAllocationMethod, "")
   recItem.InternalId  = get_Int(rec, dm.Book_InternalId, "0")
// If there are fields below, create the methods in adaptor\Book_impl.go










	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Book_NewID(r dm.Book) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

