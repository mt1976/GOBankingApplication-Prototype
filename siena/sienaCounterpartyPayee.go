package siena

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

// Defines the Fields to Fetch from SQL
var sienaCounterpartyPayeeSQL = "SourceTable, 	KeyCounterpartyFirm, 	KeyCounterpartyCentre, 	KeyCurrency, 	KeyName, 	KeyNumber, 	KeyDirection, 	KeyType, 	FullName, 	Address, 	PhoneNo, 	Country, 	Bic, 	Iban, 	AccountNo, 	FedWireNo, 	SortCode, 	BankName, 	BankPinCode, 	BankAddress, 	Reason, 	BankSettlementAcct, 	UpdatedUserId"
var sqlCPPYSourceTable, sqlCPPYKeyCounterpartyFirm, sqlCPPYKeyCounterpartyCentre, sqlCPPYKeyCurrency, sqlCPPYKeyName, sqlCPPYKeyNumber, sqlCPPYKeyDirection, sqlCPPYKeyType, sqlCPPYFullName, sqlCPPYAddress, sqlCPPYPhoneNo, sqlCPPYCountry, sqlCPPYBic, sqlCPPYIban, sqlCPPYAccountNo, sqlCPPYFedWireNo, sqlCPPYSortCode, sqlCPPYBankName, sqlCPPYBankPinCode, sqlCPPYBankAddress, sqlCPPYReason, sqlCPPYBankSettlementAcct, sqlCPPYUpdatedUserId sql.NullString

//sienaCounterpartyPayeePage is cheese
type sienaCounterpartyPayeeListPage struct {
	UserMenu                    []application.AppMenuItem
	UserRole                    string
	UserNavi                    string
	Title                       string
	PageTitle                   string
	SienaCounterpartyPayeeCount int
	SienaCounterpartyPayeeList  []sienaCounterpartyPayeeItem
}

//sienaCounterpartyPayeePage is cheese
type sienaCounterpartyPayeePage struct {
	UserMenu              []application.AppMenuItem
	UserRole              string
	UserNavi              string
	Title                 string
	PageTitle             string
	ID                    string
	SourceTable           string
	KeyCounterpartyFirm   string
	KeyCounterpartyCentre string
	KeyCurrency           string
	KeyName               string
	KeyNumber             string
	KeyDirection          string
	KeyType               string
	FullName              string
	Address               string
	PhoneNo               string
	Country               string
	Bic                   string
	Iban                  string
	AccountNo             string
	FedWireNo             string
	SortCode              string
	BankName              string
	BankPinCode           string
	BankAddress           string
	Reason                string
	BankSettlementAcct    string
	UpdatedUserId         string
	Approved              string

	Action      string
	CountryList []sienaCountryItem
}

//sienaCounterpartyPayeeItem is cheese
type sienaCounterpartyPayeeItem struct {
	SourceTable           string
	KeyCounterpartyFirm   string
	KeyCounterpartyCentre string
	KeyCurrency           string
	KeyName               string
	KeyNumber             string
	KeyDirection          string
	KeyType               string
	FullName              string
	Address               string
	PhoneNo               string
	Country               string
	Bic                   string
	Iban                  string
	AccountNo             string
	FedWireNo             string
	SortCode              string
	BankName              string
	BankPinCode           string
	BankAddress           string
	Reason                string
	BankSettlementAcct    string
	UpdatedUserId         string
	Approved              string
	Action                string
	CountryName           string
}

func listSienaCounterpartyPayeeHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "listSienaCounterpartyPayee"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyPayeeItem
	noItems, returnList, _ := getSienaCounterpartyPayeeList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCounterpartyPayeeList := sienaCounterpartyPayeeListPage{
		UserMenu:                    getappMenuData(gUserRole),
		UserRole:                    gUserRole,
		UserNavi:                    gUserNavi,
		Title:                       wctProperties["appname"],
		PageTitle:                   "List Siena CounterpartyPayees",
		SienaCounterpartyPayeeCount: noItems,
		SienaCounterpartyPayeeList:  returnList,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaCounterpartyPayeeList)

}

func viewSienaCounterpartyPayeeHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "viewSienaCounterpartyPayee"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyPayeeItem
	idSource := support.GetURLparam(r, "csrc")
	idFirm := support.GetURLparam(r, "cfrm")
	idCentre := support.GetURLparam(r, "ccen")
	idCCY := support.GetURLparam(r, "cccy")
	idName := support.GetURLparam(r, "cnam")
	idNumber := support.GetURLparam(r, "cnum")
	idDirection := support.GetURLparam(r, "cdir")
	idType := support.GetURLparam(r, "ctyp")

	noItems, returnRecord, _ := getSienaCounterpartyPayeeByKey(thisConnection, idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)
	fmt.Println("NoSienaItems", noItems, idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCounterpartyPayeeList := sienaCounterpartyPayeePage{
		UserMenu:              getappMenuData(gUserRole),
		UserRole:              gUserRole,
		UserNavi:              gUserNavi,
		Title:                 wctProperties["appname"],
		PageTitle:             "View Siena CounterpartyPayee",
		SourceTable:           returnRecord.SourceTable,
		KeyCounterpartyFirm:   returnRecord.KeyCounterpartyFirm,
		KeyCounterpartyCentre: returnRecord.KeyCounterpartyCentre,
		KeyCurrency:           returnRecord.KeyCurrency,
		KeyName:               returnRecord.KeyName,
		KeyNumber:             returnRecord.KeyNumber,
		KeyDirection:          returnRecord.KeyDirection,
		KeyType:               returnRecord.KeyType,
		FullName:              returnRecord.FullName,
		Address:               returnRecord.Address,
		PhoneNo:               returnRecord.PhoneNo,
		Country:               returnRecord.Country,
		Bic:                   returnRecord.Bic,
		Iban:                  returnRecord.Iban,
		AccountNo:             returnRecord.AccountNo,
		FedWireNo:             returnRecord.FedWireNo,
		SortCode:              returnRecord.SortCode,
		BankName:              returnRecord.BankName,
		BankPinCode:           returnRecord.BankPinCode,
		BankAddress:           returnRecord.BankAddress,
		Reason:                returnRecord.Reason,
		BankSettlementAcct:    returnRecord.BankSettlementAcct,
		UpdatedUserId:         returnRecord.UpdatedUserId,
		Approved:              returnRecord.Approved,
		Action:                "",
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaCounterpartyPayeeList)

}

func editSienaCounterpartyPayeeHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "editSienaCounterpartyPayee"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaCounterpartyPayeeItem
	idSource := support.GetURLparam(r, "csrc")
	idFirm := support.GetURLparam(r, "cfrm")
	idCentre := support.GetURLparam(r, "ccen")
	idCCY := support.GetURLparam(r, "cccy")
	idName := support.GetURLparam(r, "cnam")
	idNumber := support.GetURLparam(r, "cnum")
	idDirection := support.GetURLparam(r, "cdir")
	idType := support.GetURLparam(r, "ctyp")

	noItems, returnRecord, _ := getSienaCounterpartyPayeeByKey(thisConnection, idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)
	fmt.Println("NoSienaItems", noItems, idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)

	//fmt.Println(displayList)

	pageSienaCounterpartyPayeeList := sienaCounterpartyPayeePage{
		UserMenu:              getappMenuData(gUserRole),
		UserRole:              gUserRole,
		UserNavi:              gUserNavi,
		Title:                 wctProperties["appname"],
		PageTitle:             "View Siena CounterpartyPayee",
		SourceTable:           returnRecord.SourceTable,
		KeyCounterpartyFirm:   returnRecord.KeyCounterpartyFirm,
		KeyCounterpartyCentre: returnRecord.KeyCounterpartyCentre,
		KeyCurrency:           returnRecord.KeyCurrency,
		KeyName:               returnRecord.KeyName,
		KeyNumber:             returnRecord.KeyNumber,
		KeyDirection:          returnRecord.KeyDirection,
		KeyType:               returnRecord.KeyType,
		FullName:              returnRecord.FullName,
		Address:               returnRecord.Address,
		PhoneNo:               returnRecord.PhoneNo,
		Country:               returnRecord.Country,
		Bic:                   returnRecord.Bic,
		Iban:                  returnRecord.Iban,
		AccountNo:             returnRecord.AccountNo,
		FedWireNo:             returnRecord.FedWireNo,
		SortCode:              returnRecord.SortCode,
		BankName:              returnRecord.BankName,
		BankPinCode:           returnRecord.BankPinCode,
		BankAddress:           returnRecord.BankAddress,
		Reason:                returnRecord.Reason,
		BankSettlementAcct:    returnRecord.BankSettlementAcct,
		UpdatedUserId:         returnRecord.UpdatedUserId,
		Approved:              returnRecord.Approved,
		Action:                "",
		CountryList:           countryList,
	}
	//	fmt.Println(pageSienaCounterpartyPayeeList)

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaCounterpartyPayeeList)

}

func saveSienaCounterpartyPayeeHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := support.GetProperties(SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCounterpartyPayeeItem

	item.SourceTable = r.FormValue("SourceTable")
	item.KeyCounterpartyFirm = r.FormValue("KeyCounterpartyFirm")
	item.KeyCounterpartyCentre = r.FormValue("KeyCounterpartyCentre")
	item.KeyCurrency = r.FormValue("KeyCurrency")
	item.KeyName = r.FormValue("KeyName")
	item.KeyNumber = r.FormValue("KeyNumber")
	item.KeyDirection = r.FormValue("KeyDirection")
	item.KeyType = r.FormValue("KeyType")
	item.FullName = r.FormValue("FullName")
	item.Address = r.FormValue("Address")
	item.PhoneNo = r.FormValue("PhoneNo")
	item.Country = r.FormValue("Country")
	item.Bic = r.FormValue("Bic")
	item.Iban = r.FormValue("Iban")
	item.AccountNo = r.FormValue("AccountNo")
	item.FedWireNo = r.FormValue("FedWireNo")
	item.SortCode = r.FormValue("SortCode")
	item.BankName = r.FormValue("BankName")
	item.BankPinCode = r.FormValue("BankPinCode")
	item.BankAddress = r.FormValue("BankAddress")
	item.Reason = r.FormValue("Reason")
	item.BankSettlementAcct = r.FormValue("BankSettlementAcct")
	item.UpdatedUserId = r.FormValue("UpdatedUserId")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.SourceTable

	sFldName.Name = "fullName"
	sFldName.Text = item.FullName

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "CounterpartyPayee"
	sienaTable.Classname = "com.eurobase.siena.data.CounterpartyPayees.CounterpartyPayee"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	listSienaCounterpartyPayeeHandler(w, r)

}

func newSienaCounterpartyPayeeHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "newSienaCounterpartyPayee"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaCounterpartyPayeeList := sienaCounterpartyPayeePage{
		UserMenu:              getappMenuData(gUserRole),
		UserRole:              gUserRole,
		UserNavi:              gUserNavi,
		Title:                 wctProperties["appname"],
		PageTitle:             "View Siena CounterpartyPayee",
		ID:                    "NEW",
		SourceTable:           "",
		KeyCounterpartyFirm:   "",
		KeyCounterpartyCentre: "",
		KeyCurrency:           "",
		KeyName:               "",
		KeyNumber:             "",
		KeyDirection:          "",
		KeyType:               "",
		FullName:              "",
		Address:               "",
		PhoneNo:               "",
		Country:               "",
		Bic:                   "",
		Iban:                  "",
		AccountNo:             "",
		FedWireNo:             "",
		SortCode:              "",
		BankName:              "",
		BankPinCode:           "",
		BankAddress:           "",
		Reason:                "",
		BankSettlementAcct:    "",
		UpdatedUserId:         "",
		Approved:              "Pending",

		Action: "NEW",
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaCounterpartyPayeeList)

}

// getSienaCounterpartyPayeeList read all employees
func getSienaCounterpartyPayeeList(db *sql.DB) (int, []sienaCounterpartyPayeeItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyPayee ORDER BY KeyCounterpartyFirm, KeyCounterpartyCentre, KeyCurrency DESC;", sienaCounterpartyPayeeSQL, mssqlConfig["schema"])
	count, sienaCounterpartyPayeeList, _, _ := fetchSienaCounterpartyPayeeData(db, tsql)
	return count, sienaCounterpartyPayeeList, nil
}

// getSienaCounterpartyPayeeList read all employees
func getSienaCounterpartyPayee(db *sql.DB, id string) (int, sienaCounterpartyPayeeItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyPayee WHERE Code='%s' ORDER BY KeyCounterpartyFirm, KeyCounterpartyCentre, KeyCurrency DESC;", sienaCounterpartyPayeeSQL, mssqlConfig["schema"], id)
	_, _, sienaCounterpartyPayee, _ := fetchSienaCounterpartyPayeeData(db, tsql)
	return 1, sienaCounterpartyPayee, nil
}

// getSienaCounterpartyPayeeList read all employees
func getSienaCounterpartyPayeeListByCounterparty(db *sql.DB, idFirm string, idCentre string) (int, []sienaCounterpartyPayeeItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyPayee WHERE KeyCounterpartyFirm='%s' AND KeyCounterpartyCentre='%s' ORDER BY KeyCounterpartyFirm, KeyCounterpartyCentre, KeyCurrency DESC;", sienaCounterpartyPayeeSQL, mssqlConfig["schema"], idFirm, idCentre)
	count, sienaCounterpartyPayeeList, _, _ := fetchSienaCounterpartyPayeeData(db, tsql)
	return count, sienaCounterpartyPayeeList, nil
}

func getSienaCounterpartyPayeeByKey(db *sql.DB, idSource string, idFirm string, idCentre string, idCCY string, idName string, idNumber string, idDirection string, idType string) (int, sienaCounterpartyPayeeItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyPayee WHERE SourceTable='%s' AND KeyCounterpartyFirm='%s' AND KeyCounterpartyCentre='%s' AND KeyCurrency='%s' AND KeyName='%s' AND KeyNumber='%s' AND KeyDirection='%s' AND KeyType='%s' ORDER BY KeyCounterpartyFirm, KeyCounterpartyCentre, KeyCurrency DESC;", sienaCounterpartyPayeeSQL, mssqlConfig["schema"], idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)
	_, _, sienaCounterpartyPayee, _ := fetchSienaCounterpartyPayeeData(db, tsql)
	return 1, sienaCounterpartyPayee, nil
}

// getSienaCounterpartyPayeeList read all employees
func putSienaCounterpartyPayee(db *sql.DB, updateItem sienaCounterpartyPayeeItem) error {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCounterpartyPayeeData read all employees
func fetchSienaCounterpartyPayeeData(db *sql.DB, tsql string) (int, []sienaCounterpartyPayeeItem, sienaCounterpartyPayeeItem, error) {

	var sienaCounterpartyPayee sienaCounterpartyPayeeItem
	var sienaCounterpartyPayeeList []sienaCounterpartyPayeeItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCounterpartyPayee, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCPPYSourceTable, &sqlCPPYKeyCounterpartyFirm, &sqlCPPYKeyCounterpartyCentre, &sqlCPPYKeyCurrency, &sqlCPPYKeyName, &sqlCPPYKeyNumber, &sqlCPPYKeyDirection, &sqlCPPYKeyType, &sqlCPPYFullName, &sqlCPPYAddress, &sqlCPPYPhoneNo, &sqlCPPYCountry, &sqlCPPYBic, &sqlCPPYIban, &sqlCPPYAccountNo, &sqlCPPYFedWireNo, &sqlCPPYSortCode, &sqlCPPYBankName, &sqlCPPYBankPinCode, &sqlCPPYBankAddress, &sqlCPPYReason, &sqlCPPYBankSettlementAcct, &sqlCPPYUpdatedUserId)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCounterpartyPayee, err
		}

		sienaCounterpartyPayee.SourceTable = sqlCPPYSourceTable.String
		sienaCounterpartyPayee.KeyCounterpartyFirm = sqlCPPYKeyCounterpartyFirm.String
		sienaCounterpartyPayee.KeyCounterpartyCentre = sqlCPPYKeyCounterpartyCentre.String
		sienaCounterpartyPayee.KeyCurrency = sqlCPPYKeyCurrency.String
		sienaCounterpartyPayee.KeyName = sqlCPPYKeyName.String
		sienaCounterpartyPayee.KeyNumber = sqlCPPYKeyNumber.String
		sienaCounterpartyPayee.KeyDirection = sqlCPPYKeyDirection.String
		sienaCounterpartyPayee.KeyType = sqlCPPYKeyType.String
		sienaCounterpartyPayee.FullName = sqlCPPYFullName.String
		sienaCounterpartyPayee.Address = sqlCPPYAddress.String
		sienaCounterpartyPayee.PhoneNo = sqlCPPYPhoneNo.String
		sienaCounterpartyPayee.Country = sqlCPPYCountry.String
		sienaCounterpartyPayee.Bic = sqlCPPYBic.String
		sienaCounterpartyPayee.Iban = sqlCPPYIban.String
		sienaCounterpartyPayee.AccountNo = sqlCPPYAccountNo.String
		sienaCounterpartyPayee.FedWireNo = sqlCPPYFedWireNo.String
		sienaCounterpartyPayee.SortCode = sqlCPPYSortCode.String
		sienaCounterpartyPayee.BankName = sqlCPPYBankName.String
		sienaCounterpartyPayee.BankPinCode = sqlCPPYBankPinCode.String
		sienaCounterpartyPayee.BankAddress = sqlCPPYBankAddress.String
		sienaCounterpartyPayee.Reason = sqlCPPYReason.String
		sienaCounterpartyPayee.BankSettlementAcct = sqlCPPYBankSettlementAcct.String
		sienaCounterpartyPayee.UpdatedUserId = sqlCPPYUpdatedUserId.String

		sienaCounterpartyPayee.Approved = "Pending"
		if sienaCounterpartyPayee.SourceTable == "dbo.Payee" {
			sienaCounterpartyPayee.Approved = "Approved"
		}

		sienaCounterpartyPayeeList = append(sienaCounterpartyPayeeList, sienaCounterpartyPayee)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyPayeeList, sienaCounterpartyPayee, nil
}
