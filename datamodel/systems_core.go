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
// Date & Time		    : 17/06/2022 at 18:38:14
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
	Systems_TemplateList = "/Systems/Systems_List"
	Systems_TemplateView = "/Systems/Systems_View"
	Systems_TemplateEdit = "/Systems/Systems_Edit"
	Systems_TemplateNew  = "/Systems/Systems_New"
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
	///
	/// SQL Field Definitions
	///
Systems_SYSId_sql   = "_id" // SYSId is a Int
Systems_Id_sql   = "Id" // Id is a String
Systems_Name_sql   = "Name" // Name is a String
Systems_Staticin_sql   = "Staticin" // Staticin is a String
Systems_Staticout_sql   = "Staticout" // Staticout is a String
Systems_Txnin_sql   = "Txnin" // Txnin is a String
Systems_Txnout_sql   = "Txnout" // Txnout is a String
Systems_Fundscheckin_sql   = "Fundscheckin" // Fundscheckin is a String
Systems_Fundscheckout_sql   = "Fundscheckout" // Fundscheckout is a String
Systems_SYSCreated_sql   = "_created" // SYSCreated is a String
Systems_SYSWho_sql   = "_who" // SYSWho is a String
Systems_SYSHost_sql   = "_host" // SYSHost is a String
Systems_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Systems_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Systems_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Systems_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Systems_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
Systems_SWIFTin_sql   = "SWIFTin" // SWIFTin is a String
Systems_SWIFTout_sql   = "SWIFTout" // SWIFTout is a String

	/// Definitions End

	/// Application Field Definitions
	///
Systems_SYSId_scrn   = "SYSId" // SYSId is a Int
Systems_Id_scrn   = "Id" // Id is a String
Systems_Name_scrn   = "Name" // Name is a String
Systems_Staticin_scrn   = "Staticin" // Staticin is a String
Systems_Staticout_scrn   = "Staticout" // Staticout is a String
Systems_Txnin_scrn   = "Txnin" // Txnin is a String
Systems_Txnout_scrn   = "Txnout" // Txnout is a String
Systems_Fundscheckin_scrn   = "Fundscheckin" // Fundscheckin is a String
Systems_Fundscheckout_scrn   = "Fundscheckout" // Fundscheckout is a String
Systems_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Systems_SYSWho_scrn   = "SYSWho" // SYSWho is a String
Systems_SYSHost_scrn   = "SYSHost" // SYSHost is a String
Systems_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Systems_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Systems_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Systems_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Systems_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
Systems_SWIFTin_scrn   = "SWIFTin" // SWIFTin is a String
Systems_SWIFTout_scrn   = "SWIFTout" // SWIFTout is a String

	/// Definitions End
)
