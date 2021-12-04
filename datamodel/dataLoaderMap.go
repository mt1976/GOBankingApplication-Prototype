package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dataloadermap.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderMap (dataloadermap)
// Endpoint 	        : DataLoaderMap (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:45
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DataLoaderMap struct {

SYSId        string
Id        string
Name        string
Position        string
Loader        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Int_position        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Loader_Impl        string
LoaderDescription_Impl        string
LoaderType_Impl        string

}

const (
	DataLoaderMap_Title       = "Data Loader Data Map"
	DataLoaderMap_SQLTable    = "loaderMapStore"
	DataLoaderMap_SQLSearchID = "Id"
	DataLoaderMap_QueryString = "Id"
	///
	/// Handler Defintions
	///
	DataLoaderMap_TemplateList = "DataLoaderMap_List"
	DataLoaderMap_TemplateView = "DataLoaderMap_View"
	DataLoaderMap_TemplateEdit = "DataLoaderMap_Edit"
	DataLoaderMap_TemplateNew  = "DataLoaderMap_New"
	///
	/// Handler Monitor Paths
	///
	DataLoaderMap_PathList   = "/DataLoaderMapList/"
	DataLoaderMap_PathView   = "/DataLoaderMapView/"
	DataLoaderMap_PathEdit   = "/DataLoaderMapEdit/"
	DataLoaderMap_PathNew    = "/DataLoaderMapNew/"
	DataLoaderMap_PathSave   = "/DataLoaderMapSave/"
	DataLoaderMap_PathDelete = "/DataLoaderMapDelete/"
	///
	/// SQL Field Definitions
	///
	DataLoaderMap_SYSId   = "_id" // SYSId is a Int
	DataLoaderMap_Id   = "id" // Id is a String
	DataLoaderMap_Name   = "name" // Name is a String
	DataLoaderMap_Position   = "position" // Position is a String
	DataLoaderMap_Loader   = "loader" // Loader is a String
	DataLoaderMap_SYSCreated   = "_created" // SYSCreated is a String
	DataLoaderMap_SYSWho   = "_who" // SYSWho is a String
	DataLoaderMap_SYSHost   = "_host" // SYSHost is a String
	DataLoaderMap_SYSUpdated   = "_updated" // SYSUpdated is a String
	DataLoaderMap_Int_position   = "int_position" // Int_position is a Int
	DataLoaderMap_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	DataLoaderMap_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	DataLoaderMap_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	DataLoaderMap_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	DataLoaderMap_Loader_Impl   = "Loader_Impl" // Loader_Impl is a String
	DataLoaderMap_LoaderDescription_Impl   = "LoaderDescription_Impl" // LoaderDescription_Impl is a String
	DataLoaderMap_LoaderType_Impl   = "LoaderType_Impl" // LoaderType_Impl is a String

	/// Definitions End
)
