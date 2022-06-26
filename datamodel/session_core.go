package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Session defines the datamolde for the Session object
type Session struct {


SYSId       string
SYSId_props FieldProperties
Apptoken       string
Apptoken_props FieldProperties
Createdate       string
Createdate_props FieldProperties
Createtime       string
Createtime_props FieldProperties
Uniqueid       string
Uniqueid_props FieldProperties
Sessiontoken       string
Sessiontoken_props FieldProperties
Username       string
Username_props FieldProperties
Password       string
Password_props FieldProperties
Userip       string
Userip_props FieldProperties
Userhost       string
Userhost_props FieldProperties
Appip       string
Appip_props FieldProperties
Apphost       string
Apphost_props FieldProperties
Issued       string
Issued_props FieldProperties
Expiry       string
Expiry_props FieldProperties
Expiryraw       string
Expiryraw_props FieldProperties
Brand       string
Brand_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSWho       string
SYSWho_props FieldProperties
SYSHost       string
SYSHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
Id       string
Id_props FieldProperties
Expires       string
Expires_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
SessionRole       string
SessionRole_props FieldProperties
 // Any lookups will be added below



























}

const (
	Session_Title       = "Session"
	Session_SQLTable    = "sessionStore"
	Session_SQLSearchID = "Id"
	Session_QueryString = "SessionID"
	///
	/// Handler Defintions
	///
	Session_Template     = "Session"
	Session_TemplateList = "/Session/Session_List"
	Session_TemplateView = "/Session/Session_View"
	Session_TemplateEdit = "/Session/Session_Edit"
	Session_TemplateNew  = "/Session/Session_New"
	///
	/// Handler Monitor Paths
	///
	Session_Path       = "/API/Session/"
	Session_PathList   = "/SessionList/"
	Session_PathView   = "/SessionView/"
	Session_PathEdit   = "/SessionEdit/"
	Session_PathNew    = "/SessionNew/"
	Session_PathSave   = "/SessionSave/"
	Session_PathDelete = "/SessionDelete/"
	///
	///
	/// SQL Field Definitions
	///
Session_SYSId_sql   = "_id" // SYSId is a Int
Session_Apptoken_sql   = "Apptoken" // Apptoken is a String
Session_Createdate_sql   = "Createdate" // Createdate is a String
Session_Createtime_sql   = "Createtime" // Createtime is a String
Session_Uniqueid_sql   = "Uniqueid" // Uniqueid is a String
Session_Sessiontoken_sql   = "Sessiontoken" // Sessiontoken is a String
Session_Username_sql   = "Username" // Username is a String
Session_Password_sql   = "Password" // Password is a String
Session_Userip_sql   = "Userip" // Userip is a String
Session_Userhost_sql   = "Userhost" // Userhost is a String
Session_Appip_sql   = "Appip" // Appip is a String
Session_Apphost_sql   = "Apphost" // Apphost is a String
Session_Issued_sql   = "Issued" // Issued is a String
Session_Expiry_sql   = "Expiry" // Expiry is a String
Session_Expiryraw_sql   = "Expiryraw" // Expiryraw is a String
Session_Brand_sql   = "Brand" // Brand is a String
Session_SYSCreated_sql   = "_created" // SYSCreated is a String
Session_SYSWho_sql   = "_who" // SYSWho is a String
Session_SYSHost_sql   = "_host" // SYSHost is a String
Session_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Session_Id_sql   = "Id" // Id is a String
Session_Expires_sql   = "Expires" // Expires is a Time
Session_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Session_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Session_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Session_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
Session_SessionRole_sql   = "SessionRole" // SessionRole is a String

	/// Definitions End

	/// Application Field Definitions
	///
Session_SYSId_scrn   = "SYSId" // SYSId is a Int
Session_Apptoken_scrn   = "Apptoken" // Apptoken is a String
Session_Createdate_scrn   = "Createdate" // Createdate is a String
Session_Createtime_scrn   = "Createtime" // Createtime is a String
Session_Uniqueid_scrn   = "Uniqueid" // Uniqueid is a String
Session_Sessiontoken_scrn   = "Sessiontoken" // Sessiontoken is a String
Session_Username_scrn   = "Username" // Username is a String
Session_Password_scrn   = "Password" // Password is a String
Session_Userip_scrn   = "Userip" // Userip is a String
Session_Userhost_scrn   = "Userhost" // Userhost is a String
Session_Appip_scrn   = "Appip" // Appip is a String
Session_Apphost_scrn   = "Apphost" // Apphost is a String
Session_Issued_scrn   = "Issued" // Issued is a String
Session_Expiry_scrn   = "Expiry" // Expiry is a String
Session_Expiryraw_scrn   = "Expiryraw" // Expiryraw is a String
Session_Brand_scrn   = "Brand" // Brand is a String
Session_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Session_SYSWho_scrn   = "SYSWho" // SYSWho is a String
Session_SYSHost_scrn   = "SYSHost" // SYSHost is a String
Session_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Session_Id_scrn   = "Id" // Id is a String
Session_Expires_scrn   = "Expires" // Expires is a Time
Session_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Session_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Session_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Session_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
Session_SessionRole_scrn   = "SessionRole" // SessionRole is a String

	/// Definitions End
)
