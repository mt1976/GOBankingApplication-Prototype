package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterparty.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:49:58
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterparty_PageList provides the information for the template for a list of Counterpartys
type Counterparty_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Counterparty
}
//Counterparty_Redirect provides a page to return to aftern an action
const (
	Counterparty_Redirect = dm.Counterparty_PathList
)

//counterparty_Page provides the information for the template for an individual Counterparty
type Counterparty_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameCentre         string
	NameCentre_lookup    []dm.Lookup_Item
	NameFirm         string
	NameFirm_lookup    []dm.Lookup_Item
	FullName         string
	TelephoneNumber         string
	EmailAddress         string
	CustomerType         string
	CustomerType_lookup    []dm.Lookup_Item
	AccountOfficer         string
	CountryCode         string
	CountryCode_lookup    []dm.Lookup_Item
	SectorCode         string
	SectorCode_lookup    []dm.Lookup_Item
	CpartyGroupName         string
	CpartyGroupName_lookup    []dm.Lookup_Item
	Notes         string
	Owner         string
	Authorised         string
	Authorised_lookup    []dm.Lookup_Item
	NameFirmName         string
	NameCentreName         string
	CountryCodeName         string
	SectorCodeName         string
	CompID         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Counterparty_Publish annouces the endpoints available for this object
func Counterparty_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Counterparty_Path, Counterparty_Handler)
	mux.HandleFunc(dm.Counterparty_PathList, Counterparty_HandlerList)
	mux.HandleFunc(dm.Counterparty_PathView, Counterparty_HandlerView)
	mux.HandleFunc(dm.Counterparty_PathEdit, Counterparty_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Counterparty_PathSave, Counterparty_HandlerSave)
	mux.HandleFunc(dm.Counterparty_PathDelete, Counterparty_HandlerDelete)
	logs.Publish("Application", dm.Counterparty_Title)
    core.Catalog_Add(dm.Counterparty_Title, dm.Counterparty_Path, "", dm.Counterparty_QueryString, "Application")
}


//Counterparty_HandlerList is the handler for the list page
func Counterparty_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Counterparty
	noItems, returnList, _ := dao.Counterparty_GetList()

	pageDetail := Counterparty_PageList{
		Title:            CardTitle(dm.Counterparty_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Counterparty_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Counterparty_TemplateList, w, r, pageDetail)

}


//Counterparty_HandlerView is the handler used to View a page
func Counterparty_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)
	_, rD, _ := dao.Counterparty_GetByID(searchID)

	pageDetail := Counterparty_Page{
		Title:       CardTitle(dm.Counterparty_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Counterparty_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterparty_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Counterparty_TemplateView, w, r, pageDetail)

}


//Counterparty_HandlerEdit is the handler used generate the Edit page
func Counterparty_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)
	_, rD, _ := dao.Counterparty_GetByID(searchID)
	
	pageDetail := Counterparty_Page{
		Title:       CardTitle(dm.Counterparty_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Counterparty_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterparty_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Counterparty_TemplateEdit, w, r, pageDetail)
}


//Counterparty_HandlerSave is the handler used process the saving of an Counterparty
func Counterparty_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.Counterparty
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.NameCentre = r.FormValue(dm.Counterparty_NameCentre)
		item.NameFirm = r.FormValue(dm.Counterparty_NameFirm)
		item.FullName = r.FormValue(dm.Counterparty_FullName)
		item.TelephoneNumber = r.FormValue(dm.Counterparty_TelephoneNumber)
		item.EmailAddress = r.FormValue(dm.Counterparty_EmailAddress)
		item.CustomerType = r.FormValue(dm.Counterparty_CustomerType)
		item.AccountOfficer = r.FormValue(dm.Counterparty_AccountOfficer)
		item.CountryCode = r.FormValue(dm.Counterparty_CountryCode)
		item.SectorCode = r.FormValue(dm.Counterparty_SectorCode)
		item.CpartyGroupName = r.FormValue(dm.Counterparty_CpartyGroupName)
		item.Notes = r.FormValue(dm.Counterparty_Notes)
		item.Owner = r.FormValue(dm.Counterparty_Owner)
		item.Authorised = r.FormValue(dm.Counterparty_Authorised)
		item.NameFirmName = r.FormValue(dm.Counterparty_NameFirmName)
		item.NameCentreName = r.FormValue(dm.Counterparty_NameCentreName)
		item.CountryCodeName = r.FormValue(dm.Counterparty_CountryCodeName)
		item.SectorCodeName = r.FormValue(dm.Counterparty_SectorCodeName)
		item.CompID = r.FormValue(dm.Counterparty_CompID)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Counterparty_Store(item,r)	
	http.Redirect(w, r, Counterparty_Redirect, http.StatusFound)
}



//Counterparty_HandlerDelete is the handler used process the deletion of an Counterparty
func Counterparty_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)

	dao.Counterparty_Delete(searchID)	

	http.Redirect(w, r, Counterparty_Redirect, http.StatusFound)
}


// Builds/Popuplates the Counterparty Page 
func counterparty_PopulatePage(rD dm.Counterparty, pageDetail Counterparty_Page) Counterparty_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.FullName = rD.FullName
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.CustomerType = rD.CustomerType
	pageDetail.AccountOfficer = rD.AccountOfficer
	pageDetail.CountryCode = rD.CountryCode
	pageDetail.SectorCode = rD.SectorCode
	pageDetail.CpartyGroupName = rD.CpartyGroupName
	pageDetail.Notes = rD.Notes
	pageDetail.Owner = rD.Owner
	pageDetail.Authorised = rD.Authorised
	pageDetail.NameFirmName = rD.NameFirmName
	pageDetail.NameCentreName = rD.NameCentreName
	pageDetail.CountryCodeName = rD.CountryCodeName
	pageDetail.SectorCodeName = rD.SectorCodeName
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	pageDetail.NameCentre_lookup = dao.Centre_GetLookup()
	
	
	
	pageDetail.NameFirm_lookup = dao.Firm_GetLookup()
	
	
	
	
	
	
	
	
	
	
	pageDetail.CustomerType_lookup = dao.StubLists_Get("counterpartytypes")
	
	
	
	
	pageDetail.CountryCode_lookup = dao.Country_GetLookup()
	
	
	
	pageDetail.SectorCode_lookup = dao.Sector_GetLookup()
	
	
	
	pageDetail.CpartyGroupName_lookup = dao.CounterpartyGroup_GetLookup()
	
	
	
	
	
	
	
	
	pageDetail.Authorised_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	