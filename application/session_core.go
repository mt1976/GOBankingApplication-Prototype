package application
// ----------------------------------------------------------------
// Automatically generated  "/application/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 20:23:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//session_PageList provides the information for the template for a list of Sessions
type Session_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Session
}

//session_Page provides the information for the template for an individual Session
type Session_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Apptoken string
		Createdate string
		Createtime string
		Uniqueid string
		Sessiontoken string
		Username string
		Password string
		Userip string
		Userhost string
		Appip string
		Apphost string
		Issued string
		Expiry string
		Expiryraw string
		Brand string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Id string
		Expires string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
		SessionRole string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
}

const (
	Session_Redirect = dm.Session_PathList
)

//Session_Publish annouces the endpoints available for this object
func Session_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Session_PathList, Session_HandlerList)
	mux.HandleFunc(dm.Session_PathView, Session_HandlerView)
	
	
	
	
	logs.Publish("Application", dm.Session_Title)
	
}

//Session_HandlerList is the handler for the list page
func Session_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Session
	noItems, returnList, _ := dao.Session_GetList()


	pageDetail := Session_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Session_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Session_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Session_HandlerView is the handler used to View a page
func Session_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Session_QueryString)
	_, rD, _ := dao.Session_GetByID(searchID)

	pageDetail := Session_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Session_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Apptoken = rD.Apptoken
pageDetail.Createdate = rD.Createdate
pageDetail.Createtime = rD.Createtime
pageDetail.Uniqueid = rD.Uniqueid
pageDetail.Sessiontoken = rD.Sessiontoken
pageDetail.Username = rD.Username
pageDetail.Password = rD.Password
pageDetail.Userip = rD.Userip
pageDetail.Userhost = rD.Userhost
pageDetail.Appip = rD.Appip
pageDetail.Apphost = rD.Apphost
pageDetail.Issued = rD.Issued
pageDetail.Expiry = rD.Expiry
pageDetail.Expiryraw = rD.Expiryraw
pageDetail.Brand = rD.Brand
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Id = rD.Id
pageDetail.Expires = rD.Expires
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.SessionRole = rD.SessionRole
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Session_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Session_HandlerEdit is the handler used generate the Edit page
func Session_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Session_QueryString)
	_, rD, _ := dao.Session_GetByID(searchID)
	
	pageDetail := Session_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Session_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Apptoken = rD.Apptoken
pageDetail.Createdate = rD.Createdate
pageDetail.Createtime = rD.Createtime
pageDetail.Uniqueid = rD.Uniqueid
pageDetail.Sessiontoken = rD.Sessiontoken
pageDetail.Username = rD.Username
pageDetail.Password = rD.Password
pageDetail.Userip = rD.Userip
pageDetail.Userhost = rD.Userhost
pageDetail.Appip = rD.Appip
pageDetail.Apphost = rD.Apphost
pageDetail.Issued = rD.Issued
pageDetail.Expiry = rD.Expiry
pageDetail.Expiryraw = rD.Expiryraw
pageDetail.Brand = rD.Brand
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Id = rD.Id
pageDetail.Expires = rD.Expires
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.SessionRole = rD.SessionRole
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Session_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Session_HandlerSave is the handler used process the saving of an Session
func Session_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Session

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Session_SYSId)
		item.Apptoken = r.FormValue(dm.Session_Apptoken)
		item.Createdate = r.FormValue(dm.Session_Createdate)
		item.Createtime = r.FormValue(dm.Session_Createtime)
		item.Uniqueid = r.FormValue(dm.Session_Uniqueid)
		item.Sessiontoken = r.FormValue(dm.Session_Sessiontoken)
		item.Username = r.FormValue(dm.Session_Username)
		item.Password = r.FormValue(dm.Session_Password)
		item.Userip = r.FormValue(dm.Session_Userip)
		item.Userhost = r.FormValue(dm.Session_Userhost)
		item.Appip = r.FormValue(dm.Session_Appip)
		item.Apphost = r.FormValue(dm.Session_Apphost)
		item.Issued = r.FormValue(dm.Session_Issued)
		item.Expiry = r.FormValue(dm.Session_Expiry)
		item.Expiryraw = r.FormValue(dm.Session_Expiryraw)
		item.Brand = r.FormValue(dm.Session_Brand)
		item.SYSCreated = r.FormValue(dm.Session_SYSCreated)
		item.SYSWho = r.FormValue(dm.Session_SYSWho)
		item.SYSHost = r.FormValue(dm.Session_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Session_SYSUpdated)
		item.Id = r.FormValue(dm.Session_Id)
		item.Expires = r.FormValue(dm.Session_Expires)
		item.SYSCreatedBy = r.FormValue(dm.Session_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Session_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Session_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Session_SYSUpdatedHost)
		item.SessionRole = r.FormValue(dm.Session_SessionRole)
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	dao.Session_Store(item)	

	http.Redirect(w, r, Session_Redirect, http.StatusFound)
}

//Session_HandlerNew is the handler used process the creation of an Session
func Session_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Session_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Session_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Apptoken = ""
pageDetail.Createdate = ""
pageDetail.Createtime = ""
pageDetail.Uniqueid = ""
pageDetail.Sessiontoken = ""
pageDetail.Username = ""
pageDetail.Password = ""
pageDetail.Userip = ""
pageDetail.Userhost = ""
pageDetail.Appip = ""
pageDetail.Apphost = ""
pageDetail.Issued = ""
pageDetail.Expiry = ""
pageDetail.Expiryraw = ""
pageDetail.Brand = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Id = ""
pageDetail.Expires = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.SessionRole = ""
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Session_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Session_HandlerDelete is the handler used process the deletion of an Session
func Session_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Session_QueryString)

	dao.Session_Delete(searchID)	

	http.Redirect(w, r, Session_Redirect, http.StatusFound)
}