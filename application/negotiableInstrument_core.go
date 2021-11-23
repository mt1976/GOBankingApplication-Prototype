package application
// ----------------------------------------------------------------
// Automatically generated  "/application/negotiableinstrument.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:34:42
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

//negotiableinstrument_PageList provides the information for the template for a list of NegotiableInstruments
type NegotiableInstrument_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.NegotiableInstrument
}

//negotiableinstrument_Page provides the information for the template for an individual NegotiableInstrument
type NegotiableInstrument_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		LongName string
		Isin string
		Tidm string
		Sedol string
		IssueDate string
		MaturityDate string
		CouponValue string
		CouponType string
		Segment string
		Sector string
		CodeConventionCalculateAccrual string
		MinimumDenomination string
		DenominationCurrency string
		TradingCurrency string
		Type string
		FlatYield string
		PaymentCouponDate string
		PeriodOfCoupon string
		ExCouponDate string
		DateOfIndexInflation string
		UnitOfQuotation string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Issuer string
		IssueAmount string
		RunningYield string
		LEI string
		CUSIP string
		SYSUpdatedHost string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
}

const (
	NegotiableInstrument_Redirect = dm.NegotiableInstrument_PathList
)

//NegotiableInstrument_Publish annouces the endpoints available for this object
func NegotiableInstrument_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.NegotiableInstrument_PathList, NegotiableInstrument_HandlerList)
	mux.HandleFunc(dm.NegotiableInstrument_PathView, NegotiableInstrument_HandlerView)
	
	
	
	
	logs.Publish("Application", dm.NegotiableInstrument_Title)
	
}

//NegotiableInstrument_HandlerList is the handler for the list page
func NegotiableInstrument_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.NegotiableInstrument
	noItems, returnList, _ := dao.NegotiableInstrument_GetList()


	pageDetail := NegotiableInstrument_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.NegotiableInstrument_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.NegotiableInstrument_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//NegotiableInstrument_HandlerView is the handler used to View a page
func NegotiableInstrument_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)
	_, rD, _ := dao.NegotiableInstrument_GetByID(searchID)

	pageDetail := NegotiableInstrument_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.NegotiableInstrument_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.LongName = rD.LongName
pageDetail.Isin = rD.Isin
pageDetail.Tidm = rD.Tidm
pageDetail.Sedol = rD.Sedol
pageDetail.IssueDate = rD.IssueDate
pageDetail.MaturityDate = rD.MaturityDate
pageDetail.CouponValue = rD.CouponValue
pageDetail.CouponType = rD.CouponType
pageDetail.Segment = rD.Segment
pageDetail.Sector = rD.Sector
pageDetail.CodeConventionCalculateAccrual = rD.CodeConventionCalculateAccrual
pageDetail.MinimumDenomination = rD.MinimumDenomination
pageDetail.DenominationCurrency = rD.DenominationCurrency
pageDetail.TradingCurrency = rD.TradingCurrency
pageDetail.Type = rD.Type
pageDetail.FlatYield = rD.FlatYield
pageDetail.PaymentCouponDate = rD.PaymentCouponDate
pageDetail.PeriodOfCoupon = rD.PeriodOfCoupon
pageDetail.ExCouponDate = rD.ExCouponDate
pageDetail.DateOfIndexInflation = rD.DateOfIndexInflation
pageDetail.UnitOfQuotation = rD.UnitOfQuotation
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Issuer = rD.Issuer
pageDetail.IssueAmount = rD.IssueAmount
pageDetail.RunningYield = rD.RunningYield
pageDetail.LEI = rD.LEI
pageDetail.CUSIP = rD.CUSIP
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.NegotiableInstrument_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//NegotiableInstrument_HandlerEdit is the handler used generate the Edit page
func NegotiableInstrument_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)
	_, rD, _ := dao.NegotiableInstrument_GetByID(searchID)
	
	pageDetail := NegotiableInstrument_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.NegotiableInstrument_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.LongName = rD.LongName
pageDetail.Isin = rD.Isin
pageDetail.Tidm = rD.Tidm
pageDetail.Sedol = rD.Sedol
pageDetail.IssueDate = rD.IssueDate
pageDetail.MaturityDate = rD.MaturityDate
pageDetail.CouponValue = rD.CouponValue
pageDetail.CouponType = rD.CouponType
pageDetail.Segment = rD.Segment
pageDetail.Sector = rD.Sector
pageDetail.CodeConventionCalculateAccrual = rD.CodeConventionCalculateAccrual
pageDetail.MinimumDenomination = rD.MinimumDenomination
pageDetail.DenominationCurrency = rD.DenominationCurrency
pageDetail.TradingCurrency = rD.TradingCurrency
pageDetail.Type = rD.Type
pageDetail.FlatYield = rD.FlatYield
pageDetail.PaymentCouponDate = rD.PaymentCouponDate
pageDetail.PeriodOfCoupon = rD.PeriodOfCoupon
pageDetail.ExCouponDate = rD.ExCouponDate
pageDetail.DateOfIndexInflation = rD.DateOfIndexInflation
pageDetail.UnitOfQuotation = rD.UnitOfQuotation
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Issuer = rD.Issuer
pageDetail.IssueAmount = rD.IssueAmount
pageDetail.RunningYield = rD.RunningYield
pageDetail.LEI = rD.LEI
pageDetail.CUSIP = rD.CUSIP
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.NegotiableInstrument_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//NegotiableInstrument_HandlerSave is the handler used process the saving of an NegotiableInstrument
func NegotiableInstrument_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.NegotiableInstrument

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.NegotiableInstrument_SYSId)
		item.Id = r.FormValue(dm.NegotiableInstrument_Id)
		item.LongName = r.FormValue(dm.NegotiableInstrument_LongName)
		item.Isin = r.FormValue(dm.NegotiableInstrument_Isin)
		item.Tidm = r.FormValue(dm.NegotiableInstrument_Tidm)
		item.Sedol = r.FormValue(dm.NegotiableInstrument_Sedol)
		item.IssueDate = r.FormValue(dm.NegotiableInstrument_IssueDate)
		item.MaturityDate = r.FormValue(dm.NegotiableInstrument_MaturityDate)
		item.CouponValue = r.FormValue(dm.NegotiableInstrument_CouponValue)
		item.CouponType = r.FormValue(dm.NegotiableInstrument_CouponType)
		item.Segment = r.FormValue(dm.NegotiableInstrument_Segment)
		item.Sector = r.FormValue(dm.NegotiableInstrument_Sector)
		item.CodeConventionCalculateAccrual = r.FormValue(dm.NegotiableInstrument_CodeConventionCalculateAccrual)
		item.MinimumDenomination = r.FormValue(dm.NegotiableInstrument_MinimumDenomination)
		item.DenominationCurrency = r.FormValue(dm.NegotiableInstrument_DenominationCurrency)
		item.TradingCurrency = r.FormValue(dm.NegotiableInstrument_TradingCurrency)
		item.Type = r.FormValue(dm.NegotiableInstrument_Type)
		item.FlatYield = r.FormValue(dm.NegotiableInstrument_FlatYield)
		item.PaymentCouponDate = r.FormValue(dm.NegotiableInstrument_PaymentCouponDate)
		item.PeriodOfCoupon = r.FormValue(dm.NegotiableInstrument_PeriodOfCoupon)
		item.ExCouponDate = r.FormValue(dm.NegotiableInstrument_ExCouponDate)
		item.DateOfIndexInflation = r.FormValue(dm.NegotiableInstrument_DateOfIndexInflation)
		item.UnitOfQuotation = r.FormValue(dm.NegotiableInstrument_UnitOfQuotation)
		item.SYSCreated = r.FormValue(dm.NegotiableInstrument_SYSCreated)
		item.SYSWho = r.FormValue(dm.NegotiableInstrument_SYSWho)
		item.SYSHost = r.FormValue(dm.NegotiableInstrument_SYSHost)
		item.SYSUpdated = r.FormValue(dm.NegotiableInstrument_SYSUpdated)
		item.Issuer = r.FormValue(dm.NegotiableInstrument_Issuer)
		item.IssueAmount = r.FormValue(dm.NegotiableInstrument_IssueAmount)
		item.RunningYield = r.FormValue(dm.NegotiableInstrument_RunningYield)
		item.LEI = r.FormValue(dm.NegotiableInstrument_LEI)
		item.CUSIP = r.FormValue(dm.NegotiableInstrument_CUSIP)
		item.SYSUpdatedHost = r.FormValue(dm.NegotiableInstrument_SYSUpdatedHost)
		item.SYSCreatedBy = r.FormValue(dm.NegotiableInstrument_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.NegotiableInstrument_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.NegotiableInstrument_SYSUpdatedBy)
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	dao.NegotiableInstrument_Store(item)	

	http.Redirect(w, r, NegotiableInstrument_Redirect, http.StatusFound)
}

//NegotiableInstrument_HandlerNew is the handler used process the creation of an NegotiableInstrument
func NegotiableInstrument_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := NegotiableInstrument_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.NegotiableInstrument_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.LongName = ""
pageDetail.Isin = ""
pageDetail.Tidm = ""
pageDetail.Sedol = ""
pageDetail.IssueDate = ""
pageDetail.MaturityDate = ""
pageDetail.CouponValue = ""
pageDetail.CouponType = ""
pageDetail.Segment = ""
pageDetail.Sector = ""
pageDetail.CodeConventionCalculateAccrual = ""
pageDetail.MinimumDenomination = ""
pageDetail.DenominationCurrency = ""
pageDetail.TradingCurrency = ""
pageDetail.Type = ""
pageDetail.FlatYield = ""
pageDetail.PaymentCouponDate = ""
pageDetail.PeriodOfCoupon = ""
pageDetail.ExCouponDate = ""
pageDetail.DateOfIndexInflation = ""
pageDetail.UnitOfQuotation = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Issuer = ""
pageDetail.IssueAmount = ""
pageDetail.RunningYield = ""
pageDetail.LEI = ""
pageDetail.CUSIP = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.NegotiableInstrument_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//NegotiableInstrument_HandlerDelete is the handler used process the deletion of an NegotiableInstrument
func NegotiableInstrument_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)

	dao.NegotiableInstrument_Delete(searchID)	

	http.Redirect(w, r, NegotiableInstrument_Redirect, http.StatusFound)
}