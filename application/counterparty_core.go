package application

// ----------------------------------------------------------------
// Automatically generated  "/application/counterparty.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:55
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//counterparty_PageList provides the information for the template for a list of Counterpartys
type Counterparty_PageList struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []dm.Counterparty
}

//counterparty_Page provides the information for the template for an individual Counterparty
type Counterparty_Page struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	Title     string
	PageTitle string
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
	NameCentre      string
	NameFirm        string
	FullName        string
	TelephoneNumber string
	EmailAddress    string
	CustomerType    string
	AccountOfficer  string
	CountryCode     string
	SectorCode      string
	CpartyGroupName string
	Notes           string
	Owner           string
	Authorised      string
	NameFirmName    string
	NameCentreName  string
	CountryCodeName string
	SectorCodeName  string
	CompID          string

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
}

const (
	Counterparty_Redirect = dm.Counterparty_PathList
)

//Counterparty_Publish annouces the endpoints available for this object
func Counterparty_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Counterparty_PathList, Counterparty_HandlerList)
	mux.HandleFunc(dm.Counterparty_PathView, Counterparty_HandlerView)
	mux.HandleFunc(dm.Counterparty_PathEdit, Counterparty_HandlerEdit)

	mux.HandleFunc(dm.Counterparty_PathSave, Counterparty_HandlerSave)
	mux.HandleFunc(dm.Counterparty_PathDelete, Counterparty_HandlerDelete)
	logs.Publish("Siena", dm.Counterparty_Title)

}

//Counterparty_HandlerList is the handler for the list page
func Counterparty_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Counterparty
	noItems, returnList, _ := dao.Counterparty_GetList()

	pageDetail := Counterparty_PageList{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Counterparty_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

	ExecuteTemplate(dm.Counterparty_TemplateList, w, r, pageDetail)

}

//Counterparty_HandlerView is the handler used to View a page
func Counterparty_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)
	_, rD, _ := dao.Counterparty_GetByID(searchID)

	pageDetail := Counterparty_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Counterparty_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	ExecuteTemplate(dm.Counterparty_TemplateView, w, r, pageDetail)

}

//Counterparty_HandlerEdit is the handler used generate the Edit page
func Counterparty_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)
	_, rD, _ := dao.Counterparty_GetByID(searchID)

	pageDetail := Counterparty_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Counterparty_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	ExecuteTemplate(dm.Counterparty_TemplateEdit, w, r, pageDetail)

}

//Counterparty_HandlerSave is the handler used process the saving of an Counterparty
func Counterparty_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("CompID"))

	var item dm.Counterparty
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	dao.Counterparty_Store(item)

	http.Redirect(w, r, Counterparty_Redirect, http.StatusFound)
}

//Counterparty_HandlerNew is the handler used process the creation of an Counterparty
func Counterparty_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Counterparty_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Counterparty_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
	pageDetail.NameCentre = ""
	pageDetail.NameFirm = ""
	pageDetail.FullName = ""
	pageDetail.TelephoneNumber = ""
	pageDetail.EmailAddress = ""
	pageDetail.CustomerType = ""
	pageDetail.AccountOfficer = ""
	pageDetail.CountryCode = ""
	pageDetail.SectorCode = ""
	pageDetail.CpartyGroupName = ""
	pageDetail.Notes = ""
	pageDetail.Owner = ""
	pageDetail.Authorised = ""
	pageDetail.NameFirmName = ""
	pageDetail.NameCentreName = ""
	pageDetail.CountryCodeName = ""
	pageDetail.SectorCodeName = ""
	pageDetail.CompID = ""
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	ExecuteTemplate(dm.Counterparty_TemplateNew, w, r, pageDetail)

}

//Counterparty_HandlerDelete is the handler used process the deletion of an Counterparty
func Counterparty_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Counterparty_QueryString)

	dao.Counterparty_Delete(searchID)

	http.Redirect(w, r, Counterparty_Redirect, http.StatusFound)
}
