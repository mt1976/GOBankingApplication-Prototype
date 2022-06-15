package application
// ----------------------------------------------------------------
// Automatically generated  "/application/mandate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:58:56
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//mandate_PageList provides the information for the template for a list of Mandates
type Mandate_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Mandate
}
//Mandate_Redirect provides a page to return to aftern an action
const (
	Mandate_Redirect = dm.Mandate_PathList
)

//mandate_Page provides the information for the template for an individual Mandate
type Mandate_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	MandatedUserKeyCounterpartyFirm         string
	MandatedUserKeyCounterpartyFirm_lookup    []dm.Lookup_Item
	MandatedUserKeyCounterpartyCentre         string
	MandatedUserKeyCounterpartyCentre_lookup    []dm.Lookup_Item
	MandatedUserKeyUserName         string
	TelephoneNumber         string
	EmailAddress         string
	Active         string
	Active_lookup    []dm.Lookup_Item
	FirstName         string
	Surname         string
	DateOfBirth         string
	Postcode         string
	NationalIDNo         string
	PassportNo         string
	Country         string
	Country_lookup    []dm.Lookup_Item
	CountryName         string
	FirmName         string
	CentreName         string
	Notify         string
	Notify_lookup    []dm.Lookup_Item
	SystemUser         string
	CompID         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Mandate_Publish annouces the endpoints available for this object
func Mandate_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Mandate_Path, Mandate_Handler)
	mux.HandleFunc(dm.Mandate_PathList, Mandate_HandlerList)
	mux.HandleFunc(dm.Mandate_PathView, Mandate_HandlerView)
	mux.HandleFunc(dm.Mandate_PathEdit, Mandate_HandlerEdit)
	mux.HandleFunc(dm.Mandate_PathNew, Mandate_HandlerNew)
	mux.HandleFunc(dm.Mandate_PathSave, Mandate_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Mandate_Title)
    core.Catalog_Add(dm.Mandate_Title, dm.Mandate_Path, "", dm.Mandate_QueryString, "Application")
}


//Mandate_HandlerList is the handler for the list page
func Mandate_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Mandate
	noItems, returnList, _ := dao.Mandate_GetList()

	pageDetail := Mandate_PageList{
		Title:            CardTitle(dm.Mandate_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Mandate_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Mandate_TemplateList, w, r, pageDetail)

}


//Mandate_HandlerView is the handler used to View a page
func Mandate_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)

	pageDetail := Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateView, w, r, pageDetail)

}


//Mandate_HandlerEdit is the handler used generate the Edit page
func Mandate_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)
	
	pageDetail := Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateEdit, w, r, pageDetail)
}


//Mandate_HandlerSave is the handler used process the saving of an Mandate
func Mandate_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.Mandate
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.MandatedUserKeyCounterpartyFirm = r.FormValue(dm.Mandate_MandatedUserKeyCounterpartyFirm)
		item.MandatedUserKeyCounterpartyCentre = r.FormValue(dm.Mandate_MandatedUserKeyCounterpartyCentre)
		item.MandatedUserKeyUserName = r.FormValue(dm.Mandate_MandatedUserKeyUserName)
		item.TelephoneNumber = r.FormValue(dm.Mandate_TelephoneNumber)
		item.EmailAddress = r.FormValue(dm.Mandate_EmailAddress)
		item.Active = r.FormValue(dm.Mandate_Active)
		item.FirstName = r.FormValue(dm.Mandate_FirstName)
		item.Surname = r.FormValue(dm.Mandate_Surname)
		item.DateOfBirth = r.FormValue(dm.Mandate_DateOfBirth)
		item.Postcode = r.FormValue(dm.Mandate_Postcode)
		item.NationalIDNo = r.FormValue(dm.Mandate_NationalIDNo)
		item.PassportNo = r.FormValue(dm.Mandate_PassportNo)
		item.Country = r.FormValue(dm.Mandate_Country)
		item.CountryName = r.FormValue(dm.Mandate_CountryName)
		item.FirmName = r.FormValue(dm.Mandate_FirmName)
		item.CentreName = r.FormValue(dm.Mandate_CentreName)
		item.Notify = r.FormValue(dm.Mandate_Notify)
		item.SystemUser = r.FormValue(dm.Mandate_SystemUser)
		item.CompID = r.FormValue(dm.Mandate_CompID)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Mandate_Store(item,r)	
	http.Redirect(w, r, Mandate_Redirect, http.StatusFound)
}


//Mandate_HandlerNew is the handler used process the creation of an Mandate
func Mandate_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(dm.Mandate{} , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Mandate Page 
func mandate_PopulatePage(rD dm.Mandate, pageDetail Mandate_Page) Mandate_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.MandatedUserKeyCounterpartyFirm = rD.MandatedUserKeyCounterpartyFirm
	pageDetail.MandatedUserKeyCounterpartyCentre = rD.MandatedUserKeyCounterpartyCentre
	pageDetail.MandatedUserKeyUserName = rD.MandatedUserKeyUserName
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Active = rD.Active
	pageDetail.FirstName = rD.FirstName
	pageDetail.Surname = rD.Surname
	pageDetail.DateOfBirth = rD.DateOfBirth
	pageDetail.Postcode = rD.Postcode
	pageDetail.NationalIDNo = rD.NationalIDNo
	pageDetail.PassportNo = rD.PassportNo
	pageDetail.Country = rD.Country
	pageDetail.CountryName = rD.CountryName
	pageDetail.FirmName = rD.FirmName
	pageDetail.CentreName = rD.CentreName
	pageDetail.Notify = rD.Notify
	pageDetail.SystemUser = rD.SystemUser
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	pageDetail.MandatedUserKeyCounterpartyFirm_lookup = dao.Firm_GetLookup()
	
	
	
	pageDetail.MandatedUserKeyCounterpartyCentre_lookup = dao.Centre_GetLookup()
	
	
	
	
	
	
	
	
	
	
	pageDetail.Active_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	
	
	
	
	
	
	
	pageDetail.Notify_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	