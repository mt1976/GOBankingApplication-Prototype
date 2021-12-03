package application

// ----------------------------------------------------------------
// Automatically generated  "/application/mandate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:59
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//mandate_PageList provides the information for the template for a list of Mandates
type Mandate_PageList struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []dm.Mandate
}

//mandate_Page provides the information for the template for an individual Mandate
type Mandate_Page struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	Title     string
	PageTitle string
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
	MandatedUserKeyCounterpartyFirm   string
	MandatedUserKeyCounterpartyCentre string
	MandatedUserKeyUserName           string
	TelephoneNumber                   string
	EmailAddress                      string
	Active                            string
	FirstName                         string
	Surname                           string
	DateOfBirth                       string
	Postcode                          string
	NationalIDNo                      string
	PassportNo                        string
	Country                           string
	CountryName                       string
	FirmName                          string
	CentreName                        string
	Notify                            string
	SystemUser                        string
	CompID                            string
	Country_Impl                      string
	Firm_Impl                         string
	Centre_Impl                       string

	Country_Impl_List []dm.Country
	Firm_Impl_List    []dm.Firm
	Centre_Impl_List  []dm.Centre

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
}

const (
	Mandate_Redirect = dm.Mandate_PathList
)

//Mandate_Publish annouces the endpoints available for this object
func Mandate_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Mandate_PathList, Mandate_HandlerList)
	mux.HandleFunc(dm.Mandate_PathView, Mandate_HandlerView)
	mux.HandleFunc(dm.Mandate_PathEdit, Mandate_HandlerEdit)
	mux.HandleFunc(dm.Mandate_PathNew, Mandate_HandlerNew)
	mux.HandleFunc(dm.Mandate_PathSave, Mandate_HandlerSave)

	logs.Publish("Siena", dm.Mandate_Title)

	// mandate_PublishImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override function should be defined as
	// mandate_PublishImpl(mux http.ServeMux) {...}
	// TODO - this is a temporary hack to get the special case working
	// Add to main.go >>> mandate_PublishImpl(mux)

}

//Mandate_HandlerList is the handler for the list page
func Mandate_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Mandate
	noItems, returnList, _ := dao.Mandate_GetList()

	pageDetail := Mandate_PageList{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

	ExecuteTemplate(dm.Mandate_TemplateList, w, r, pageDetail)

}

//Mandate_HandlerView is the handler used to View a page
func Mandate_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)

	pageDetail := Mandate_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Mandate_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	_, Country_Lookup_Name, _ := dao.Country_GetByID(rD.Country)
	pageDetail.Country_Impl = Country_Lookup_Name.Name
	_, MandatedUserKeyCounterpartyFirm_Lookup_FullName, _ := dao.Firm_GetByID(rD.MandatedUserKeyCounterpartyFirm)
	pageDetail.Firm_Impl = MandatedUserKeyCounterpartyFirm_Lookup_FullName.FullName
	_, MandatedUserKeyCounterpartyCentre_Lookup_Name, _ := dao.Centre_GetByID(rD.MandatedUserKeyCounterpartyCentre)
	pageDetail.Centre_Impl = MandatedUserKeyCounterpartyCentre_Lookup_Name.Name
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	// mandate_HandlerViewImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func mandate_HandlerViewImpl(pageDetail Mandate_Page) Mandate_Page {return pageDetail}
	pageDetail = mandate_HandlerViewImpl(pageDetail)

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	ExecuteTemplate(dm.Mandate_TemplateView, w, r, pageDetail)

}

//Mandate_HandlerEdit is the handler used generate the Edit page
func Mandate_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)

	pageDetail := Mandate_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Mandate_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	_, Country_Lookup_Name, _ := dao.Country_GetByID(rD.Country)
	pageDetail.Country_Impl = Country_Lookup_Name.Name
	_, pageDetail.Country_Impl_List, _ = dao.Country_GetList()
	_, MandatedUserKeyCounterpartyFirm_Lookup_FullName, _ := dao.Firm_GetByID(rD.MandatedUserKeyCounterpartyFirm)
	pageDetail.Firm_Impl = MandatedUserKeyCounterpartyFirm_Lookup_FullName.FullName
	_, pageDetail.Firm_Impl_List, _ = dao.Firm_GetList()
	_, MandatedUserKeyCounterpartyCentre_Lookup_Name, _ := dao.Centre_GetByID(rD.MandatedUserKeyCounterpartyCentre)
	pageDetail.Centre_Impl = MandatedUserKeyCounterpartyCentre_Lookup_Name.Name
	_, pageDetail.Centre_Impl_List, _ = dao.Centre_GetList()
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	// mandate_HandlerEditImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func mandate_HandlerEditImpl(pageDetail Mandate_Page) Mandate_Page {return pageDetail}
	pageDetail = mandate_HandlerEditImpl(pageDetail)

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	ExecuteTemplate(dm.Mandate_TemplateEdit, w, r, pageDetail)

}

//Mandate_HandlerSave is the handler used process the saving of an Mandate
func Mandate_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path + r.FormValue("CompID"))

	var item dm.Mandate
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
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
	item.Country_Impl = r.FormValue(dm.Mandate_Country_Impl)
	item.Firm_Impl = r.FormValue(dm.Mandate_Firm_Impl)
	item.Centre_Impl = r.FormValue(dm.Mandate_Centre_Impl)

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	// mandate_HandlerSaveImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func mandate_HandlerSaveImpl(item dm.Mandate) dm.Mandate {return item}
	item = mandate_HandlerSaveImpl(item)

	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END

	dao.Mandate_Store(item)

	http.Redirect(w, r, Mandate_Redirect, http.StatusFound)
}

//Mandate_HandlerNew is the handler used process the creation of an Mandate
func Mandate_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Mandate_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Mandate_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
	}

	//
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
	pageDetail.MandatedUserKeyCounterpartyFirm = ""
	pageDetail.MandatedUserKeyCounterpartyCentre = ""
	pageDetail.MandatedUserKeyUserName = ""
	pageDetail.TelephoneNumber = ""
	pageDetail.EmailAddress = ""
	pageDetail.Active = ""
	pageDetail.FirstName = ""
	pageDetail.Surname = ""
	pageDetail.DateOfBirth = ""
	pageDetail.Postcode = ""
	pageDetail.NationalIDNo = ""
	pageDetail.PassportNo = ""
	pageDetail.Country = ""
	pageDetail.CountryName = ""
	pageDetail.FirmName = ""
	pageDetail.CentreName = ""
	pageDetail.Notify = ""
	pageDetail.SystemUser = ""
	pageDetail.CompID = ""
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	pageDetail.Country_Impl = ""
	_, pageDetail.Country_Impl_List, _ = dao.Country_GetList()
	pageDetail.Firm_Impl = ""
	_, pageDetail.Firm_Impl_List, _ = dao.Firm_GetList()
	pageDetail.Centre_Impl = ""
	_, pageDetail.Centre_Impl_List, _ = dao.Centre_GetList()
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
	//

	ExecuteTemplate(dm.Mandate_TemplateNew, w, r, pageDetail)

}

//Mandate_HandlerDelete is the handler used process the deletion of an Mandate
func Mandate_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Mandate_QueryString)

	dao.Mandate_Delete(searchID)

	http.Redirect(w, r, Mandate_Redirect, http.StatusFound)
}
