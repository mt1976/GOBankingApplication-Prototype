package application
// ----------------------------------------------------------------
// Automatically generated  "/application/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//translation_PageList provides the information for the template for a list of Translations
type Translation_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Translation
}
//Translation_Redirect provides a page to return to aftern an action
const (
	Translation_Redirect = dm.Translation_PathList
)

//translation_Page provides the information for the template for an individual Translation
type Translation_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	Id         string
	Class         string
	Message         string
	Translation         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	SYSUpdated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Translation_Publish annouces the endpoints available for this object
func Translation_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Translation_PathList, Translation_HandlerList)
	mux.HandleFunc(dm.Translation_PathView, Translation_HandlerView)
	mux.HandleFunc(dm.Translation_PathEdit, Translation_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Translation_PathSave, Translation_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Translation_Title)
    //No API
}


//Translation_HandlerList is the handler for the list page
func Translation_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Translation
	noItems, returnList, _ := dao.Translation_GetList()

	pageDetail := Translation_PageList{
		Title:            CardTitle(dm.Translation_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Translation_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Translation_TemplateList, w, r, pageDetail)

}


//Translation_HandlerView is the handler used to View a page
func Translation_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)

	pageDetail := Translation_Page{
		Title:       CardTitle(dm.Translation_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = translation_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Translation_TemplateView, w, r, pageDetail)

}


//Translation_HandlerEdit is the handler used generate the Edit page
func Translation_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)
	
	pageDetail := Translation_Page{
		Title:       CardTitle(dm.Translation_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = translation_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Translation_TemplateEdit, w, r, pageDetail)
}


//Translation_HandlerSave is the handler used process the saving of an Translation
func Translation_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Translation
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Translation_SYSId)
		item.Id = r.FormValue(dm.Translation_Id)
		item.Class = r.FormValue(dm.Translation_Class)
		item.Message = r.FormValue(dm.Translation_Message)
		item.Translation = r.FormValue(dm.Translation_Translation)
		item.SYSCreated = r.FormValue(dm.Translation_SYSCreated)
		item.SYSWho = r.FormValue(dm.Translation_SYSWho)
		item.SYSHost = r.FormValue(dm.Translation_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Translation_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Translation_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Translation_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Translation_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Translation_SYSUpdatedHost)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Translation_Store(item,r)	
	http.Redirect(w, r, Translation_Redirect, http.StatusFound)
}




// Builds/Popuplates the Translation Page 
func translation_PopulatePage(rD dm.Translation, pageDetail Translation_Page) Translation_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Class = rD.Class
	pageDetail.Message = rD.Message
	pageDetail.Translation = rD.Translation
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	