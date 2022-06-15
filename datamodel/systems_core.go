package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/systems.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Systems (systems)
// Endpoint 	        : Systems (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Systems defines the datamolde for the Systems object
type Systems struct {

SYSId       string
Id       string
Name       string
Staticin       string
Staticout       string
Txnin       string
Txnout       string
Fundscheckin       string
Fundscheckout       string
SYSCreated       string
SYSWho       string
SYSHost       string
SYSUpdated       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdatedBy       string
SYSUpdatedHost       string
SWIFTin       string
SWIFTout       string

}

const (
	Systems_Title       = "Connected System"
	Systems_SQLTable    = "systemStore"
	Systems_SQLSearchID = "Id"
	Systems_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Systems_Template     = "Systems"
	Systems_TemplateList = "Systems_List"
	Systems_TemplateView = "Systems_View"
	Systems_TemplateEdit = "Systems_Edit"
	Systems_TemplateNew  = "Systems_New"
	///
	/// Handler Monitor Paths
	///
	Systems_Path       = "/API/Systems/"
	Systems_PathList   = "/SystemsList/"
	Systems_PathView   = "/SystemsView/"
	Systems_PathEdit   = "/SystemsEdit/"
	Systems_PathNew    = "/SystemsNew/"
	Systems_PathSave   = "/SystemsSave/"
	Systems_PathDelete = "/SystemsDelete/"
	///
	/// SQL Field Definitions
	///
Systems_SYSId   = "_id" // SYSId is a Int
Systems_Id   = "Id" // Id is a String
Systems_Name   = "Name" // Name is a String
Systems_Staticin   = "Staticin" // Staticin is a String
Systems_Staticout   = "Staticout" // Staticout is a String
Systems_Txnin   = "Txnin" // Txnin is a String
Systems_Txnout   = "Txnout" // Txnout is a String
Systems_Fundscheckin   = "Fundscheckin" // Fundscheckin is a String
Systems_Fundscheckout   = "Fundscheckout" // Fundscheckout is a String
Systems_SYSCreated   = "_created" // SYSCreated is a String
Systems_SYSWho   = "_who" // SYSWho is a String
Systems_SYSHost   = "_host" // SYSHost is a String
Systems_SYSUpdated   = "_updated" // SYSUpdated is a String
Systems_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
Systems_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
Systems_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
Systems_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
Systems_SWIFTin   = "SWIFTin" // SWIFTin is a String
Systems_SWIFTout   = "SWIFTout" // SWIFTout is a String

	/// Definitions End
)
