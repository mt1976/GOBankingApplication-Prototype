package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appLSEGiltsDataStoreSQL = "id, 	longName, 	isin, 	tidm, 	sedol, 	issueDate, 	maturityDate, 	couponValue, 	couponType, 	segment, 	sector, 	codeConventionCalculateAccrual, 	minimumDenomination, 	denominationCurrency, 	tradingCurrency, 	type, 	flatYield, 	paymentCouponDate, 	periodOfCoupon, 	exCouponDate, 	dateOfIndexInflation, 	unitOfQuotation, 	_Created, 	_Who, 	_Host, 	_Updated, 	issuer, 	issueAmount, 	runningYield, 	lei, 	cusip"

var sqlLSEGiltsDataStoreId, sqlLSEGiltsDataStoreLongName, sqlLSEGiltsDataStoreIsin, sqlLSEGiltsDataStoreTidm, sqlLSEGiltsDataStoreSedol, sqlLSEGiltsDataStoreIssueDate, sqlLSEGiltsDataStoreMaturityDate, sqlLSEGiltsDataStoreCouponValue, sqlLSEGiltsDataStoreCouponType, sqlLSEGiltsDataStoreSegment, sqlLSEGiltsDataStoreSector, sqlLSEGiltsDataStoreCodeConventionCalculateAccrual, sqlLSEGiltsDataStoreMinimumDenomination, sqlLSEGiltsDataStoreDenominationCurrency, sqlLSEGiltsDataStoreTradingCurrency, sqlLSEGiltsDataStoreType, sqlLSEGiltsDataStoreFlatYield, sqlLSEGiltsDataStorePaymentCouponDate, sqlLSEGiltsDataStorePeriodOfCoupon, sqlLSEGiltsDataStoreExCouponDate, sqlLSEGiltsDataStoreDateOfIndexInflation, sqlLSEGiltsDataStoreUnitOfQuotation, sqlLSEGiltsDataStoreSYSCreated, sqlLSEGiltsDataStoreSYSWho, sqlLSEGiltsDataStoreSYSHost, sqlLSEGiltsDataStoreSYSUpdated, sqlLSEGiltsDataStoreIssuer, sqlLSEGiltsDataStoreIssueAmount, sqlLSEGiltsDataStoreRunningYield, sqlLSEGiltsDataStoreLei, sqlLSEGiltsDataStoreCusip sql.NullString

var appLSEGiltsDataStoreNoParams = strings.Count(appLSEGiltsDataStoreSQL, ",") + 1
var appLSEGiltsDataStoreParams = strings.Repeat("'%s',", appLSEGiltsDataStoreNoParams)
var appLSEGiltsDataStoreSQLINSERT = "INSERT INTO %s.lseGiltsDataStore(%s) VALUES(" + strings.TrimRight(appLSEGiltsDataStoreParams, ",") + ");"
var appLSEGiltsDataStoreSQLDELETE = "DELETE FROM %s.lseGiltsDataStore WHERE id='%s';"
var appLSEGiltsDataStoreSQLSELECT = "SELECT %s FROM %s.lseGiltsDataStore;"
var appLSEGiltsDataStoreSQLGET = "SELECT %s FROM %s.lseGiltsDataStore WHERE id='%s';"

//appLSEGiltsDataStorePage is cheese
type appLSEGiltsDataStoreListPage struct {
	UserMenu               []AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	LSEGiltsDataStoreCount int
	LSEGiltsDataStoreList  []AppLSEGiltsDataStoreItem
}

//appLSEGiltsDataStorePage is cheese
type appLSEGiltsDataStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Part Below
	Id                             string
	LongName                       string
	Isin                           string
	Tidm                           string
	Sedol                          string
	IssueDate                      string
	MaturityDate                   string
	CouponValue                    string
	CouponType                     string
	Segment                        string
	Sector                         string
	CodeConventionCalculateAccrual string
	MinimumDenomination            string
	DenominationCurrency           string
	TradingCurrency                string
	Type                           string
	FlatYield                      string
	PaymentCouponDate              string
	PeriodOfCoupon                 string
	ExCouponDate                   string
	DateOfIndexInflation           string
	UnitOfQuotation                string
	SYSCreated                     string
	SYSWho                         string
	SYSHost                        string
	SYSUpdated                     string
	Issuer                         string
	IssueAmount                    string
	RunningYield                   string
	Lei                            string
	Cusip                          string
}

//AppLSEGiltsDataStoreItem is cheese
type AppLSEGiltsDataStoreItem struct {
	Id                             string
	LongName                       string
	Isin                           string
	Tidm                           string
	Sedol                          string
	IssueDate                      string
	MaturityDate                   string
	CouponValue                    string
	CouponType                     string
	Segment                        string
	Sector                         string
	CodeConventionCalculateAccrual string
	MinimumDenomination            string
	DenominationCurrency           string
	TradingCurrency                string
	Type                           string
	FlatYield                      string
	PaymentCouponDate              string
	PeriodOfCoupon                 string
	ExCouponDate                   string
	DateOfIndexInflation           string
	UnitOfQuotation                string
	SYSCreated                     string
	SYSWho                         string
	SYSHost                        string
	SYSUpdated                     string
	Issuer                         string
	IssueAmount                    string
	RunningYield                   string
	Lei                            string
	Cusip                          string
}

func ListLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LSEGiltsDataStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []AppLSEGiltsDataStoreItem

	noItems, returnList, _ := GetLSEGiltsDataStoreList(globals.ApplicationDB)

	pageLSEGiltsDataStoreList := appLSEGiltsDataStoreListPage{
		UserMenu:               GetUserMenu(r),
		UserRole:               GetUserRole(r),
		UserNavi:               "NOT USED",
		Title:                  globals.ApplicationProperties["appname"],
		PageTitle:              globals.ApplicationProperties["appname"] + " - " + "LSE Gilts",
		LSEGiltsDataStoreCount: noItems,
		LSEGiltsDataStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLSEGiltsDataStoreList)

}

func ViewLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LSEGiltsDataStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LSEGiltsDataStore")
	_, returnRecord, _ := GetLSEGiltsDataStoreByID(searchID)

	pageCredentialStoreList := appLSEGiltsDataStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "LSE Gilts - View",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:                             returnRecord.Id,
		LongName:                       returnRecord.LongName,
		Isin:                           returnRecord.Isin,
		Tidm:                           returnRecord.Tidm,
		Sedol:                          returnRecord.Sedol,
		IssueDate:                      returnRecord.IssueDate,
		MaturityDate:                   returnRecord.MaturityDate,
		CouponValue:                    returnRecord.CouponValue,
		CouponType:                     returnRecord.CouponType,
		Segment:                        returnRecord.Segment,
		Sector:                         returnRecord.Sector,
		CodeConventionCalculateAccrual: returnRecord.CodeConventionCalculateAccrual,
		MinimumDenomination:            returnRecord.MinimumDenomination,
		DenominationCurrency:           returnRecord.DenominationCurrency,
		TradingCurrency:                returnRecord.TradingCurrency,
		Type:                           returnRecord.Type,
		FlatYield:                      returnRecord.FlatYield,
		PaymentCouponDate:              returnRecord.PaymentCouponDate,
		PeriodOfCoupon:                 returnRecord.PeriodOfCoupon,
		ExCouponDate:                   returnRecord.ExCouponDate,
		DateOfIndexInflation:           returnRecord.DateOfIndexInflation,
		UnitOfQuotation:                returnRecord.UnitOfQuotation,
		SYSCreated:                     returnRecord.SYSCreated,
		SYSWho:                         returnRecord.SYSWho,
		SYSHost:                        returnRecord.SYSHost,
		SYSUpdated:                     returnRecord.SYSUpdated,
		Issuer:                         returnRecord.Issuer,
		IssueAmount:                    returnRecord.IssueAmount,
		RunningYield:                   returnRecord.RunningYield,
		Lei:                            returnRecord.Lei,
		Cusip:                          returnRecord.Cusip,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func EditLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LSEGiltsDataStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LSEGiltsDataStore")
	_, returnRecord, _ := GetLSEGiltsDataStoreByID(searchID)

	pageCredentialStoreList := appLSEGiltsDataStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "LSE Gilts - Edit",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:                             returnRecord.Id,
		LongName:                       returnRecord.LongName,
		Isin:                           returnRecord.Isin,
		Tidm:                           returnRecord.Tidm,
		Sedol:                          returnRecord.Sedol,
		IssueDate:                      returnRecord.IssueDate,
		MaturityDate:                   returnRecord.MaturityDate,
		CouponValue:                    returnRecord.CouponValue,
		CouponType:                     returnRecord.CouponType,
		Segment:                        returnRecord.Segment,
		Sector:                         returnRecord.Sector,
		CodeConventionCalculateAccrual: returnRecord.CodeConventionCalculateAccrual,
		MinimumDenomination:            returnRecord.MinimumDenomination,
		DenominationCurrency:           returnRecord.DenominationCurrency,
		TradingCurrency:                returnRecord.TradingCurrency,
		Type:                           returnRecord.Type,
		FlatYield:                      returnRecord.FlatYield,
		PaymentCouponDate:              returnRecord.PaymentCouponDate,
		PeriodOfCoupon:                 returnRecord.PeriodOfCoupon,
		ExCouponDate:                   returnRecord.ExCouponDate,
		DateOfIndexInflation:           returnRecord.DateOfIndexInflation,
		UnitOfQuotation:                returnRecord.UnitOfQuotation,
		SYSCreated:                     returnRecord.SYSCreated,
		SYSWho:                         returnRecord.SYSWho,
		SYSHost:                        returnRecord.SYSHost,
		SYSUpdated:                     returnRecord.SYSUpdated,
		Issuer:                         returnRecord.Issuer,
		IssueAmount:                    returnRecord.IssueAmount,
		RunningYield:                   returnRecord.RunningYield,
		Lei:                            returnRecord.Lei,
		Cusip:                          returnRecord.Cusip,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func SaveLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s AppLSEGiltsDataStoreItem

	s.Id = r.FormValue("Id")
	s.LongName = r.FormValue("LongName")
	s.Isin = r.FormValue("Isin")
	s.Tidm = r.FormValue("Tidm")
	s.Sedol = r.FormValue("Sedol")
	s.IssueDate = r.FormValue("IssueDate")
	s.MaturityDate = r.FormValue("MaturityDate")
	s.CouponValue = r.FormValue("CouponValue")
	s.CouponType = r.FormValue("CouponType")
	s.Segment = r.FormValue("Segment")
	s.Sector = r.FormValue("Sector")
	s.CodeConventionCalculateAccrual = r.FormValue("CodeConventionCalculateAccrual")
	s.MinimumDenomination = r.FormValue("MinimumDenomination")
	s.DenominationCurrency = r.FormValue("DenominationCurrency")
	s.TradingCurrency = r.FormValue("TradingCurrency")
	s.Type = r.FormValue("Type")
	s.FlatYield = r.FormValue("FlatYield")
	s.PaymentCouponDate = r.FormValue("PaymentCouponDate")
	s.PeriodOfCoupon = r.FormValue("PeriodOfCoupon")
	s.ExCouponDate = r.FormValue("ExCouponDate")
	s.DateOfIndexInflation = r.FormValue("DateOfIndexInflation")
	s.UnitOfQuotation = r.FormValue("UnitOfQuotation")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	s.Issuer = r.FormValue("Issuer")
	s.IssueAmount = r.FormValue("IssueAmount")
	s.RunningYield = r.FormValue("RunningYield")
	s.Lei = r.FormValue("Lei")
	s.Cusip = r.FormValue("Cusip")
	//log.Println("save", s)

	putLSEGiltsDataStore(s, r)

	ListLSEGiltsDataStoreHandler(w, r)

}

func DeleteLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LSEGiltsDataStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	deleteLSEGiltsDataStore(searchID)
	ListLSEGiltsDataStoreHandler(w, r)

}

func BanLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LSEGiltsDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Ban", searchID)
	banLSEGiltsDataStore(searchID, r)
	ListLSEGiltsDataStoreHandler(w, r)
}

func ActivateLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LSEGiltsDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Activate", searchID)
	activateLSEGiltsDataStore(searchID, r)
	ListLSEGiltsDataStoreHandler(w, r)

}

func NewLSEGiltsDataStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LSEGiltsDataStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageCredentialStoreList := appLSEGiltsDataStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "LSE Gilts - New",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

// getLSEGiltsDataStoreList read all employees
func GetLSEGiltsDataStoreList(unused *sql.DB) (int, []AppLSEGiltsDataStoreItem, error) {
	tsql := fmt.Sprintf(appLSEGiltsDataStoreSQLSELECT, appLSEGiltsDataStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appLSEGiltsDataStoreList, _, _ := fetchLSEGiltsDataStoreData(globals.ApplicationDB, tsql)
	return count, appLSEGiltsDataStoreList, nil
}

// getLSEGiltsDataStoreList read all employees
func GetLSEGiltsDataStoreByID(id string) (int, AppLSEGiltsDataStoreItem, error) {
	tsql := fmt.Sprintf(appLSEGiltsDataStoreSQLGET, appLSEGiltsDataStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, AppLSEGiltsDataStoreItem, _ := fetchLSEGiltsDataStoreData(globals.ApplicationDB, tsql)
	return 1, AppLSEGiltsDataStoreItem, nil
}

func putLSEGiltsDataStore(r AppLSEGiltsDataStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	userID := GetUserName(req)
	putLSEGiltsDataStoreUser(r, userID)
}

func PutLSEGiltsDataStoreSystem(r AppLSEGiltsDataStoreItem) {
	//fmt.Println(credentialStore)
	putLSEGiltsDataStoreSystem(r)
}
func putLSEGiltsDataStoreSystem(r AppLSEGiltsDataStoreItem) {
	//fmt.Println(credentialStore)
	putLSEGiltsDataStoreUser(r, "AutoGenerated")
}

func putLSEGiltsDataStoreUser(r AppLSEGiltsDataStoreItem, userID string) {
	//fmt.Println(credentialStore)

	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.Id) == 0 {
		r.Id = getLSEGiltsDataStoreID(r)
	}
	//	currentUserID, _ := user.Current()
	//	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
		// Default in info for a new RECORD
	}

	r.SYSUpdated = createDate

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appLSEGiltsDataStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLSEGiltsDataStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appLSEGiltsDataStoreSQL, r.Id, r.LongName, r.Isin, r.Tidm, r.Sedol, r.IssueDate, r.MaturityDate, r.CouponValue, r.CouponType, r.Segment, r.Sector, r.CodeConventionCalculateAccrual, r.MinimumDenomination, r.DenominationCurrency, r.TradingCurrency, r.Type, r.FlatYield, r.PaymentCouponDate, r.PeriodOfCoupon, r.ExCouponDate, r.DateOfIndexInflation, r.UnitOfQuotation, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Issuer, r.IssueAmount, r.RunningYield, r.Lei, r.Cusip)

	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	//log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panicf("%e", err2)
	}
	_, err := globals.ApplicationDB.Exec(inserttsql)
	if err != nil {
		log.Panicf("%e", err)
	}
}

func deleteLSEGiltsDataStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLSEGiltsDataStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql)
	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//log.Println(fred2, err2)
}

func banLSEGiltsDataStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	_, r, err2 := GetLSEGiltsDataStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putLSEGiltsDataStore(r, req)
}

func activateLSEGiltsDataStore(id string, req *http.Request) {
	fmt.Println("RECORD", id)
	_, r, err2 := GetLSEGiltsDataStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putLSEGiltsDataStore(r, req)
}

// fetchLSEGiltsDataStoreData read all employees
func fetchLSEGiltsDataStoreData(unused *sql.DB, tsql string) (int, []AppLSEGiltsDataStoreItem, AppLSEGiltsDataStoreItem, error) {
	//log.Println(tsql)
	var appLSEGiltsDataStore AppLSEGiltsDataStoreItem
	var appLSEGiltsDataStoreList []AppLSEGiltsDataStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appLSEGiltsDataStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlLSEGiltsDataStoreId, &sqlLSEGiltsDataStoreLongName, &sqlLSEGiltsDataStoreIsin, &sqlLSEGiltsDataStoreTidm, &sqlLSEGiltsDataStoreSedol, &sqlLSEGiltsDataStoreIssueDate, &sqlLSEGiltsDataStoreMaturityDate, &sqlLSEGiltsDataStoreCouponValue, &sqlLSEGiltsDataStoreCouponType, &sqlLSEGiltsDataStoreSegment, &sqlLSEGiltsDataStoreSector, &sqlLSEGiltsDataStoreCodeConventionCalculateAccrual, &sqlLSEGiltsDataStoreMinimumDenomination, &sqlLSEGiltsDataStoreDenominationCurrency, &sqlLSEGiltsDataStoreTradingCurrency, &sqlLSEGiltsDataStoreType, &sqlLSEGiltsDataStoreFlatYield, &sqlLSEGiltsDataStorePaymentCouponDate, &sqlLSEGiltsDataStorePeriodOfCoupon, &sqlLSEGiltsDataStoreExCouponDate, &sqlLSEGiltsDataStoreDateOfIndexInflation, &sqlLSEGiltsDataStoreUnitOfQuotation, &sqlLSEGiltsDataStoreSYSCreated, &sqlLSEGiltsDataStoreSYSWho, &sqlLSEGiltsDataStoreSYSHost, &sqlLSEGiltsDataStoreSYSUpdated, &sqlLSEGiltsDataStoreIssuer, &sqlLSEGiltsDataStoreIssueAmount, &sqlLSEGiltsDataStoreRunningYield, &sqlLSEGiltsDataStoreLei, &sqlLSEGiltsDataStoreCusip)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appLSEGiltsDataStore, err
		}
		// Mapping from SQL to Struct
		appLSEGiltsDataStore.Id = sqlLSEGiltsDataStoreId.String
		appLSEGiltsDataStore.LongName = sqlLSEGiltsDataStoreLongName.String
		appLSEGiltsDataStore.Isin = sqlLSEGiltsDataStoreIsin.String
		appLSEGiltsDataStore.Tidm = sqlLSEGiltsDataStoreTidm.String
		appLSEGiltsDataStore.Sedol = sqlLSEGiltsDataStoreSedol.String
		appLSEGiltsDataStore.IssueDate = sqlLSEGiltsDataStoreIssueDate.String
		appLSEGiltsDataStore.MaturityDate = sqlLSEGiltsDataStoreMaturityDate.String
		appLSEGiltsDataStore.CouponValue = sqlLSEGiltsDataStoreCouponValue.String
		appLSEGiltsDataStore.CouponType = sqlLSEGiltsDataStoreCouponType.String
		appLSEGiltsDataStore.Segment = sqlLSEGiltsDataStoreSegment.String
		appLSEGiltsDataStore.Sector = sqlLSEGiltsDataStoreSector.String
		appLSEGiltsDataStore.CodeConventionCalculateAccrual = sqlLSEGiltsDataStoreCodeConventionCalculateAccrual.String
		appLSEGiltsDataStore.MinimumDenomination = sqlLSEGiltsDataStoreMinimumDenomination.String
		appLSEGiltsDataStore.DenominationCurrency = sqlLSEGiltsDataStoreDenominationCurrency.String
		appLSEGiltsDataStore.TradingCurrency = sqlLSEGiltsDataStoreTradingCurrency.String
		appLSEGiltsDataStore.Type = sqlLSEGiltsDataStoreType.String
		appLSEGiltsDataStore.FlatYield = sqlLSEGiltsDataStoreFlatYield.String
		appLSEGiltsDataStore.PaymentCouponDate = sqlLSEGiltsDataStorePaymentCouponDate.String
		appLSEGiltsDataStore.PeriodOfCoupon = sqlLSEGiltsDataStorePeriodOfCoupon.String
		appLSEGiltsDataStore.ExCouponDate = sqlLSEGiltsDataStoreExCouponDate.String
		appLSEGiltsDataStore.DateOfIndexInflation = sqlLSEGiltsDataStoreDateOfIndexInflation.String
		appLSEGiltsDataStore.UnitOfQuotation = sqlLSEGiltsDataStoreUnitOfQuotation.String
		appLSEGiltsDataStore.SYSCreated = sqlLSEGiltsDataStoreSYSCreated.String
		appLSEGiltsDataStore.SYSWho = sqlLSEGiltsDataStoreSYSWho.String
		appLSEGiltsDataStore.SYSHost = sqlLSEGiltsDataStoreSYSHost.String
		appLSEGiltsDataStore.SYSUpdated = sqlLSEGiltsDataStoreSYSUpdated.String
		appLSEGiltsDataStore.Issuer = sqlLSEGiltsDataStoreIssuer.String
		appLSEGiltsDataStore.IssueAmount = sqlLSEGiltsDataStoreIssueAmount.String
		appLSEGiltsDataStore.RunningYield = sqlLSEGiltsDataStoreRunningYield.String
		appLSEGiltsDataStore.Lei = sqlLSEGiltsDataStoreLei.String
		appLSEGiltsDataStore.Cusip = sqlLSEGiltsDataStoreCusip.String
		// no change below
		appLSEGiltsDataStoreList = append(appLSEGiltsDataStoreList, appLSEGiltsDataStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appLSEGiltsDataStoreList, appLSEGiltsDataStore)
	return count, appLSEGiltsDataStoreList, appLSEGiltsDataStore, nil
}

func newLSEGiltsDataStoreID() string {
	id := uuid.New().String()
	return id
}

// Returns a valid lseGiltsDataStore store ID
func getLSEGiltsDataStoreID(r AppLSEGiltsDataStoreItem) string {
	var lseGiltsDataStoreID string
	lseGiltsDataStoreID = r.Isin
	return lseGiltsDataStoreID
}
