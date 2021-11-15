package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/marketrates.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : MarketRates
// Endpoint Root 	  : MarketRates
// Search QueryString : ID
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 19:03:09
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type MarketRates struct {
	AppInternalID string  // Special field for internal use only
SYSId        string
Id        string
Bid        string
Mid        string
Offer        string
Market        string
Tenor        string
Series        string
Name        string
Source        string
Destination        string
Class        string
SYSCreated        string
SYSWho        string
SYSHost        string
Date        string
SYSUpdated        string

}

const (
	MarketRates_Title       = "Rates Store Contents"
	MarketRates_SQLTable    = "rateDataStore"
	MarketRates_SQLSearchID = "Id"
	MarketRates_QueryString = "ID"
	///
	/// Handler Defintions
	///
	MarketRates_TemplateList = "MarketRates_List"
	MarketRates_TemplateView = "MarketRates_View"
	MarketRates_TemplateEdit = "MarketRates_Edit"
	MarketRates_TemplateNew  = "MarketRates_New"
	///
	/// Handler Monitor Paths
	///
	MarketRates_PathList   = "/MarketRatesList/"
	MarketRates_PathView   = "/MarketRatesView/"
	MarketRates_PathEdit   = "/MarketRatesEdit/"
	MarketRates_PathNew    = "/MarketRatesNew/"
	MarketRates_PathSave   = "/MarketRatesSave/"
	MarketRates_PathDelete = "/MarketRatesDelete/"
	///
	/// SQL Field Definitions
	///
	MarketRates_SYSId   = "_id" // SYSId is a Int
	MarketRates_Id   = "id" // Id is a String
	MarketRates_Bid   = "bid" // Bid is a String
	MarketRates_Mid   = "mid" // Mid is a String
	MarketRates_Offer   = "offer" // Offer is a String
	MarketRates_Market   = "market" // Market is a String
	MarketRates_Tenor   = "tenor" // Tenor is a String
	MarketRates_Series   = "series" // Series is a String
	MarketRates_Name   = "name" // Name is a String
	MarketRates_Source   = "source" // Source is a String
	MarketRates_Destination   = "destination" // Destination is a String
	MarketRates_Class   = "class" // Class is a String
	MarketRates_SYSCreated   = "_created" // SYSCreated is a String
	MarketRates_SYSWho   = "_who" // SYSWho is a String
	MarketRates_SYSHost   = "_host" // SYSHost is a String
	MarketRates_Date   = "date" // Date is a String
	MarketRates_SYSUpdated   = "_updated" // SYSUpdated is a String

	/// Definitions End
)
