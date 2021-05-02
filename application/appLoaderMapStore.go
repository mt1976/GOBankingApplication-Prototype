package application

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var appLoaderMapStoreSQL = "id, 	name, 	position, 	loader, 	_created, 	_who, 	_host, 	_updated"

var sqlLoaderMapStoreId, sqlLoaderMapStoreName, sqlLoaderMapStorePosition, sqlLoaderMapStoreLoader, sqlLoaderMapStoreSYSCreated, sqlLoaderMapStoreSYSWho, sqlLoaderMapStoreSYSHost, sqlLoaderMapStoreSYSUpdated sql.NullString

var appLoaderMapStoreSQLINSERT = "INSERT INTO %s.loaderMapStore(%s) VALUES('%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s',	'%s');"
var appLoaderMapStoreSQLDELETE = "DELETE FROM %s.loaderMapStore WHERE id='%s';"
var appLoaderMapStoreSQLSELECT = "SELECT %s FROM %s.loaderMapStore;"
var appLoaderMapStoreSQLGET = "SELECT %s FROM %s.loaderMapStore WHERE id='%s';"
var appLoaderMapStoreSQLSELECTBYLOADER = "SELECT %s FROM %s.loaderMapStore WHERE loader='%s';"

//appLoaderMapStorePage is cheese
type appLoaderMapStoreListPage struct {
	UserMenu            []AppMenuItem
	UserRole            string
	UserNavi            string
	Title               string
	PageTitle           string
	LoaderMapStoreCount int
	LoaderMapStoreList  []LoaderMapStoreItem
}

//appLoaderMapStorePage is cheese
type appLoaderMapStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id         string
	Name       string
	Position   string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//LoaderMapStoreItem is cheese
type LoaderMapStoreItem struct {
	Id         string
	Name       string
	Position   string
	Loader     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreList"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []LoaderMapStoreItem
	noItems, returnList, _ := GetLoaderMapStoreList()

	pageLoaderMapStoreList := appLoaderMapStoreListPage{
		UserMenu:            GetUserMenu(r),
		UserRole:            GetUserRole(r),
		UserNavi:            "NOT USED",
		Title:               globals.ApplicationProperties["appname"],
		PageTitle:           "List Dispatch",
		LoaderMapStoreCount: noItems,
		LoaderMapStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderMapStoreList)

}

func ViewLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LoaderMapStore")
	_, returnRecord, _ := GetLoaderMapStoreByID(searchID)

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Dispatch",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Name:       returnRecord.Name,
		Position:   returnRecord.Position,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageLoaderMapStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderMapStoreList)

}

func EditLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "LoaderMapStore")
	_, returnRecord, _ := GetLoaderMapStoreByID(searchID)

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "Edit Dispatch",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		Name:       returnRecord.Name,
		Position:   returnRecord.Position,
		Loader:     returnRecord.Loader,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageLoaderMapStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderMapStoreList)

}

func SaveLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s LoaderMapStoreItem

	s.Id = r.FormValue("Id")
	s.Name = r.FormValue("Name")
	s.Position = r.FormValue("Position")
	s.Loader = r.FormValue("Loader")
	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")

	//log.Println("save", s)

	putLoaderMapStore(s)

	ListLoaderMapStoreHandler(w, r)

}

func DeleteLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderMapStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	deleteLoaderMapStore(searchID)
	ListLoaderMapStoreHandler(w, r)

}

func BanLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderMapStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Ban", searchID)
	banLoaderMapStore(searchID)
	ListLoaderMapStoreHandler(w, r)
}

func ActivateLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "LoaderMapStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Activate", searchID)
	activateLoaderMapStore(searchID)
	ListLoaderMapStoreHandler(w, r)

}

func NewLoaderMapStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderMapStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageLoaderMapStoreList := appLoaderMapStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: "View Siena Broker",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageLoaderMapStoreList)

}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreList() (int, []LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLSELECT, appLoaderMapStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appLoaderMapStoreList, _, _ := fetchLoaderMapStoreData(tsql)
	return count, appLoaderMapStoreList, nil
}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreListByLoader(id string) (int, []LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLSELECTBYLOADER, appLoaderMapStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	count, appLoaderMapStoreList, _, _ := fetchLoaderMapStoreData(tsql)
	return count, appLoaderMapStoreList, nil
}

// getLoaderMapStoreList read all employees
func GetLoaderMapStoreByID(id string) (int, LoaderMapStoreItem, error) {
	tsql := fmt.Sprintf(appLoaderMapStoreSQLGET, appLoaderMapStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, LoaderMapStoreItem, _ := fetchLoaderMapStoreData(tsql)
	return 1, LoaderMapStoreItem, nil
}

func putLoaderMapStore(r LoaderMapStoreItem) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
		// Default in info for a new RECORD
	}

	r.SYSUpdated = createDate

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	deletesql := fmt.Sprintf(appLoaderMapStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appLoaderMapStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appLoaderMapStoreSQL, r.Id, r.Name, r.Position, r.Loader, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//	log.Println("DELETE:", deletesql, globals.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, globals.ApplicationDB)

	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//	log.Println(fred2, err2)
	_, err := globals.ApplicationDB.Exec(inserttsql)
	//	log.Println(fred, err)
	if err != nil {
		log.Println(err.Error())
	}
}

func deleteLoaderMapStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appLoaderMapStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, globals.ApplicationDB)
	_, err := globals.ApplicationDB.Exec(deletesql)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(fred2, err2)
}

func banLoaderMapStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderMapStoreByID(id)

	putLoaderMapStore(r)
}

func activateLoaderMapStore(id string) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	_, r, _ := GetLoaderMapStoreByID(id)

	putLoaderMapStore(r)
}

// fetchLoaderMapStoreData read all employees
func fetchLoaderMapStoreData(tsql string) (int, []LoaderMapStoreItem, LoaderMapStoreItem, error) {
	//log.Println(tsql)
	var appLoaderMapStore LoaderMapStoreItem
	var appLoaderMapStoreList []LoaderMapStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appLoaderMapStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlLoaderMapStoreId, &sqlLoaderMapStoreName, &sqlLoaderMapStorePosition, &sqlLoaderMapStoreLoader, &sqlLoaderMapStoreSYSCreated, &sqlLoaderMapStoreSYSWho, &sqlLoaderMapStoreSYSHost, &sqlLoaderMapStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appLoaderMapStore, err
		}
		// Populate Arrays etc.
		appLoaderMapStore.Id = sqlLoaderMapStoreId.String
		appLoaderMapStore.Name = sqlLoaderMapStoreName.String
		appLoaderMapStore.Position = sqlLoaderMapStorePosition.String
		appLoaderMapStore.Loader = sqlLoaderMapStoreLoader.String
		appLoaderMapStore.SYSCreated = sqlLoaderMapStoreSYSCreated.String
		appLoaderMapStore.SYSWho = sqlLoaderMapStoreSYSWho.String
		appLoaderMapStore.SYSHost = sqlLoaderMapStoreSYSHost.String
		appLoaderMapStore.SYSUpdated = sqlLoaderMapStoreSYSUpdated.String
		// no change below
		appLoaderMapStoreList = append(appLoaderMapStoreList, appLoaderMapStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appLoaderMapStoreList, appLoaderMapStore)
	return count, appLoaderMapStoreList, appLoaderMapStore, nil
}
