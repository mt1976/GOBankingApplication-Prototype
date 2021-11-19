package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/sector.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Sector
// Endpoint Root 	  : Sector
// Search QueryString : Sector
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 19/11/2021 at 17:16:05
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Sector struct {
	AppInternalID string  // Special field for internal use only
Code        string
Name        string

}

const (
	Sector_Title       = "Sectors"
	Sector_SQLTable    = "sienaSector"
	Sector_SQLSearchID = "Code"
	Sector_QueryString = "Sector"
	///
	/// Handler Defintions
	///
	Sector_TemplateList = "Sector_List"
	Sector_TemplateView = "Sector_View"
	Sector_TemplateEdit = "Sector_Edit"
	Sector_TemplateNew  = "Sector_New"
	///
	/// Handler Monitor Paths
	///
	Sector_PathList   = "/SectorList/"
	Sector_PathView   = "/SectorView/"
	Sector_PathEdit   = "/SectorEdit/"
	Sector_PathNew    = "/SectorNew/"
	Sector_PathSave   = "/SectorSave/"
	Sector_PathDelete = "/SectorDelete/"
	///
	/// SQL Field Definitions
	///
	Sector_Code   = "Code" // Code is a String
	Sector_Name   = "Name" // Name is a String

	/// Definitions End
)
