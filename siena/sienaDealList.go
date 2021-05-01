package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var sienaDealListSQL = "SienaReference, 	CustomerSienaView, 	Status, 	ValueDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	Book, 	MandatedUser, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	ISIN, 	UTI, 	BookName, 	Centre, 	Firm, 	DealTypeShortName, 	FullDealType, 	TradeDate, 	DealtCcy, 	DealtAmount, 	AgainstAmount, 	AgainstCcy, 	AllInRate, 	MktRate, 	SettleCcy, 	Direction, 	NpvRate, 	OriginUser, 	PayInstruction, 	ReceiptInstruction, 	NIName, 	CCYPair, 	Instrument, 	PortfolioName, 	RVDate, 	RVMTM, 	CounterBook, 	CounterBookName, 	Party, 	PartyName"

// Defined the Temp Variables used to extract data returned from SQL
var sqlDLSTSienaReference, sqlDLSTCustomerSienaView, sqlDLSTStatus, sqlDLSTValueDate, sqlDLSTMaturityDate, sqlDLSTContractNumber, sqlDLSTExternalReference, sqlDLSTBook, sqlDLSTMandatedUser, sqlDLSTPortfolio, sqlDLSTAgreementId, sqlDLSTBackOfficeRefNo, sqlDLSTISIN, sqlDLSTUTI, sqlDLSTBookName, sqlDLSTCentre, sqlDLSTFirm, sqlDLSTDealTypeShortName, sqlDLSTFullDealType, sqlDLSTTradeDate, sqlDLSTDealtCcy, sqlDLSTDealtAmount, sqlDLSTAgainstAmount, sqlDLSTAgainstCcy, sqlDLSTAllInRate, sqlDLSTMktRate, sqlDLSTSettleCcy, sqlDLSTDirection, sqlDLSTNpvRate, sqlDLSTOriginUser, sqlDLSTPayInstruction, sqlDLSTReceiptInstruction, sqlDLSTNIName, sqlDLSTCCYPair, sqlDLSTInstrument, sqlDLSTPortfolioName, sqlDLSTRVDate, sqlDLSTRVMTM, sqlDLSTCounterBook, sqlDLSTCounterBookName, sqlDLSTParty, sqlDLSTPartyName sql.NullString

//sienaDealListPage is cheese
type sienaDealListListPage struct {
	UserMenu           []application.AppMenuItem
	UserRole           string
	UserNavi           string
	Title              string
	PageTitle          string
	SienaDealListCount int
	SienaDealListList  []sienaDealListItem
}

//sienaDealListPage is cheese
type sienaDealListPage struct {
	UserMenu           []application.AppMenuItem
	UserRole           string
	UserNavi           string
	Title              string
	PageTitle          string
	ID                 string
	Action             string
	SienaReference     string
	CustomerSienaView  string
	Status             string
	ValueDate          string
	MaturityDate       string
	ContractNumber     string
	ExternalReference  string
	Book               string
	MandatedUser       string
	Portfolio          string
	AgreementId        string
	BackOfficeRefNo    string
	ISIN               string
	UTI                string
	BookName           string
	Centre             string
	Firm               string
	DealTypeShortName  string
	FullDealType       string
	TradeDate          string
	DealtCcy           string
	DealtAmount        string
	AgainstAmount      string
	AgainstCcy         string
	AllInRate          string
	MktRate            string
	SettleCcy          string
	Direction          string
	NpvRate            string
	OriginUser         string
	PayInstruction     string
	ReceiptInstruction string
	NIName             string
	CCYPair            string
	Instrument         string
	PortfolioName      string
	RVDate             string
	RVMTM              string
	CounterBook        string
	CounterBookName    string
	Party              string
	PartyName          string
}

//sienaDealListItem is cheese
type sienaDealListItem struct {
	Action             string
	SienaReference     string
	CustomerSienaView  string
	Status             string
	ValueDate          string
	MaturityDate       string
	ContractNumber     string
	ExternalReference  string
	Book               string
	MandatedUser       string
	Portfolio          string
	AgreementId        string
	BackOfficeRefNo    string
	ISIN               string
	UTI                string
	BookName           string
	Centre             string
	Firm               string
	DealTypeShortName  string
	FullDealType       string
	TradeDate          string
	DealtCcy           string
	DealtAmount        string
	AgainstAmount      string
	AgainstCcy         string
	AllInRate          string
	MktRate            string
	SettleCcy          string
	Direction          string
	NpvRate            string
	OriginUser         string
	PayInstruction     string
	ReceiptInstruction string
	NIName             string
	CCYPair            string
	Instrument         string
	PortfolioName      string
	RVDate             string
	RVMTM              string
	CounterBook        string
	CounterBookName    string
	Party              string
	PartyName          string
}

func ListSienaDealListHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaDealList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	var returnList []sienaDealListItem
	noItems, returnList, _ := getSienaDealListList()

	pageSienaDealListList := sienaDealListListPage{
		UserMenu:           application.GetUserMenu(r),
		UserRole:           application.GetUserRole(r),
		UserNavi:           "NOT USED",
		Title:              globals.ApplicationProperties["appname"],
		PageTitle:          "List Siena DealLists",
		SienaDealListCount: noItems,
		SienaDealListList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaDealListList)

}

func ViewSienaDealListHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaDealList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessageAction(inUTL, tmpl, "")

	sienaDealListID := application.GetURLparam(r, "SienaRef")
	noItems, returnRecord, _ := getSienaDealList(sienaDealListID)
	fmt.Println("NoSienaItems", noItems, sienaDealListID, returnRecord.Status)

	pageSienaDealListList := sienaDealListPage{
		UserMenu:           application.GetUserMenu(r),
		UserRole:           application.GetUserRole(r),
		UserNavi:           "NOT USED",
		Title:              globals.ApplicationProperties["appname"],
		PageTitle:          "View Siena DealList",
		Action:             "",
		SienaReference:     returnRecord.SienaReference,
		CustomerSienaView:  returnRecord.CustomerSienaView,
		Status:             returnRecord.Status,
		ValueDate:          returnRecord.ValueDate,
		MaturityDate:       returnRecord.MaturityDate,
		ContractNumber:     returnRecord.ContractNumber,
		ExternalReference:  returnRecord.ExternalReference,
		Book:               returnRecord.Book,
		MandatedUser:       returnRecord.MandatedUser,
		Portfolio:          returnRecord.Portfolio,
		AgreementId:        returnRecord.AgreementId,
		BackOfficeRefNo:    returnRecord.BackOfficeRefNo,
		ISIN:               returnRecord.ISIN,
		UTI:                returnRecord.UTI,
		BookName:           returnRecord.BookName,
		Centre:             returnRecord.Centre,
		Firm:               returnRecord.Firm,
		DealTypeShortName:  returnRecord.DealTypeShortName,
		FullDealType:       returnRecord.FullDealType,
		TradeDate:          returnRecord.TradeDate,
		DealtCcy:           returnRecord.DealtCcy,
		DealtAmount:        returnRecord.DealtAmount,
		AgainstAmount:      returnRecord.AgainstAmount,
		AgainstCcy:         returnRecord.AgainstCcy,
		AllInRate:          returnRecord.AllInRate,
		MktRate:            returnRecord.MktRate,
		SettleCcy:          returnRecord.SettleCcy,
		Direction:          returnRecord.Direction,
		NpvRate:            returnRecord.NpvRate,
		OriginUser:         returnRecord.OriginUser,
		PayInstruction:     returnRecord.PayInstruction,
		ReceiptInstruction: returnRecord.ReceiptInstruction,
		NIName:             returnRecord.NIName,
		CCYPair:            returnRecord.CCYPair,
		Instrument:         returnRecord.Instrument,
		PortfolioName:      returnRecord.PortfolioName,
		RVDate:             returnRecord.RVDate,
		RVMTM:              returnRecord.RVMTM,
		CounterBook:        returnRecord.CounterBook,
		CounterBookName:    returnRecord.CounterBookName,
		Party:              returnRecord.Party,
		PartyName:          returnRecord.PartyName,
	}
	fmt.Println("PAGE", pageSienaDealListList)
	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaDealListList)

}

// getSienaDealListList read all employees
func getSienaDealListList() (int, []sienaDealListItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaDealList;", sienaDealListSQL, globals.SienaPropertiesDB["schema"])
	count, sienaDealListList, _, _ := fetchSienaDealListData(tsql)

	return count, sienaDealListList, nil
}

// getSienaDealListList read all employees
func getSienaDealListListByCounterparty(idFirm string, idCentre string) (int, []sienaDealListItem, error) {
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaDealList WHERE Firm='%s' AND Centre='%s';", sienaDealListSQL, globals.SienaPropertiesDB["schema"], idFirm, idCentre)
	count, sienaDealListList, _, _ := fetchSienaDealListData(tsql)
	return count, sienaDealListList, nil
}

// getSienaDealListList read all employees
func getSienaDealList(sienaDealListID string) (int, sienaDealListItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaDealList WHERE SienaReference='%s';", sienaDealListSQL, globals.SienaPropertiesDB["schema"], sienaDealListID)
	_, _, sienaDealList, _ := fetchSienaDealListData(tsql)
	return 1, sienaDealList, nil
}

// getSienaDealListList read all employees
func putSienaDealList(updateItem sienaDealListItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaDealListData read all employees
func fetchSienaDealListData(tsql string) (int, []sienaDealListItem, sienaDealListItem, error) {

	var sienaDealList sienaDealListItem
	var sienaDealListList []sienaDealListItem

	rows, err := globals.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaDealList, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlDLSTSienaReference, &sqlDLSTCustomerSienaView, &sqlDLSTStatus, &sqlDLSTValueDate, &sqlDLSTMaturityDate, &sqlDLSTContractNumber, &sqlDLSTExternalReference, &sqlDLSTBook, &sqlDLSTMandatedUser, &sqlDLSTPortfolio, &sqlDLSTAgreementId, &sqlDLSTBackOfficeRefNo, &sqlDLSTISIN, &sqlDLSTUTI, &sqlDLSTBookName, &sqlDLSTCentre, &sqlDLSTFirm, &sqlDLSTDealTypeShortName, &sqlDLSTFullDealType, &sqlDLSTTradeDate, &sqlDLSTDealtCcy, &sqlDLSTDealtAmount, &sqlDLSTAgainstAmount, &sqlDLSTAgainstCcy, &sqlDLSTAllInRate, &sqlDLSTMktRate, &sqlDLSTSettleCcy, &sqlDLSTDirection, &sqlDLSTNpvRate, &sqlDLSTOriginUser, &sqlDLSTPayInstruction, &sqlDLSTReceiptInstruction, &sqlDLSTNIName, &sqlDLSTCCYPair, &sqlDLSTInstrument, &sqlDLSTPortfolioName, &sqlDLSTRVDate, &sqlDLSTRVMTM, &sqlDLSTCounterBook, &sqlDLSTCounterBookName, &sqlDLSTParty, &sqlDLSTPartyName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaDealList, err
		}

		sienaDealList.SienaReference = sqlDLSTSienaReference.String
		sienaDealList.CustomerSienaView = sqlDLSTCustomerSienaView.String
		sienaDealList.Status = sqlDLSTStatus.String
		sienaDealList.ValueDate = application.SqlDateToHTMLDate(sqlDLSTValueDate.String)
		sienaDealList.MaturityDate = application.SqlDateToHTMLDate(sqlDLSTMaturityDate.String)
		sienaDealList.ContractNumber = sqlDLSTContractNumber.String
		sienaDealList.ExternalReference = sqlDLSTExternalReference.String
		sienaDealList.Book = sqlDLSTBook.String
		sienaDealList.MandatedUser = sqlDLSTMandatedUser.String
		sienaDealList.Portfolio = sqlDLSTPortfolio.String
		sienaDealList.AgreementId = sqlDLSTAgreementId.String
		sienaDealList.BackOfficeRefNo = sqlDLSTBackOfficeRefNo.String
		sienaDealList.ISIN = sqlDLSTISIN.String
		sienaDealList.UTI = sqlDLSTUTI.String
		sienaDealList.BookName = sqlDLSTBookName.String
		sienaDealList.Centre = sqlDLSTCentre.String
		sienaDealList.Firm = sqlDLSTFirm.String
		sienaDealList.DealTypeShortName = sqlDLSTDealTypeShortName.String
		sienaDealList.FullDealType = sqlDLSTFullDealType.String
		sienaDealList.TradeDate = application.SqlDateToHTMLDate(sqlDLSTTradeDate.String)
		sienaDealList.DealtCcy = sqlDLSTDealtCcy.String
		sienaDealList.DealtAmount = application.FormatCurrency(sqlDLSTDealtAmount.String, sqlDLSTDealtCcy.String)
		sienaDealList.AgainstCcy = sqlDLSTAgainstCcy.String
		sienaDealList.AgainstAmount = application.FormatCurrency(sqlDLSTAgainstAmount.String, sqlDLSTAgainstCcy.String)
		sienaDealList.AllInRate = sqlDLSTAllInRate.String
		sienaDealList.MktRate = sqlDLSTMktRate.String
		sienaDealList.SettleCcy = sqlDLSTSettleCcy.String
		sienaDealList.Direction = sqlDLSTDirection.String
		sienaDealList.NpvRate = sqlDLSTNpvRate.String
		sienaDealList.OriginUser = sqlDLSTOriginUser.String
		sienaDealList.PayInstruction = sqlDLSTPayInstruction.String
		sienaDealList.ReceiptInstruction = sqlDLSTReceiptInstruction.String
		sienaDealList.NIName = sqlDLSTNIName.String
		sienaDealList.CCYPair = sqlDLSTCCYPair.String
		sienaDealList.Instrument = sqlDLSTInstrument.String
		sienaDealList.PortfolioName = sqlDLSTPortfolioName.String
		sienaDealList.RVDate = application.SqlDateToHTMLDate(sqlDLSTRVDate.String)
		sienaDealList.RVMTM = sqlDLSTRVMTM.String
		sienaDealList.CounterBook = sqlDLSTCounterBook.String
		sienaDealList.CounterBookName = sqlDLSTCounterBookName.String
		sienaDealList.Party = sqlDLSTParty.String
		sienaDealList.PartyName = sqlDLSTPartyName.String

		if sienaDealList.Firm != "" {
			sienaDealList.PartyName = fmt.Sprintf("%s%s%s", sienaDealList.Firm, globals.SIENACPTYSEP, sienaDealList.Centre)
		}

		if sienaDealList.Instrument == "" {
			sienaDealList.Instrument = sienaDealList.DealtCcy + " " + sienaDealList.FullDealType + " " + sienaDealList.Direction
		}

		sienaDealListList = append(sienaDealListList, sienaDealList)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaDealListList, sienaDealList, nil
}
