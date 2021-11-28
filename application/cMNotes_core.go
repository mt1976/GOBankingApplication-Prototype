package application
// ----------------------------------------------------------------
// Automatically generated  "/application/cmnotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:54:53
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

//cmnotes_PageList provides the information for the template for a list of CMNotess
type CMNotes_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CMNotes
}

//cmnotes_Page provides the information for the template for an individual CMNotes
type CMNotes_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		NoteId string
		StreamId string
		Summary string
		Details string
		RecordState string
		CreatedBy string
		CreatedDateTime string
	
	
	
	
	
	
	
	
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
}

const (
	CMNotes_Redirect = dm.CMNotes_PathList
)

//CMNotes_Publish annouces the endpoints available for this object
func CMNotes_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.CMNotes_PathList, CMNotes_HandlerList)
	mux.HandleFunc(dm.CMNotes_PathView, CMNotes_HandlerView)
	
	
	
	
	logs.Publish("Siena", dm.CMNotes_Title)
	
}

//CMNotes_HandlerList is the handler for the list page
func CMNotes_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CMNotes
	noItems, returnList, _ := dao.CMNotes_GetList()


	pageDetail := CMNotes_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.CMNotes_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CMNotes_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//CMNotes_HandlerView is the handler used to View a page
func CMNotes_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)
	_, rD, _ := dao.CMNotes_GetByID(searchID)

	pageDetail := CMNotes_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = rD.NoteId
pageDetail.StreamId = rD.StreamId
pageDetail.Summary = rD.Summary
pageDetail.Details = rD.Details
pageDetail.RecordState = rD.RecordState
pageDetail.CreatedBy = rD.CreatedBy
pageDetail.CreatedDateTime = rD.CreatedDateTime
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.CMNotes_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CMNotes_HandlerEdit is the handler used generate the Edit page
func CMNotes_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)
	_, rD, _ := dao.CMNotes_GetByID(searchID)
	
	pageDetail := CMNotes_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = rD.NoteId
pageDetail.StreamId = rD.StreamId
pageDetail.Summary = rD.Summary
pageDetail.Details = rD.Details
pageDetail.RecordState = rD.RecordState
pageDetail.CreatedBy = rD.CreatedBy
pageDetail.CreatedDateTime = rD.CreatedDateTime
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CMNotes_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CMNotes_HandlerSave is the handler used process the saving of an CMNotes
func CMNotes_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("NoteId"))

	var item dm.CMNotes
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		item.NoteId = r.FormValue(dm.CMNotes_NoteId)
		item.StreamId = r.FormValue(dm.CMNotes_StreamId)
		item.Summary = r.FormValue(dm.CMNotes_Summary)
		item.Details = r.FormValue(dm.CMNotes_Details)
		item.RecordState = r.FormValue(dm.CMNotes_RecordState)
		item.CreatedBy = r.FormValue(dm.CMNotes_CreatedBy)
		item.CreatedDateTime = r.FormValue(dm.CMNotes_CreatedDateTime)
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	dao.CMNotes_Store(item)	

	http.Redirect(w, r, CMNotes_Redirect, http.StatusFound)
}

//CMNotes_HandlerNew is the handler used process the creation of an CMNotes
func CMNotes_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CMNotes_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = ""
pageDetail.StreamId = ""
pageDetail.Summary = ""
pageDetail.Details = ""
pageDetail.RecordState = ""
pageDetail.CreatedBy = ""
pageDetail.CreatedDateTime = ""
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CMNotes_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CMNotes_HandlerDelete is the handler used process the deletion of an CMNotes
func CMNotes_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)

	dao.CMNotes_Delete(searchID)	

	http.Redirect(w, r, CMNotes_Redirect, http.StatusFound)
}