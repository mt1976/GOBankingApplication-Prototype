package application

import (
	"html/template"
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// Defines the Fields to Fetch from SQL
/*
var simFundsCheckSQL = "id, 	name, 	staticIn, 	staticout, 	txnin, 	txnout, 	fundscheckin, 	fundscheckout, 	_created, 	_who, 	_host, 	_updated"

var sqlFundsCheckId, sqlFundsCheckName, sqlFundsCheckStaticin, sqlFundsCheckStaticout, sqlFundsCheckTxnin, sqlFundsCheckTxnout, sqlFundsCheckFundscheckin, sqlFundsCheckFundscheckout, sqlFundsCheckSYSCreated, sqlFundsCheckSYSWho, sqlFundsCheckSYSHost, sqlFundsCheckSYSUpdated sql.NullString

var simFundsCheckNoParams = strings.Count(simFundsCheckSQL, ",") + 1
var simFundsCheckParams = strings.Repeat("'%s',", simFundsCheckNoParams)
var simFundsCheckSQLINSERT = "INSERT INTO %s.fundsCheck(%s) VALUES(" + strings.TrimRight(simFundsCheckParams, ",") + ");"
var simFundsCheckSQLDELETE = "DELETE FROM %s.fundsCheck WHERE id='%s';"
var simFundsCheckSQLSELECT = "SELECT %s FROM %s.fundsCheck;"
var simFundsCheckSQLGET = "SELECT %s FROM %s.fundsCheck WHERE id='%s';"
*/
//simFundsCheckPage is cheese
type Simulator_FundsChecker_PageList struct {
	UserMenu        []dm.AppMenuItem
	UserRole        string
	UserNavi        string
	Title           string
	PageTitle       string
	FundsCheckCount int
	FundsCheckList  []dm.Simulator_FundsChecker_Item
}

//Simulator_FundsChecker_Page is cheese
type Simulator_FundsChecker_Page struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id      string
	Name    string
	Source  string
	Content string
	Request dm.Simulator_FundsChecker_Request
}

//NegotiableInstrument_Publish annouces the endpoints available for this object
func Simulator_FundsChecker_Publish_Impl(mux http.ServeMux) {
	//mux.HandleFunc("/listGiltsDataStore/", application.ListLSEGiltsDataStoreHandler)
	//mux.HandleFunc("/viewLSEGiltsDataStore/", application.ViewLSEGiltsDataStoreHandler)
	mux.HandleFunc(dm.Simulator_FundsChecker_PathList, Simulator_FundsChecker_HandlerList)
	mux.HandleFunc(dm.Simulator_FundsChecker_PathView, Simulator_FundsChecker_HandlerView)
	mux.HandleFunc(dm.Simulator_FundsChecker_PathAction, Simulator_FundsChecker_HandlerAction)
	mux.HandleFunc(dm.Simulator_FundsChecker_PathSubmit, Simulator_FundsChecker_HandlerSubmit)
	mux.HandleFunc(dm.Simulator_FundsChecker_PathDelete, Simulator_FundsChecker_HandlerDelete)

	logs.Publish("Application", dm.Simulator_FundsChecker_Title)

}

func Simulator_FundsChecker_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	noItems, returnList, _ := dao.Simulator_FundsChecker_GetList()

	fundsCheckPage := Simulator_FundsChecker_PageList{
		UserMenu:        UserMenu_Get(r),
		UserRole:        core.GetUserRole(r),
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       PageTitle(dm.Simulator_FundsChecker_Title, core.Action_Requests),
		FundsCheckCount: noItems,
		FundsCheckList:  returnList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Simulator_FundsChecker_TemplateList, core.GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

func Simulator_FundsChecker_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	fundsCheckPage := simulator_FundsCheck_BuildPage(w, r)
	fundsCheckPage.Title = core.ApplicationProperties["appname"]
	fundsCheckPage.PageTitle = PageTitle(dm.Simulator_FundsChecker_Title, core.Action_View)

	//log.Println(fundsCheckPage)
	t, _ := template.ParseFiles(core.GetTemplateID(dm.Simulator_FundsChecker_TemplateView, core.GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

func Simulator_FundsChecker_HandlerAction(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	fundsCheckPage := simulator_FundsCheck_BuildPage(w, r)

	fundsCheckPage.Title = core.ApplicationProperties["appname"]
	fundsCheckPage.PageTitle = PageTitle(dm.Simulator_FundsChecker_Title, core.Action_Process)

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Simulator_FundsChecker_TemplateAction, core.GetUserRole(r)))
	t.Execute(w, fundsCheckPage)
}

func Simulator_FundsChecker_HandlerSubmit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	core.ServiceMessageAction(inUTL, "Save", r.FormValue("Id"))

	thisID := r.FormValue("ID")
	balance := r.FormValue("Balance")
	resultCode := r.FormValue("ResultCode")

	dao.Simulator_FundsChecker_Store(thisID, balance, resultCode)

	http.Redirect(w, r, dm.Simulator_FundsChecker_PathList, http.StatusFound)
}

func Simulator_FundsChecker_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "FundsCheck")
	core.ServiceMessageAction(inUTL, "Delete", searchID)
	dao.Simulator_FundsChecker_DeleteByID(searchID)

	http.Redirect(w, r, dm.Simulator_FundsChecker_PathList, http.StatusFound)
}

func simulator_FundsCheck_BuildPage(w http.ResponseWriter, r *http.Request) Simulator_FundsChecker_Page {

	searchID := core.GetURLparam(r, "FundsCheck")
	_, returnRecord, _ := dao.Simulator_FundsChecker_GetByID(searchID)

	fundsCheckPage := Simulator_FundsChecker_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Simulator_FundsChecker_Title, core.Action_Request),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		// Above are mandatory
		// Below are variable
		Id:      returnRecord.Id,
		Name:    returnRecord.Name,
		Source:  returnRecord.Source,
		Content: returnRecord.Content,
		Request: returnRecord.Request,
	}

	return fundsCheckPage //fmt.Println(fundsCheckPage)
}
