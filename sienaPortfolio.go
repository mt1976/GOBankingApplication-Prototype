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

var sienaPortfolioSQL = "Code, 	Name"
var sqlPRTCode, sqlPRTName sql.NullString

//sienaPortfolioPage is cheese
type sienaPortfolioListPage struct {
	UserMenu            []AppMenuItem
	UserRole            string
	UserNavi            string
	Title               string
	PageTitle           string
	SienaPortfolioCount int
	SienaPortfolioList  []sienaPortfolioItem
}

//sienaPortfolioPage is cheese
type sienaPortfolioPage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
}

//sienaPortfolioItem is cheese
type sienaPortfolioItem struct {
	Code   string
	Name   string
	Action string
}

func listSienaPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaPortfolioItem
	noItems, returnList, _ := getSienaPortfolioList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaPortfolioList := sienaPortfolioListPage{
		UserMenu:            getappMenuData(gUserRole),
		UserRole:            gUserRole,
		UserNavi:            gUserNavi,
		Title:               wctProperties["appname"],
		PageTitle:           "List Siena Portfolios",
		SienaPortfolioCount: noItems,
		SienaPortfolioList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaPortfolioList)

}

func viewSienaPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaPortfolioItem
	searchID := getURLparam(r, "SienaPortfolio")
	noItems, returnRecord, _ := getSienaPortfolio(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaPortfolioList := sienaPortfolioPage{
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Portfolio",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaPortfolioList)

}

func editSienaPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaPortfolioItem
	searchID := getURLparam(r, "SienaPortfolio")
	noItems, returnRecord, _ := getSienaPortfolio(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaPortfolioList := sienaPortfolioPage{
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Portfolio",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaPortfolioList)

}

func saveSienaPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaPortfolioItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("name")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	// POPULATE THE XML FIELDS
	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "Description1"
	sFldName.Text = item.Name
	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Portfolio"
	sienaTable.Classname = "com.eurobase.siena.data.portfolio.Portfolio"
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

	listSienaPortfolioHandler(w, r)

}

func newSienaPortfolioHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaPortfolioList := sienaPortfolioPage{
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Portfolio",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaPortfolioList)

}

// getSienaPortfolioList read all employees
func getSienaPortfolioList(db *sql.DB) (int, []sienaPortfolioItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaPortfolio;", sienaPortfolioSQL, mssqlConfig["schema"])
	count, sienaPortfolioList, _, _ := fetchSienaPortfolioData(db, tsql)
	return count, sienaPortfolioList, nil
}

// getSienaPortfolioList read all employees
func getSienaPortfolio(db *sql.DB, id string) (int, sienaPortfolioItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaPortfolio WHERE Code='%s';", sienaPortfolioSQL, mssqlConfig["schema"], id)
	_, _, sienaPortfolio, _ := fetchSienaPortfolioData(db, tsql)
	return 1, sienaPortfolio, nil
}

// getSienaPortfolioList read all employees
func putSienaPortfolio(db *sql.DB, updateItem sienaPortfolioItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// getSienaPortfolioList read all employees
func fetchSienaPortfolioData(db *sql.DB, tsql string) (int, []sienaPortfolioItem, sienaPortfolioItem, error) {
	log.Println("QUERY", tsql)
	var sienaPortfolioList []sienaPortfolioItem
	var sienaPortfolio sienaPortfolioItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaPortfolio, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlPRTCode, &sqlPRTName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaPortfolio, err
		}
		sienaPortfolio.Code = sqlPRTCode.String
		sienaPortfolio.Name = sqlPRTName.String
		sienaPortfolioList = append(sienaPortfolioList, sienaPortfolio)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, sienaPortfolioList, sienaPortfolio)
	return count, sienaPortfolioList, sienaPortfolio, nil
}
