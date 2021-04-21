package main

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
)

var sienaCounterpartyExtensionsSQL = "NameFirm, 	NameCentre, 	BICCode, 	ContactIndicator, 	CoverTrade, 	CustomerCategory, 	FSCSInclusive, 	FeeFactor, 	InactiveStatus, 	Indemnity, 	KnowYourCustomerStatus, 	LERLimitCarveOut, 	LastAmended, 	LastLogin, 	LossGivenDefault, 	MIC, 	ProtectedDepositor, 	RPTCurrency, 	RateTimeout, 	RateValidation, 	Registered, 	RegulatoryCategory, 	SecuredSettlement, 	SettlementLimitCarveOut, 	SortCode, 	Training, 	TrainingCode, 	TrainingReceived, 	Unencumbered, 	LEIExpiryDate, 	MIFIDReviewDate, 	GDPRReviewDate, 	DelegatedReporting, 	BOReconcile, 	MIFIDReportableDealsAllowed, 	SignedInvestmentAgreement, 	AppropriatenessAssessment, 	FinancialCounterparty, 	Collateralisation, 	PortfolioCode, 	ReconciliationLetterFrequency, 	DirectDealing"

var sqlCPEXNameFirm, sqlCPEXNameCentre, sqlCPEXBICCode, sqlCPEXContactIndicator, sqlCPEXCoverTrade, sqlCPEXCustomerCategory, sqlCPEXFSCSInclusive, sqlCPEXFeeFactor, sqlCPEXInactiveStatus, sqlCPEXIndemnity, sqlCPEXKnowYourCustomerStatus, sqlCPEXLERLimitCarveOut, sqlCPEXLastAmended, sqlCPEXLastLogin, sqlCPEXLossGivenDefault, sqlCPEXMIC, sqlCPEXProtectedDepositor, sqlCPEXRPTCurrency, sqlCPEXRateTimeout, sqlCPEXRateValidation, sqlCPEXRegistered, sqlCPEXRegulatoryCategory, sqlCPEXSecuredSettlement, sqlCPEXSettlementLimitCarveOut, sqlCPEXSortCode, sqlCPEXTraining, sqlCPEXTrainingCode, sqlCPEXTrainingReceived, sqlCPEXUnencumbered, sqlCPEXLEIExpiryDate, sqlCPEXMIFIDReviewDate, sqlCPEXGDPRReviewDate, sqlCPEXDelegatedReporting, sqlCPEXBOReconcile, sqlCPEXMIFIDReportableDealsAllowed, sqlCPEXSignedInvestmentAgreement, sqlCPEXAppropriatenessAssessment, sqlCPEXFinancialCounterparty, sqlCPEXCollateralisation, sqlCPEXPortfolioCode, sqlCPEXReconciliationLetterFrequency, sqlCPEXDirectDealing sql.NullString

//sienaCounterpartyExtensionsPage is cheese
type sienaCounterpartyExtensionsListPage struct {
	UserMenu                         []AppMenuItem
	UserRole                         string
	UserNavi                         string
	Title                            string
	PageTitle                        string
	SienaCounterpartyExtensionsCount int
	SienaCounterpartyExtensionsList  []sienaCounterpartyExtensionsItem
}

//sienaCounterpartyExtensionsPage is cheese
type sienaCounterpartyExtensionsPage struct {
	UserMenu                      []AppMenuItem
	UserRole                      string
	UserNavi                      string
	Title                         string
	PageTitle                     string
	ID                            string
	Action                        string
	NameFirm                      string
	NameCentre                    string
	BICCode                       string
	ContactIndicator              string
	CoverTrade                    string
	CustomerCategory              string
	FSCSInclusive                 string
	FeeFactor                     string
	InactiveStatus                string
	Indemnity                     string
	KnowYourCustomerStatus        string
	LERLimitCarveOut              string
	LastAmended                   string
	LastLogin                     string
	LossGivenDefault              string
	MIC                           string
	ProtectedDepositor            string
	RPTCurrency                   string
	RateTimeout                   string
	RateValidation                string
	Registered                    string
	RegulatoryCategory            string
	SecuredSettlement             string
	SettlementLimitCarveOut       string
	SortCode                      string
	Training                      string
	TrainingCode                  string
	TrainingReceived              string
	Unencumbered                  string
	LEIExpiryDate                 string
	MIFIDReviewDate               string
	GDPRReviewDate                string
	DelegatedReporting            string
	BOReconcile                   string
	MIFIDReportableDealsAllowed   string
	SignedInvestmentAgreement     string
	AppropriatenessAssessment     string
	FinancialCounterparty         string
	Collateralisation             string
	PortfolioCode                 string
	ReconciliationLetterFrequency string
	DirectDealing                 string
	YNList                        []sienaYNItem
}

//sienaCounterpartyExtensionsItem is cheese
type sienaCounterpartyExtensionsItem struct {
	NameFirm                      string
	NameCentre                    string
	BICCode                       string
	ContactIndicator              string
	CoverTrade                    string
	CustomerCategory              string
	FSCSInclusive                 string
	FeeFactor                     string
	InactiveStatus                string
	Indemnity                     string
	KnowYourCustomerStatus        string
	LERLimitCarveOut              string
	LastAmended                   string
	LastLogin                     string
	LossGivenDefault              string
	MIC                           string
	ProtectedDepositor            string
	RPTCurrency                   string
	RateTimeout                   string
	RateValidation                string
	Registered                    string
	RegulatoryCategory            string
	SecuredSettlement             string
	SettlementLimitCarveOut       string
	SortCode                      string
	Training                      string
	TrainingCode                  string
	TrainingReceived              string
	Unencumbered                  string
	LEIExpiryDate                 string
	MIFIDReviewDate               string
	GDPRReviewDate                string
	DelegatedReporting            string
	BOReconcile                   string
	MIFIDReportableDealsAllowed   string
	SignedInvestmentAgreement     string
	AppropriatenessAssessment     string
	FinancialCounterparty         string
	Collateralisation             string
	PortfolioCode                 string
	ReconciliationLetterFrequency string
	DirectDealing                 string
	Action                        string
}

func listSienaCounterpartyExtensionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "listSienaCounterpartyExtensions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyExtensionsItem
	noItems, returnList, _ := getSienaCounterpartyExtensionsList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCounterpartyExtensionsList := sienaCounterpartyExtensionsListPage{
		UserMenu:                         getappMenuData(gUserRole),
		UserRole:                         gUserRole,
		UserNavi:                         gUserNavi,
		Title:                            wctProperties["appname"],
		PageTitle:                        "List Siena CounterpartyExtensionss",
		SienaCounterpartyExtensionsCount: noItems,
		SienaCounterpartyExtensionsList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyExtensionsList)

}

func viewSienaCounterpartyExtensionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "viewSienaCounterpartyExtensions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyExtensionsItem
	suID := getURLparam(r, "SU")
	sfID := getURLparam(r, "SF")
	scID := getURLparam(r, "SC")
	noItems, returnRecord, _ := getSienaCounterpartyExtensions(thisConnection, sfID, scID)
	fmt.Println("NoSienaItems", noItems, suID, sfID, scID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCounterpartyExtensionsList := sienaCounterpartyExtensionsPage{
		UserMenu:                      getappMenuData(gUserRole),
		UserRole:                      gUserRole,
		UserNavi:                      gUserNavi,
		Title:                         wctProperties["appname"],
		PageTitle:                     "View Siena CounterpartyExtensions",
		NameFirm:                      returnRecord.NameFirm,
		NameCentre:                    returnRecord.NameCentre,
		BICCode:                       returnRecord.BICCode,
		ContactIndicator:              returnRecord.ContactIndicator,
		CoverTrade:                    returnRecord.CoverTrade,
		CustomerCategory:              returnRecord.CustomerCategory,
		FSCSInclusive:                 returnRecord.FSCSInclusive,
		FeeFactor:                     returnRecord.FeeFactor,
		InactiveStatus:                returnRecord.InactiveStatus,
		Indemnity:                     returnRecord.Indemnity,
		KnowYourCustomerStatus:        returnRecord.KnowYourCustomerStatus,
		LERLimitCarveOut:              returnRecord.LERLimitCarveOut,
		LastAmended:                   returnRecord.LastAmended,
		LastLogin:                     returnRecord.LastLogin,
		LossGivenDefault:              returnRecord.LossGivenDefault,
		MIC:                           returnRecord.MIC,
		ProtectedDepositor:            returnRecord.ProtectedDepositor,
		RPTCurrency:                   returnRecord.RPTCurrency,
		RateTimeout:                   returnRecord.RateTimeout,
		RateValidation:                returnRecord.RateValidation,
		Registered:                    returnRecord.Registered,
		RegulatoryCategory:            returnRecord.RegulatoryCategory,
		SecuredSettlement:             returnRecord.SecuredSettlement,
		SettlementLimitCarveOut:       returnRecord.SettlementLimitCarveOut,
		SortCode:                      returnRecord.SortCode,
		Training:                      returnRecord.Training,
		TrainingCode:                  returnRecord.TrainingCode,
		TrainingReceived:              returnRecord.TrainingReceived,
		Unencumbered:                  returnRecord.Unencumbered,
		LEIExpiryDate:                 returnRecord.LEIExpiryDate,
		MIFIDReviewDate:               returnRecord.MIFIDReviewDate,
		GDPRReviewDate:                returnRecord.GDPRReviewDate,
		DelegatedReporting:            returnRecord.DelegatedReporting,
		BOReconcile:                   returnRecord.BOReconcile,
		MIFIDReportableDealsAllowed:   returnRecord.MIFIDReportableDealsAllowed,
		SignedInvestmentAgreement:     returnRecord.SignedInvestmentAgreement,
		AppropriatenessAssessment:     returnRecord.AppropriatenessAssessment,
		FinancialCounterparty:         returnRecord.FinancialCounterparty,
		Collateralisation:             returnRecord.Collateralisation,
		PortfolioCode:                 returnRecord.PortfolioCode,
		ReconciliationLetterFrequency: returnRecord.ReconciliationLetterFrequency,
		DirectDealing:                 returnRecord.DirectDealing,
		Action:                        "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyExtensionsList)

}

func editSienaCounterpartyExtensionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "editSienaCounterpartyExtensions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyExtensionsItem
	sfID := getURLparam(r, "SienaFirm")
	scID := getURLparam(r, "SienaCentre")
	noItems, returnRecord, _ := getSienaCounterpartyExtensions(thisConnection, sfID, scID)
	fmt.Println("NoSienaItems", noItems, sfID, scID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, ynList, _ := getSienaYNList()
	//fmt.Println(displayList)

	pageSienaCounterpartyExtensionsList := sienaCounterpartyExtensionsPage{
		UserMenu:                      getappMenuData(gUserRole),
		UserRole:                      gUserRole,
		UserNavi:                      gUserNavi,
		Title:                         wctProperties["appname"],
		PageTitle:                     "View Siena CounterpartyExtensions",
		NameFirm:                      returnRecord.NameFirm,
		NameCentre:                    returnRecord.NameCentre,
		BICCode:                       returnRecord.BICCode,
		ContactIndicator:              returnRecord.ContactIndicator,
		CoverTrade:                    returnRecord.CoverTrade,
		CustomerCategory:              returnRecord.CustomerCategory,
		FSCSInclusive:                 returnRecord.FSCSInclusive,
		FeeFactor:                     returnRecord.FeeFactor,
		InactiveStatus:                returnRecord.InactiveStatus,
		Indemnity:                     returnRecord.Indemnity,
		KnowYourCustomerStatus:        returnRecord.KnowYourCustomerStatus,
		LERLimitCarveOut:              returnRecord.LERLimitCarveOut,
		LastAmended:                   returnRecord.LastAmended,
		LastLogin:                     returnRecord.LastLogin,
		LossGivenDefault:              returnRecord.LossGivenDefault,
		MIC:                           returnRecord.MIC,
		ProtectedDepositor:            returnRecord.ProtectedDepositor,
		RPTCurrency:                   returnRecord.RPTCurrency,
		RateTimeout:                   returnRecord.RateTimeout,
		RateValidation:                returnRecord.RateValidation,
		Registered:                    returnRecord.Registered,
		RegulatoryCategory:            returnRecord.RegulatoryCategory,
		SecuredSettlement:             returnRecord.SecuredSettlement,
		SettlementLimitCarveOut:       returnRecord.SettlementLimitCarveOut,
		SortCode:                      returnRecord.SortCode,
		Training:                      returnRecord.Training,
		TrainingCode:                  returnRecord.TrainingCode,
		TrainingReceived:              returnRecord.TrainingReceived,
		Unencumbered:                  returnRecord.Unencumbered,
		LEIExpiryDate:                 returnRecord.LEIExpiryDate,
		MIFIDReviewDate:               returnRecord.MIFIDReviewDate,
		GDPRReviewDate:                returnRecord.GDPRReviewDate,
		DelegatedReporting:            returnRecord.DelegatedReporting,
		BOReconcile:                   returnRecord.BOReconcile,
		MIFIDReportableDealsAllowed:   returnRecord.MIFIDReportableDealsAllowed,
		SignedInvestmentAgreement:     returnRecord.SignedInvestmentAgreement,
		AppropriatenessAssessment:     returnRecord.AppropriatenessAssessment,
		FinancialCounterparty:         returnRecord.FinancialCounterparty,
		Collateralisation:             returnRecord.Collateralisation,
		PortfolioCode:                 returnRecord.PortfolioCode,
		ReconciliationLetterFrequency: returnRecord.ReconciliationLetterFrequency,
		DirectDealing:                 returnRecord.DirectDealing,
		Action:                        "",
		YNList:                        ynList,
	}

	fmt.Println(pageSienaCounterpartyExtensionsList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyExtensionsList)

}

func saveSienaCounterpartyExtensionsHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(SIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCounterpartyExtensionsItem

	item.NameFirm = r.FormValue("NameFirm")
	item.NameCentre = r.FormValue("NameCentre")
	item.BICCode = r.FormValue("BICCode")
	item.ContactIndicator = r.FormValue("ContactIndicator")
	item.CoverTrade = r.FormValue("CoverTrade")
	item.CustomerCategory = r.FormValue("CustomerCategory")
	item.FSCSInclusive = r.FormValue("FSCSInclusive")
	item.FeeFactor = r.FormValue("FeeFactor")
	item.InactiveStatus = r.FormValue("InactiveStatus")
	item.Indemnity = r.FormValue("Indemnity")
	item.KnowYourCustomerStatus = r.FormValue("KnowYourCustomerStatus")
	item.LERLimitCarveOut = r.FormValue("LERLimitCarveOut")
	item.LastAmended = r.FormValue("LastAmended")
	item.LastLogin = r.FormValue("LastLogin")
	item.LossGivenDefault = r.FormValue("LossGivenDefault")
	item.MIC = r.FormValue("MIC")
	item.ProtectedDepositor = r.FormValue("ProtectedDepositor")
	item.RPTCurrency = r.FormValue("RPTCurrency")
	item.RateTimeout = r.FormValue("RateTimeout")
	item.RateValidation = r.FormValue("RateValidation")
	item.Registered = r.FormValue("Registered")
	item.RegulatoryCategory = r.FormValue("RegulatoryCategory")
	item.SecuredSettlement = r.FormValue("SecuredSettlement")
	item.SettlementLimitCarveOut = r.FormValue("SettlementLimitCarveOut")
	item.SortCode = r.FormValue("SortCode")
	item.Training = r.FormValue("Training")
	item.TrainingCode = r.FormValue("TrainingCode")
	item.TrainingReceived = r.FormValue("TrainingReceived")
	item.Unencumbered = r.FormValue("Unencumbered")
	item.LEIExpiryDate = r.FormValue("LEIExpiryDate")
	item.MIFIDReviewDate = r.FormValue("MIFIDReviewDate")
	item.GDPRReviewDate = r.FormValue("GDPRReviewDate")
	item.DelegatedReporting = r.FormValue("DelegatedReporting")
	item.BOReconcile = r.FormValue("BOReconcile")
	item.MIFIDReportableDealsAllowed = r.FormValue("MIFIDReportableDealsAllowed")
	item.SignedInvestmentAgreement = r.FormValue("SignedInvestmentAgreement")
	item.AppropriatenessAssessment = r.FormValue("AppropriatenessAssessment")
	item.FinancialCounterparty = r.FormValue("FinancialCounterparty")
	item.Collateralisation = r.FormValue("Collateralisation")
	item.PortfolioCode = r.FormValue("PortfolioCode")
	item.ReconciliationLetterFrequency = r.FormValue("ReconciliationLetterFrequency")
	item.DirectDealing = r.FormValue("DirectDealing")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "CounterpartyExtensionsKeyUserName"
	sFldCode.Text = item.NameFirm

	sFldName.Name = "CounterpartyExtensionsKeyCounterpartyFirm"
	sFldName.Text = item.NameCentre

	sFldCountry.Name = "CounterpartyExtensionsKeyCounterpartyCentre"
	sFldCountry.Text = item.NameCentre

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
	sienaTable.Name = "CounterpartyExtensions"
	sienaTable.Classname = "com.eurobase.siena.data.CounterpartyExtensionss.CounterpartyExtensions"
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

	viewSienaCounterpartyHandler(w, r)

}

func newSienaCounterpartyExtensionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(APPCONFIG)
	tmpl := "newSienaCounterpartyExtensions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items

	pageSienaCounterpartyExtensionsList := sienaCounterpartyExtensionsPage{
		UserMenu:                      getappMenuData(gUserRole),
		UserRole:                      gUserRole,
		UserNavi:                      gUserNavi,
		Title:                         wctProperties["appname"],
		PageTitle:                     "View Siena CounterpartyExtensions",
		ID:                            "NEW",
		NameFirm:                      "",
		NameCentre:                    "",
		BICCode:                       "",
		ContactIndicator:              "",
		CoverTrade:                    "",
		CustomerCategory:              "",
		FSCSInclusive:                 "",
		FeeFactor:                     "",
		InactiveStatus:                "",
		Indemnity:                     "",
		KnowYourCustomerStatus:        "",
		LERLimitCarveOut:              "",
		LastAmended:                   "",
		LastLogin:                     "",
		LossGivenDefault:              "",
		MIC:                           "",
		ProtectedDepositor:            "",
		RPTCurrency:                   "",
		RateTimeout:                   "",
		RateValidation:                "",
		Registered:                    "",
		RegulatoryCategory:            "",
		SecuredSettlement:             "",
		SettlementLimitCarveOut:       "",
		SortCode:                      "",
		Training:                      "",
		TrainingCode:                  "",
		TrainingReceived:              "",
		Unencumbered:                  "",
		LEIExpiryDate:                 "",
		MIFIDReviewDate:               "",
		GDPRReviewDate:                "",
		DelegatedReporting:            "",
		BOReconcile:                   "",
		MIFIDReportableDealsAllowed:   "",
		SignedInvestmentAgreement:     "",
		AppropriatenessAssessment:     "",
		FinancialCounterparty:         "",
		Collateralisation:             "",
		PortfolioCode:                 "",
		ReconciliationLetterFrequency: "",
		DirectDealing:                 "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyExtensionsList)

}

// getSienaCounterpartyExtensionsList read all employees
func getSienaCounterpartyExtensionsList(db *sql.DB) (int, []sienaCounterpartyExtensionsItem, error) {
	mssqlConfig := getProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyExtensions;", sienaCounterpartyExtensionsSQL, mssqlConfig["schema"])
	count, sienaCounterpartyExtensionsList, _, _ := fetchSienaCounterpartyExtensionsData(db, tsql)
	return count, sienaCounterpartyExtensionsList, nil
}

// getSienaCounterpartyExtensionsList read all employees
func getSienaCounterpartyExtensionsListByCounterparty(db *sql.DB, idFirm string, idCentre string) (int, []sienaCounterpartyExtensionsItem, error) {
	mssqlConfig := getProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyExtensions WHERE NameFirm='%s' AND NameCentre='%s';", sienaCounterpartyExtensionsSQL, mssqlConfig["schema"], idFirm, idCentre)
	count, sienaCounterpartyExtensionsList, _, _ := fetchSienaCounterpartyExtensionsData(db, tsql)
	return count, sienaCounterpartyExtensionsList, nil
}

// getSienaCounterpartyExtensionsList read all employees
func getSienaCounterpartyExtensions(db *sql.DB, sfid string, scid string) (int, sienaCounterpartyExtensionsItem, error) {
	mssqlConfig := getProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyExtensions WHERE NameFirm='%s' AND NameCentre='%s';", sienaCounterpartyExtensionsSQL, mssqlConfig["schema"], sfid, scid)
	_, _, sienaCounterpartyExtensions, _ := fetchSienaCounterpartyExtensionsData(db, tsql)
	return 1, sienaCounterpartyExtensions, nil
}

// getSienaCounterpartyExtensionsList read all employees
func putSienaCounterpartyExtensions(db *sql.DB, updateItem sienaCounterpartyExtensionsItem) error {
	mssqlConfig := getProperties(SQLCONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCounterpartyExtensionsData read all employees
func fetchSienaCounterpartyExtensionsData(db *sql.DB, tsql string) (int, []sienaCounterpartyExtensionsItem, sienaCounterpartyExtensionsItem, error) {

	var sienaCounterpartyExtensions sienaCounterpartyExtensionsItem
	var sienaCounterpartyExtensionsList []sienaCounterpartyExtensionsItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCounterpartyExtensions, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCPEXNameFirm, &sqlCPEXNameCentre, &sqlCPEXBICCode, &sqlCPEXContactIndicator, &sqlCPEXCoverTrade, &sqlCPEXCustomerCategory, &sqlCPEXFSCSInclusive, &sqlCPEXFeeFactor, &sqlCPEXInactiveStatus, &sqlCPEXIndemnity, &sqlCPEXKnowYourCustomerStatus, &sqlCPEXLERLimitCarveOut, &sqlCPEXLastAmended, &sqlCPEXLastLogin, &sqlCPEXLossGivenDefault, &sqlCPEXMIC, &sqlCPEXProtectedDepositor, &sqlCPEXRPTCurrency, &sqlCPEXRateTimeout, &sqlCPEXRateValidation, &sqlCPEXRegistered, &sqlCPEXRegulatoryCategory, &sqlCPEXSecuredSettlement, &sqlCPEXSettlementLimitCarveOut, &sqlCPEXSortCode, &sqlCPEXTraining, &sqlCPEXTrainingCode, &sqlCPEXTrainingReceived, &sqlCPEXUnencumbered, &sqlCPEXLEIExpiryDate, &sqlCPEXMIFIDReviewDate, &sqlCPEXGDPRReviewDate, &sqlCPEXDelegatedReporting, &sqlCPEXBOReconcile, &sqlCPEXMIFIDReportableDealsAllowed, &sqlCPEXSignedInvestmentAgreement, &sqlCPEXAppropriatenessAssessment, &sqlCPEXFinancialCounterparty, &sqlCPEXCollateralisation, &sqlCPEXPortfolioCode, &sqlCPEXReconciliationLetterFrequency, &sqlCPEXDirectDealing)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCounterpartyExtensions, err
		}

		sienaCounterpartyExtensions.NameFirm = sqlCPEXNameFirm.String
		sienaCounterpartyExtensions.NameCentre = sqlCPEXNameCentre.String
		sienaCounterpartyExtensions.BICCode = sqlCPEXBICCode.String
		sienaCounterpartyExtensions.ContactIndicator = sqlCPEXContactIndicator.String
		sienaCounterpartyExtensions.CoverTrade = sqlCPEXCoverTrade.String
		sienaCounterpartyExtensions.CustomerCategory = sqlCPEXCustomerCategory.String
		sienaCounterpartyExtensions.FSCSInclusive = sienaYN(sqlCPEXFSCSInclusive.String)
		sienaCounterpartyExtensions.FeeFactor = sqlCPEXFeeFactor.String
		sienaCounterpartyExtensions.InactiveStatus = sqlCPEXInactiveStatus.String
		sienaCounterpartyExtensions.Indemnity = sqlCPEXIndemnity.String
		sienaCounterpartyExtensions.KnowYourCustomerStatus = sienaYN(sqlCPEXKnowYourCustomerStatus.String)
		sienaCounterpartyExtensions.LERLimitCarveOut = sqlCPEXLERLimitCarveOut.String
		sienaCounterpartyExtensions.LastAmended = sqlCPEXLastAmended.String
		sienaCounterpartyExtensions.LastLogin = sqlCPEXLastLogin.String
		sienaCounterpartyExtensions.LossGivenDefault = sqlCPEXLossGivenDefault.String
		sienaCounterpartyExtensions.MIC = sqlCPEXMIC.String
		sienaCounterpartyExtensions.ProtectedDepositor = sqlCPEXProtectedDepositor.String
		sienaCounterpartyExtensions.RPTCurrency = sqlCPEXRPTCurrency.String
		sienaCounterpartyExtensions.RateTimeout = sqlCPEXRateTimeout.String
		sienaCounterpartyExtensions.RateValidation = sqlCPEXRateValidation.String
		sienaCounterpartyExtensions.Registered = sienaYN(sqlCPEXRegistered.String)
		sienaCounterpartyExtensions.RegulatoryCategory = sqlCPEXRegulatoryCategory.String
		sienaCounterpartyExtensions.SecuredSettlement = sqlCPEXSecuredSettlement.String
		sienaCounterpartyExtensions.SettlementLimitCarveOut = sqlCPEXSettlementLimitCarveOut.String
		sienaCounterpartyExtensions.SortCode = sqlCPEXSortCode.String
		sienaCounterpartyExtensions.Training = sienaYN(sqlCPEXTraining.String)
		sienaCounterpartyExtensions.TrainingCode = sqlCPEXTrainingCode.String
		sienaCounterpartyExtensions.TrainingReceived = sienaYN(sqlCPEXTrainingReceived.String)
		sienaCounterpartyExtensions.Unencumbered = sqlCPEXUnencumbered.String
		sienaCounterpartyExtensions.LEIExpiryDate = sqlDateToHTMLDate(sqlCPEXLEIExpiryDate.String)
		sienaCounterpartyExtensions.MIFIDReviewDate = sqlDateToHTMLDate(sqlCPEXMIFIDReviewDate.String)
		sienaCounterpartyExtensions.GDPRReviewDate = sqlDateToHTMLDate(sqlCPEXGDPRReviewDate.String)
		sienaCounterpartyExtensions.DelegatedReporting = sqlCPEXDelegatedReporting.String
		sienaCounterpartyExtensions.BOReconcile = sqlCPEXBOReconcile.String
		sienaCounterpartyExtensions.MIFIDReportableDealsAllowed = sqlCPEXMIFIDReportableDealsAllowed.String
		sienaCounterpartyExtensions.SignedInvestmentAgreement = sienaYN(sqlCPEXSignedInvestmentAgreement.String)
		sienaCounterpartyExtensions.AppropriatenessAssessment = sienaYN(sqlCPEXAppropriatenessAssessment.String)
		sienaCounterpartyExtensions.FinancialCounterparty = sqlCPEXFinancialCounterparty.String
		sienaCounterpartyExtensions.Collateralisation = sqlCPEXCollateralisation.String
		sienaCounterpartyExtensions.PortfolioCode = sqlCPEXPortfolioCode.String
		sienaCounterpartyExtensions.ReconciliationLetterFrequency = sqlCPEXReconciliationLetterFrequency.String
		sienaCounterpartyExtensions.DirectDealing = sqlCPEXDirectDealing.String

		sienaCounterpartyExtensionsList = append(sienaCounterpartyExtensionsList, sienaCounterpartyExtensions)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyExtensionsList, sienaCounterpartyExtensions, nil
}