package siena

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaFirmSQL = "FirmName, 	FullName, 	Country, 	Sector, 	SectorName, 	CountryName"
var sqlFRMFirmName, sqlFRMFullName, sqlFRMCountry, sqlFRMSector, sqlFRMSectorName, sqlFRMCountryName sql.NullString

//sienaFirmPage is cheese
type sienaFirmListPage struct {
	UserMenu       []application.AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	SienaFirmCount int
	SienaFirmList  []sienaFirmItem
}

//sienaFirmPage is cheese
type sienaFirmPage struct {
	UserMenu    []application.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
	CountryList []sienaCountryItem
	SectorList  []sienaSectorItem
}

//sienaFirmItem is cheese
type sienaFirmItem struct {
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
}

func Firm_MUX(mux http.ServeMux) {
	log.Println("MUX Siena Firm")
	mux.HandleFunc("/listSienaFirm/", ListSienaFirmHandler)
	mux.HandleFunc("/viewSienaFirm/", ViewSienaFirmHandler)
	mux.HandleFunc("/editSienaFirm/", EditSienaFirmHandler)
	mux.HandleFunc("/saveSienaFirm/", SaveSienaFirmHandler)
	mux.HandleFunc("/newSienaFirm/", NewSienaFirmHandler)
}

func ListSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	var returnList []sienaFirmItem
	noItems, returnList, _ := getSienaFirmList()

	pageSienaFirmList := sienaFirmListPage{
		UserMenu:       application.GetUserMenu(r),
		UserRole:       application.GetUserRole(r),
		UserNavi:       "NOT USED",
		Title:          globals.ApplicationProperties["appname"],
		PageTitle:      globals.ApplicationProperties["appname"] + " - " + "Firms",
		SienaFirmCount: noItems,
		SienaFirmList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func ViewSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(searchID)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       globals.ApplicationProperties["appname"],
		PageTitle:   globals.ApplicationProperties["appname"] + " - " + "Firm - View",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func EditSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(searchID)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	_, sectorList, _ := getSienaSectorList()

	//fmt.Println(displayList)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       globals.ApplicationProperties["appname"],
		PageTitle:   globals.ApplicationProperties["appname"] + " - " + "Firm - Edit",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
		CountryList: countryList,
		SectorList:  sectorList,
	}
	//fmt.Println(pageSienaFirmList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func SaveSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessageAction(inUTL, "Save", "")

	var item sienaFirmItem

	item.FirmName = r.FormValue("firmName")
	if len(item.FirmName) == 0 {
		item.FirmName = r.FormValue("ID")
	}
	item.FullName = r.FormValue("fullName")
	item.Country = r.FormValue("country")

	item.Sector = r.FormValue("sector")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldFirmName sienaKEYFIELD
	var sFldFullName sienaFIELD
	var sFldCountry sienaFIELD
	var sFldSector sienaFIELD

	// POPULATE THE XML FIELDS
	sFldFirmName.Name = "firmName"
	sFldFirmName.Text = item.FirmName

	sFldFullName.Name = "fullName"
	sFldFullName.Text = item.FullName

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	sFldSector.Name = "sector"
	sFldSector.Text = item.Sector

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldFirmName)
	sienaFields = append(sienaFields, sFldFullName)
	sienaFields = append(sienaFields, sFldCountry)
	sienaFields = append(sienaFields, sFldSector)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Firm"
	sienaTable.Classname = "com.eurobase.siena.data.firms.Firm"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := globals.SienaProperties["static_in"]
	fileID := uuid.New()
	fileName := staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}
	xmlFile.WriteString(globals.SienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(sienaXMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}

	ListSienaFirmHandler(w, r)

}

func NewSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	_, sectorList, _ := getSienaSectorList()

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       globals.ApplicationProperties["appname"],
		PageTitle:   globals.ApplicationProperties["appname"] + " - " + "Firm - New",
		ID:          "NEW",
		FirmName:    "",
		FullName:    "",
		Country:     "",
		Sector:      "",
		Action:      "NEW",
		CountryList: countryList,
		SectorList:  sectorList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

// getSienaFirmList read all employees
func getSienaFirmList() (int, []sienaFirmItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm;", sienaFirmSQL, globals.SienaPropertiesDB["schema"])
	count, sienaFirmList, _, _ := fetchSienaFirmData(tsql)

	return count, sienaFirmList, nil
}

// getSienaFirmList read all employees
func getSienaFirm(id string) (int, sienaFirmItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm WHERE FirmName='%s';", sienaFirmSQL, globals.SienaPropertiesDB["schema"], id)
	_, _, sienaFirm, _ := fetchSienaFirmData(tsql)

	return 1, sienaFirm, nil
}

// getSienaFirmList read all employees
func putSienaFirm(updateItem sienaFirmItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaFirmData read all employees
func fetchSienaFirmData(tsql string) (int, []sienaFirmItem, sienaFirmItem, error) {

	var sienaFirm sienaFirmItem
	var sienaFirmList []sienaFirmItem

	rows, err := globals.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaFirm, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlFRMFirmName, &sqlFRMFullName, &sqlFRMCountry, &sqlFRMSector, &sqlFRMSectorName, &sqlFRMCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaFirm, err
		}

		sienaFirm.FirmName = sqlFRMFirmName.String
		sienaFirm.FullName = sqlFRMFullName.String
		sienaFirm.Country = sqlFRMCountry.String
		sienaFirm.Sector = sqlFRMSector.String
		sienaFirm.SectorName = sqlFRMSectorName.String
		sienaFirm.CountryName = sqlFRMCountryName.String

		sienaFirmList = append(sienaFirmList, sienaFirm)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaFirmList, sienaFirm, nil
}
