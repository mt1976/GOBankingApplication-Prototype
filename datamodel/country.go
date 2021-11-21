package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/country.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:01
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Country struct {
	AppInternalID string  // Special field for internal use only
Code        string
Name        string
ShortCode        string
EU_EEA        string
HolidaysWeekend        string

}

const (
	Country_Title       = "Countries"
	Country_SQLTable    = "sienaCountry"
	Country_SQLSearchID = "Code"
	Country_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Country_TemplateList = "Country_List"
	Country_TemplateView = "Country_View"
	Country_TemplateEdit = "Country_Edit"
	Country_TemplateNew  = "Country_New"
	///
	/// Handler Monitor Paths
	///
	Country_PathList   = "/CountryList/"
	Country_PathView   = "/CountryView/"
	Country_PathEdit   = "/CountryEdit/"
	Country_PathNew    = "/CountryNew/"
	Country_PathSave   = "/CountrySave/"
	Country_PathDelete = "/CountryDelete/"
	///
	/// SQL Field Definitions
	///
	Country_Code   = "Code" // Code is a String
	Country_Name   = "Name" // Name is a String
	Country_ShortCode   = "ShortCode" // ShortCode is a String
	Country_EU_EEA   = "EU_EEA" // EU_EEA is a Bool
	Country_HolidaysWeekend   = "HolidaysWeekend" // HolidaysWeekend is a String

	/// Definitions End
)
