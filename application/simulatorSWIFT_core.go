package application
// ----------------------------------------------------------------
// Automatically generated  "/application/simulatorswift.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SimulatorSWIFT (simulatorswift)
// Endpoint 	        : SimulatorSWIFT (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:57
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//SimulatorSWIFT_Publish annouces the endpoints available for this object
func SimulatorSWIFT_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SimulatorSWIFT_Path, SimulatorSWIFT_Handler)
	mux.HandleFunc(dm.SimulatorSWIFT_PathList, SimulatorSWIFT_HandlerList)
	mux.HandleFunc(dm.SimulatorSWIFT_PathView, SimulatorSWIFT_HandlerView)
	mux.HandleFunc(dm.SimulatorSWIFT_PathEdit, SimulatorSWIFT_HandlerEdit)
	mux.HandleFunc(dm.SimulatorSWIFT_PathNew, SimulatorSWIFT_HandlerNew)
	mux.HandleFunc(dm.SimulatorSWIFT_PathSave, SimulatorSWIFT_HandlerSave)
	mux.HandleFunc(dm.SimulatorSWIFT_PathDelete, SimulatorSWIFT_HandlerDelete)
	logs.Publish("Application", dm.SimulatorSWIFT_Title)
    core.Catalog_Add(dm.SimulatorSWIFT_Title, dm.SimulatorSWIFT_Path, "", dm.SimulatorSWIFT_QueryString, "Application")
}


//SimulatorSWIFT_HandlerList is the handler for the list page
func SimulatorSWIFT_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SimulatorSWIFT
	noItems, returnList, _ := dao.SimulatorSWIFT_GetList()

	pageDetail := dm.SimulatorSWIFT_PageList{
		Title:            CardTitle(dm.SimulatorSWIFT_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SimulatorSWIFT_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SimulatorSWIFT_TemplateList, w, r, pageDetail)

}


//SimulatorSWIFT_HandlerView is the handler used to View a page
func SimulatorSWIFT_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)
	_, rD, _ := dao.SimulatorSWIFT_GetByID(searchID)

	pageDetail := dm.SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = simulatorswift_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateView, w, r, pageDetail)

}


//SimulatorSWIFT_HandlerEdit is the handler used generate the Edit page
func SimulatorSWIFT_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)
	_, rD, _ := dao.SimulatorSWIFT_GetByID(searchID)
	
	pageDetail := dm.SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = simulatorswift_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateEdit, w, r, pageDetail)
}


//SimulatorSWIFT_HandlerSave is the handler used process the saving of an SimulatorSWIFT
func SimulatorSWIFT_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.SimulatorSWIFT
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.SimulatorSWIFT_ID_scrn)
		item.FileName = r.FormValue(dm.SimulatorSWIFT_FileName_scrn)
		item.MessageRaw = r.FormValue(dm.SimulatorSWIFT_MessageRaw_scrn)
		item.MessageFmt = r.FormValue(dm.SimulatorSWIFT_MessageFmt_scrn)
		item.Action = r.FormValue(dm.SimulatorSWIFT_Action_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.SimulatorSWIFT_Store(item,r)	
	http.Redirect(w, r, dm.SimulatorSWIFT_Redirect, http.StatusFound)
}


//SimulatorSWIFT_HandlerNew is the handler used process the creation of an SimulatorSWIFT
func SimulatorSWIFT_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.SimulatorSWIFT_New()

	pageDetail := dm.SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_New),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = simulatorswift_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateNew, w, r, pageDetail)

}	


//SimulatorSWIFT_HandlerDelete is the handler used process the deletion of an SimulatorSWIFT
func SimulatorSWIFT_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)

	dao.SimulatorSWIFT_Delete(searchID)	

	http.Redirect(w, r, dm.SimulatorSWIFT_Redirect, http.StatusFound)
}


// Builds/Popuplates the SimulatorSWIFT Page 
func simulatorswift_PopulatePage(rD dm.SimulatorSWIFT, pageDetail dm.SimulatorSWIFT_Page) dm.SimulatorSWIFT_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.FileName = rD.FileName
	pageDetail.MessageRaw = rD.MessageRaw
	pageDetail.MessageFmt = rD.MessageFmt
	pageDetail.Action = rD.Action
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.FileName_props = rD.FileName_props
	pageDetail.MessageRaw_props = rD.MessageRaw_props
	pageDetail.MessageFmt_props = rD.MessageFmt_props
	pageDetail.Action_props = rD.Action_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	