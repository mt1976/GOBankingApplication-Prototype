package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/account.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:38
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
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
)

// Account_GetList() returns a list of all Account records
func Account_GetList() (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	count, accountList, _, _ := account_Fetch(tsql)
	return count, accountList, nil
}

// Account_GetByID() returns a single Account record
func Account_GetByID(id string) (int, dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_SQLSearchID + "='" + id + "'"

	_, _, accountItem, _ := account_Fetch(tsql)
	return 1, accountItem, nil
}

// Account_GetByReverseLookup(id string) returns a single Account record
func Account_GetByReverseLookup(id string) (int, dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE CashBalance = '" + id + "'"

	_, _, accountItem, _ := account_Fetch(tsql)
	return 1, accountItem, nil
}

// Account_DeleteByID() deletes a single Account record
func Account_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Account_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Account_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Account_Store() saves/stores a Account record to the database
func Account_Store(r dm.Account) error {

	logs.Storing("Account",fmt.Sprintf("%s", r))

	if len(r.SienaReference) == 0 {
		r.SienaReference = Account_NewID(r)
	}



	adaptor.Account_Delete(r.SienaReference)
	adaptor.Account_Update(r)


	return nil
}

// account_Fetch read all employees
func account_Fetch(tsql string) (int, []dm.Account, dm.Account, error) {

	var recItem dm.Account
	var recList []dm.Account

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Account_SienaReference,"")
   recItem.SienaReference  = get_String(rec, dm.Account_SienaReference, "")
   recItem.CustomerSienaView  = get_String(rec, dm.Account_CustomerSienaView, "")
   recItem.SienaCommonRef  = get_String(rec, dm.Account_SienaCommonRef, "")
   recItem.Status  = get_String(rec, dm.Account_Status, "")
   recItem.StartDate  = get_Time(rec, dm.Account_StartDate, "")
   recItem.MaturityDate  = get_Time(rec, dm.Account_MaturityDate, "")
   recItem.ContractNumber  = get_String(rec, dm.Account_ContractNumber, "")
   recItem.ExternalReference  = get_String(rec, dm.Account_ExternalReference, "")
   recItem.CCY  = get_String(rec, dm.Account_CCY, "")
   recItem.Book  = get_String(rec, dm.Account_Book, "")
   recItem.MandatedUser  = get_String(rec, dm.Account_MandatedUser, "")
   recItem.BackOfficeNotes  = get_String(rec, dm.Account_BackOfficeNotes, "")
   recItem.CashBalance  = get_Float(rec, dm.Account_CashBalance, "0.00")
   recItem.AccountNumber  = get_String(rec, dm.Account_AccountNumber, "")
   recItem.AccountName  = get_String(rec, dm.Account_AccountName, "")
   recItem.LedgerBalance  = get_Float(rec, dm.Account_LedgerBalance, "0.00")
   recItem.Portfolio  = get_String(rec, dm.Account_Portfolio, "")
   recItem.AgreementId  = get_Int(rec, dm.Account_AgreementId, "0")
   recItem.BackOfficeRefNo  = get_String(rec, dm.Account_BackOfficeRefNo, "")
   recItem.ISIN  = get_String(rec, dm.Account_ISIN, "")
   recItem.UTI  = get_String(rec, dm.Account_UTI, "")
   recItem.CCYName  = get_String(rec, dm.Account_CCYName, "")
   recItem.BookName  = get_String(rec, dm.Account_BookName, "")
   recItem.PortfolioName  = get_String(rec, dm.Account_PortfolioName, "")
   recItem.Centre  = get_String(rec, dm.Account_Centre, "")
   recItem.Firm  = get_String(rec, dm.Account_Firm, "")
   recItem.CCYDp  = get_Int(rec, dm.Account_CCYDp, "0")





// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Account_NewID(r dm.Account) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

