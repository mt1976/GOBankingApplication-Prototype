package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/cache.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Cache (cache)
// Endpoint 	        : Cache (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Cache defines the datamolde for the Cache object
type Cache struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
Object       string
Object_props FieldProperties
Field       string
Field_props FieldProperties
Value       string
Value_props FieldProperties
Expiry       string
Expiry_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSWho       string
SYSWho_props FieldProperties
SYSHost       string
SYSHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
Source       string
Source_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
 // Any lookups will be added below















}

const (
	Cache_Title       = "Cache Content"
	Cache_SQLTable    = "cacheStore"
	Cache_SQLSearchID = "Id"
	Cache_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Cache_Template     = "Cache"
	Cache_TemplateList = "/Cache/Cache_List"
	Cache_TemplateView = "/Cache/Cache_View"
	Cache_TemplateEdit = "/Cache/Cache_Edit"
	Cache_TemplateNew  = "/Cache/Cache_New"
	///
	/// Handler Monitor Paths
	///
	Cache_Path       = "/API/Cache/"
	Cache_PathList   = "/CacheList/"
	Cache_PathView   = "/CacheView/"
	Cache_PathEdit   = "/CacheEdit/"
	Cache_PathNew    = "/CacheNew/"
	Cache_PathSave   = "/CacheSave/"
	Cache_PathDelete = "/CacheDelete/"
	///
	//Cache_Redirect provides a page to return to aftern an action
	Cache_Redirect = Cache_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Cache_SYSId_sql   = "_id" // SYSId is a Int
	Cache_Id_sql   = "id" // Id is a String
	Cache_Object_sql   = "object" // Object is a String
	Cache_Field_sql   = "field" // Field is a String
	Cache_Value_sql   = "value" // Value is a String
	Cache_Expiry_sql   = "expiry" // Expiry is a String
	Cache_SYSCreated_sql   = "_created" // SYSCreated is a String
	Cache_SYSWho_sql   = "_who" // SYSWho is a String
	Cache_SYSHost_sql   = "_host" // SYSHost is a String
	Cache_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	Cache_Source_sql   = "source" // Source is a String
	Cache_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	Cache_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	Cache_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
	Cache_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Cache_SYSId_scrn   = "SYSId" // SYSId is a Int
	Cache_Id_scrn   = "Id" // Id is a String
	Cache_Object_scrn   = "Object" // Object is a String
	Cache_Field_scrn   = "Field" // Field is a String
	Cache_Value_scrn   = "Value" // Value is a String
	Cache_Expiry_scrn   = "Expiry" // Expiry is a String
	Cache_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	Cache_SYSWho_scrn   = "SYSWho" // SYSWho is a String
	Cache_SYSHost_scrn   = "SYSHost" // SYSHost is a String
	Cache_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	Cache_Source_scrn   = "Source" // Source is a String
	Cache_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	Cache_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	Cache_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
	Cache_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
	///
)

//cache_PageList provides the information for the template for a list of Caches
type Cache_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Cache
	Context	 appContext
}

//cache_Page provides the information for the template for an individual Cache
type Cache_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	Object         string
	Object_props     FieldProperties
	Field         string
	Field_props     FieldProperties
	Value         string
	Value_props     FieldProperties
	Expiry         string
	Expiry_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSWho         string
	SYSWho_props     FieldProperties
	SYSHost         string
	SYSHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	Source         string
	Source_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}