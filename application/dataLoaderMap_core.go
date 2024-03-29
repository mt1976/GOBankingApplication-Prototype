package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dataloadermap.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderMap (dataloadermap)
// Endpoint 	        : DataLoaderMap (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//DataLoaderMap_Publish annouces the endpoints available for this object
func DataLoaderMap_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DataLoaderMap_PathList, DataLoaderMap_HandlerList)
	mux.HandleFunc(dm.DataLoaderMap_PathView, DataLoaderMap_HandlerView)
	mux.HandleFunc(dm.DataLoaderMap_PathEdit, DataLoaderMap_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.DataLoaderMap_PathSave, DataLoaderMap_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DataLoaderMap_Title)
    //No API
}


//DataLoaderMap_HandlerList is the handler for the list page
func DataLoaderMap_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DataLoaderMap
	noItems, returnList, _ := dao.DataLoaderMap_GetList()

	pageDetail := dm.DataLoaderMap_PageList{
		Title:            CardTitle(dm.DataLoaderMap_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DataLoaderMap_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DataLoaderMap_TemplateList, w, r, pageDetail)

}


//DataLoaderMap_HandlerView is the handler used to View a page
func DataLoaderMap_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderMap_QueryString)
	_, rD, _ := dao.DataLoaderMap_GetByID(searchID)

	pageDetail := dm.DataLoaderMap_Page{
		Title:       CardTitle(dm.DataLoaderMap_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataLoaderMap_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloadermap_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoaderMap_TemplateView, w, r, pageDetail)

}


//DataLoaderMap_HandlerEdit is the handler used generate the Edit page
func DataLoaderMap_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderMap_QueryString)
	_, rD, _ := dao.DataLoaderMap_GetByID(searchID)
	
	pageDetail := dm.DataLoaderMap_Page{
		Title:       CardTitle(dm.DataLoaderMap_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataLoaderMap_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloadermap_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoaderMap_TemplateEdit, w, r, pageDetail)
}


//DataLoaderMap_HandlerSave is the handler used process the saving of an DataLoaderMap
func DataLoaderMap_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.DataLoaderMap
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.DataLoaderMap_SYSId_scrn)
		item.Id = r.FormValue(dm.DataLoaderMap_Id_scrn)
		item.Name = r.FormValue(dm.DataLoaderMap_Name_scrn)
		item.Position = r.FormValue(dm.DataLoaderMap_Position_scrn)
		item.Loader = r.FormValue(dm.DataLoaderMap_Loader_scrn)
		item.SYSCreated = r.FormValue(dm.DataLoaderMap_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.DataLoaderMap_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.DataLoaderMap_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.DataLoaderMap_SYSUpdated_scrn)
		item.Int_position = r.FormValue(dm.DataLoaderMap_Int_position_scrn)
		item.SYSCreatedBy = r.FormValue(dm.DataLoaderMap_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.DataLoaderMap_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.DataLoaderMap_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.DataLoaderMap_SYSUpdatedHost_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DataLoaderMap_Store(item,r)	
	http.Redirect(w, r, dm.DataLoaderMap_Redirect, http.StatusFound)
}




// Builds/Popuplates the DataLoaderMap Page 
func dataloadermap_PopulatePage(rD dm.DataLoaderMap, pageDetail dm.DataLoaderMap_Page) dm.DataLoaderMap_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Position = rD.Position
	pageDetail.Loader = rD.Loader
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Int_position = rD.Int_position
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	pageDetail.Loader_lookup = dao.DataLoader_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Position_props = rD.Position_props
	pageDetail.Loader_props = rD.Loader_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.Int_position_props = rD.Int_position_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	