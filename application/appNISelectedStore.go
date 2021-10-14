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
var appNISelectedStoreSQL = "id, 	_created, 	_who, 	_host, 	_updated"

var sqlNISelectedStoreId, sqlNISelectedStoreSYSCreated, sqlNISelectedStoreSYSWho, sqlNISelectedStoreSYSHost, sqlNISelectedStoreSYSUpdated sql.NullString

var appNISelectedStoreNoParams = strings.Count(appNISelectedStoreSQL, ",") + 1
var appNISelectedStoreParams = strings.Repeat("'%s',", appNISelectedStoreNoParams)
var appNISelectedStoreSQLINSERT = "INSERT INTO %s.niSelectedStore(%s) VALUES(" + strings.TrimRight(appNISelectedStoreParams, ",") + ");"
var appNISelectedStoreSQLDELETE = "DELETE FROM %s.niSelectedStore WHERE id='%s';"
var appNISelectedStoreSQLSELECT = "SELECT %s FROM %s.niSelectedStore;"
var appNISelectedStoreSQLGET = "SELECT %s FROM %s.niSelectedStore WHERE id='%s';"

//appNISelectedStorePage is cheese
type appNISelectedStoreListPage struct {
	UserMenu             []AppMenuItem
	UserRole             string
	UserNavi             string
	Title                string
	PageTitle            string
	NISelectedStoreCount int
	NISelectedStoreList  []AppNISelectedStoreItem
}

//appNISelectedStorePage is cheese
type appNISelectedStorePage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Part Below
	Id         string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

//AppNISelectedStoreItem is cheese
type AppNISelectedStoreItem struct {
	Id         string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

func ListNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "NISelectedStoreList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []AppNISelectedStoreItem

	noItems, returnList, _ := GetNISelectedStoreList(globals.ApplicationDB)

	pageNISelectedStoreList := appNISelectedStoreListPage{
		UserMenu:             GetUserMenu(r),
		UserRole:             GetUserRole(r),
		UserNavi:             "NOT USED",
		Title:                globals.ApplicationProperties["appname"],
		PageTitle:            globals.ApplicationProperties["appname"] + " - " + "LSE Gilts",
		NISelectedStoreCount: noItems,
		NISelectedStoreList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageNISelectedStoreList)

}

func ViewNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "NISelectedStoreView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "NISelectedStore")
	_, returnRecord, _ := GetNISelectedStoreByID(searchID)

	pageCredentialStoreList := appNISelectedStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "LSE Gilts - View",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}

	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func EditNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "NISelectedStoreEdit"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	searchID := GetURLparam(r, "NISelectedStore")
	_, returnRecord, _ := GetNISelectedStoreByID(searchID)

	pageCredentialStoreList := appNISelectedStorePage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - " + "LSE Gilts - Edit",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable
		Id:         returnRecord.Id,
		SYSCreated: returnRecord.SYSCreated,
		SYSWho:     returnRecord.SYSWho,
		SYSHost:    returnRecord.SYSHost,
		SYSUpdated: returnRecord.SYSUpdated,
	}
	//fmt.Println(pageCredentialStoreList)

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, pageCredentialStoreList)

}

func SaveNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	var s AppNISelectedStoreItem

	s.Id = r.FormValue("Id")

	s.SYSCreated = r.FormValue("SYSCreated")
	s.SYSWho = r.FormValue("SYSWho")
	s.SYSHost = r.FormValue("SYSHost")
	s.SYSUpdated = r.FormValue("SYSUpdated")
	//log.Println("save", s)

	putNISelectedStore(s, r)

	ListNISelectedStoreHandler(w, r)

}

func DeleteNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "NISelectedStore")
	serviceMessageAction(inUTL, "Delete", searchID)
	deleteNISelectedStore(searchID)
	ListNISelectedStoreHandler(w, r)

}

func SelectNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "NISelectedStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Ban", searchID)
	banNISelectedStore(searchID, r)
	ListNISelectedStoreHandler(w, r)
}

func DeselectNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "NISelectedStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	serviceMessageAction(inUTL, "Activate", searchID)
	activateNISelectedStore(searchID, r)
	ListNISelectedStoreHandler(w, r)

}

func NewNISelectedStoreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "NISelectedStoreNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	pageCredentialStoreList := appNISelectedStorePage{
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

// getNISelectedStoreList read all employees
func GetNISelectedStoreList(unused *sql.DB) (int, []AppNISelectedStoreItem, error) {
	tsql := fmt.Sprintf(appNISelectedStoreSQLSELECT, appNISelectedStoreSQL, globals.ApplicationPropertiesDB["schema"])
	count, appNISelectedStoreList, _, _ := fetchNISelectedStoreData(globals.ApplicationDB, tsql)
	return count, appNISelectedStoreList, nil
}

// getNISelectedStoreList read all employees
func GetNISelectedStoreByID(id string) (int, AppNISelectedStoreItem, error) {
	tsql := fmt.Sprintf(appNISelectedStoreSQLGET, appNISelectedStoreSQL, globals.ApplicationPropertiesDB["schema"], id)
	_, _, AppNISelectedStoreItem, _ := fetchNISelectedStoreData(globals.ApplicationDB, tsql)
	return 1, AppNISelectedStoreItem, nil
}

func putNISelectedStore(r AppNISelectedStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)
	userID := GetUserName(req)
	putNISelectedStoreUser(r, userID)
}

func PutNISelectedStoreSystem(r AppNISelectedStoreItem) {
	//fmt.Println(credentialStore)
	putNISelectedStoreSystem(r)
}
func putNISelectedStoreSystem(r AppNISelectedStoreItem) {
	//fmt.Println(credentialStore)
	putNISelectedStoreUser(r, "AutoGenerated")
}

func putNISelectedStoreUser(r AppNISelectedStoreItem, userID string) {
	//fmt.Println(credentialStore)

	createDate := time.Now().Format(globals.DATETIMEFORMATUSER)
	if len(r.Id) == 0 {
		r.Id = getNISelectedStoreID(r)
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

	deletesql := fmt.Sprintf(appNISelectedStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], r.Id)
	inserttsql := fmt.Sprintf(appNISelectedStoreSQLINSERT, globals.ApplicationPropertiesDB["schema"], appNISelectedStoreSQL, r.Id, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

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

func deleteNISelectedStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(appNISelectedStoreSQLDELETE, globals.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql)
	_, err2 := globals.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Println(err2.Error())
	}
	//log.Println(fred2, err2)
}

func banNISelectedStore(id string, req *http.Request) {
	//fmt.Println(credentialStore)
	//fmt.Println("RECORD", id)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	_, r, err2 := GetNISelectedStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putNISelectedStore(r, req)
}

func activateNISelectedStore(id string, req *http.Request) {
	fmt.Println("RECORD", id)
	_, r, err2 := GetNISelectedStoreByID(id)
	if err2 != nil {
		log.Println(err2.Error())
	}
	putNISelectedStore(r, req)
}

// fetchNISelectedStoreData read all employees
func fetchNISelectedStoreData(unused *sql.DB, tsql string) (int, []AppNISelectedStoreItem, AppNISelectedStoreItem, error) {
	//log.Println(tsql)
	var appNISelectedStore AppNISelectedStoreItem
	var appNISelectedStoreList []AppNISelectedStoreItem

	rows, err := globals.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appNISelectedStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlNISelectedStoreId, &sqlNISelectedStoreSYSCreated, &sqlNISelectedStoreSYSWho, &sqlNISelectedStoreSYSHost, &sqlNISelectedStoreSYSUpdated)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appNISelectedStore, err
		}
		// Mapping from SQL to Struct
		appNISelectedStore.Id = sqlNISelectedStoreId.String
		appNISelectedStore.SYSCreated = sqlNISelectedStoreSYSCreated.String
		appNISelectedStore.SYSWho = sqlNISelectedStoreSYSWho.String
		appNISelectedStore.SYSHost = sqlNISelectedStoreSYSHost.String
		appNISelectedStore.SYSUpdated = sqlNISelectedStoreSYSUpdated.String
		// no change below
		appNISelectedStoreList = append(appNISelectedStoreList, appNISelectedStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appNISelectedStoreList, appNISelectedStore)
	return count, appNISelectedStoreList, appNISelectedStore, nil
}

func newNISelectedStoreID() string {
	id := uuid.New().String()
	return id
}

// Returns a valid niSelectedStore store ID
func getNISelectedStoreID(r AppNISelectedStoreItem) string {
	var niSelectedStoreID string
	niSelectedStoreID = r.Id
	return niSelectedStoreID
}