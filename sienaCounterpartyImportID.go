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

//sienaCounterpartyImportIDPage is cheese
type sienaCounterpartyImportIDListPage struct {
	Title                          string
	PageTitle                      string
	SienaCounterpartyImportIDCount int
	SienaCounterpartyImportIDList  []sienaCounterpartyImportIDItem
}

//sienaCounterpartyImportIDPage is cheese
type sienaCounterpartyImportIDPage struct {
	Title       string
	PageTitle   string
	ID          string
	KeyImportID string
	Firm        string
	Centre      string
	FirmName    string
	CentreName  string
	KeyOriginID string
	Action      string
}

//sienaCounterpartyImportIDItem is cheese
type sienaCounterpartyImportIDItem struct {
	Code        string
	Action      string
	KeyImportID string
	Firm        string
	Centre      string
	FirmName    string
	CentreName  string
	KeyOriginID string
}

func listSienaCounterpartyImportIDHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCounterpartyImportID"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyImportIDItem
	noItems, returnList, _ := getSienaCounterpartyImportIDList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCounterpartyImportIDList := sienaCounterpartyImportIDListPage{
		Title:                          wctProperties["appname"],
		PageTitle:                      "List Siena CounterpartyImportIDs",
		SienaCounterpartyImportIDCount: noItems,
		SienaCounterpartyImportIDList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyImportIDList)

}

func viewSienaCounterpartyImportIDHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCounterpartyImportID"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyImportIDItem
	idImport := getURLparam(r, "II")
	idOrigin := getURLparam(r, "IO")
	noItems, returnRecord, _ := getSienaCounterpartyImportID(thisConnection, idImport, idOrigin)
	fmt.Println("NoSienaItems", noItems, idImport, idOrigin)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCounterpartyImportIDList := sienaCounterpartyImportIDPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena CounterpartyImportID",
		ID:          returnRecord.Code,
		KeyImportID: returnRecord.KeyImportID,
		Firm:        returnRecord.Firm,
		Centre:      returnRecord.Centre,
		FirmName:    returnRecord.FirmName,
		CentreName:  returnRecord.CentreName,
		KeyOriginID: returnRecord.KeyOriginID,
		Action:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyImportIDList)

}

func editSienaCounterpartyImportIDHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCounterpartyImportID"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCounterpartyImportIDItem
	idImport := getURLparam(r, "II")
	idOrigin := getURLparam(r, "IO")
	noItems, returnRecord, _ := getSienaCounterpartyImportID(thisConnection, idImport, idOrigin)
	fmt.Println("NoSienaItems", noItems, idImport, idOrigin)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	//fmt.Println(displayList)

	pageSienaCounterpartyImportIDList := sienaCounterpartyImportIDPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena CounterpartyImportID",
		ID:          returnRecord.Code,
		KeyImportID: returnRecord.KeyImportID,
		Firm:        returnRecord.Firm,
		Centre:      returnRecord.Centre,
		FirmName:    returnRecord.FirmName,
		CentreName:  returnRecord.CentreName,
		KeyOriginID: returnRecord.KeyOriginID,
		Action:      "",
	}
	fmt.Println(pageSienaCounterpartyImportIDList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyImportIDList)

}

func saveSienaCounterpartyImportIDHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaCounterpartyImportIDItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.KeyImportID = r.FormValue("KeyImportID")
	item.Firm = r.FormValue("Firm")
	item.Centre = r.FormValue("Centre")
	item.FirmName = r.FormValue("FirmName")
	item.CentreName = r.FormValue("CentreName")
	item.KeyOriginID = r.FormValue("KeyOriginID")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.Code

	sFldName.Name = "fullName"
	sFldName.Text = item.Firm

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Centre

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
	sienaTable.Name = "CounterpartyImportID"
	sienaTable.Classname = "com.eurobase.siena.data.CounterpartyImportIDs.CounterpartyImportID"
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

	listSienaCounterpartyImportIDHandler(w, r)

}

func newSienaCounterpartyImportIDHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCounterpartyImportID"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items

	pageSienaCounterpartyImportIDList := sienaCounterpartyImportIDPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena CounterpartyImportID",
		ID:          "NEW",
		KeyImportID: "externalDealImporter",
		Firm:        "",
		Centre:      "",
		FirmName:    "",
		CentreName:  "",
		KeyOriginID: "",

		Action: "NEW",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCounterpartyImportIDList)

}

// getSienaCounterpartyImportIDList read all employees
func getSienaCounterpartyImportIDList(db *sql.DB) (int, []sienaCounterpartyImportIDItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCounterpartyImportIDList []sienaCounterpartyImportIDItem
	var sienaCounterpartyImportID sienaCounterpartyImportIDItem
	tsql := fmt.Sprintf("SELECT KeyImportID, 	Firm, 	Centre, 	FirmName, 	CentreName, 	KeyOriginID FROM %s.sienaCounterpartyImportID;", mssqlConfig["schema"])
	//	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var KeyImportID, Firm, Centre, FirmName, CentreName, KeyOriginID sql.NullString
		err := rows.Scan(&KeyImportID, &Firm, &Centre, &FirmName, &CentreName, &KeyOriginID)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCounterpartyImportID.KeyImportID = KeyImportID.String
		sienaCounterpartyImportID.Firm = Firm.String
		sienaCounterpartyImportID.Centre = Centre.String
		sienaCounterpartyImportID.FirmName = FirmName.String
		sienaCounterpartyImportID.CentreName = CentreName.String
		sienaCounterpartyImportID.KeyOriginID = KeyOriginID.String
		sienaCounterpartyImportIDList = append(sienaCounterpartyImportIDList, sienaCounterpartyImportID)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyImportIDList, nil
}

// getSienaCounterpartyImportIDList read all employees
func getSienaCounterpartyImportID(db *sql.DB, idImportID string, idOriginID string) (int, sienaCounterpartyImportIDItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCounterpartyImportID sienaCounterpartyImportIDItem
	tsql := fmt.Sprintf("SELECT KeyImportID, 	Firm, 	Centre, 	FirmName, 	CentreName, 	KeyOriginID FROM %s.sienaCounterpartyImportID WHERE KeyImportID='%s' AND KeyOriginID='%s';", mssqlConfig["schema"], idImportID, idOriginID)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCounterpartyImportID, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var KeyImportID, Firm, Centre, FirmName, CentreName, KeyOriginID sql.NullString
		err := rows.Scan(&KeyImportID, &Firm, &Centre, &FirmName, &CentreName, &KeyOriginID)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCounterpartyImportID, err
		}
		sienaCounterpartyImportID.KeyImportID = KeyImportID.String
		sienaCounterpartyImportID.Firm = Firm.String
		sienaCounterpartyImportID.Centre = Centre.String
		sienaCounterpartyImportID.FirmName = FirmName.String
		sienaCounterpartyImportID.CentreName = CentreName.String
		sienaCounterpartyImportID.KeyOriginID = KeyOriginID.String

		count++
	}
	return 1, sienaCounterpartyImportID, nil
}

// getSienaCounterpartyImportIDList read all employees
func putSienaCounterpartyImportID(db *sql.DB, updateItem sienaCounterpartyImportIDItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// getSienaCounterpartyImportIDList read all employees
func getSienaCounterpartyImportIDListByCounterparty(db *sql.DB, idFirm string, idCentre string) (int, []sienaCounterpartyImportIDItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCounterpartyImportIDList []sienaCounterpartyImportIDItem
	var sienaCounterpartyImportID sienaCounterpartyImportIDItem
	tsql := fmt.Sprintf("SELECT KeyImportID, 	Firm, 	Centre, 	FirmName, 	CentreName, 	KeyOriginID FROM %s.sienaCounterpartyImportID WHERE Firm='%s' AND Centre='%s';", mssqlConfig["schema"], idFirm, idCentre)
	//	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var KeyImportID, Firm, Centre, FirmName, CentreName, KeyOriginID sql.NullString
		err := rows.Scan(&KeyImportID, &Firm, &Centre, &FirmName, &CentreName, &KeyOriginID)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCounterpartyImportID.KeyImportID = KeyImportID.String
		sienaCounterpartyImportID.Firm = Firm.String
		sienaCounterpartyImportID.Centre = Centre.String
		sienaCounterpartyImportID.FirmName = FirmName.String
		sienaCounterpartyImportID.CentreName = CentreName.String
		sienaCounterpartyImportID.KeyOriginID = KeyOriginID.String
		sienaCounterpartyImportIDList = append(sienaCounterpartyImportIDList, sienaCounterpartyImportID)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyImportIDList, nil
}
