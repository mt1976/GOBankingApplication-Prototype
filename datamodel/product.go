package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/product.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Product
// Endpoint Root 	  : Product
// Search QueryString : Code
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 19/11/2021 at 17:16:05
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Product struct {
	AppInternalID string  // Special field for internal use only
Code        string
Name        string
Factor        string
MaxTermPrecedence        string
InternalId        string
InternalDeleted        string
UpdatedTransactionId        string
UpdatedUserId        string
UpdatedDateTime        string
DeletedTransactionId        string
DeletedUserId        string
ChangeType        string

}

const (
	Product_Title       = "Product Groups"
	Product_SQLTable    = "sienaProduct"
	Product_SQLSearchID = "Code"
	Product_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Product_TemplateList = "Product_List"
	Product_TemplateView = "Product_View"
	Product_TemplateEdit = "Product_Edit"
	Product_TemplateNew  = "Product_New"
	///
	/// Handler Monitor Paths
	///
	Product_PathList   = "/ProductList/"
	Product_PathView   = "/ProductView/"
	Product_PathEdit   = "/ProductEdit/"
	Product_PathNew    = "/ProductNew/"
	Product_PathSave   = "/ProductSave/"
	Product_PathDelete = "/ProductDelete/"
	///
	/// SQL Field Definitions
	///
	Product_Code   = "Code" // Code is a String
	Product_Name   = "Name" // Name is a String
	Product_Factor   = "Factor" // Factor is a Float
	Product_MaxTermPrecedence   = "MaxTermPrecedence" // MaxTermPrecedence is a Bool
	Product_InternalId   = "InternalId" // InternalId is a Int
	Product_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
	Product_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Product_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	Product_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Product_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
	Product_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
	Product_ChangeType   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
