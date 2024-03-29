package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/cmnotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 15:01:35
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CMNotes defines the datamolde for the CMNotes object
type CMNotes struct {

NoteId       string
StreamId       string
Summary       string
Details       string
RecordState       string
CreatedBy       string
CreatedDateTime       string

}

const (
	CMNotes_Title       = "Customer Note"
	CMNotes_SQLTable    = "contactManagerNote"
	CMNotes_SQLSearchID = "noteId"
	CMNotes_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CMNotes_Template     = "CMNotes"
	CMNotes_TemplateList = "/CMNotes/CMNotes_List"
	CMNotes_TemplateView = "/CMNotes/CMNotes_View"
	CMNotes_TemplateEdit = "/CMNotes/CMNotes_Edit"
	CMNotes_TemplateNew  = "/CMNotes/CMNotes_New"
	///
	/// Handler Monitor Paths
	///
	CMNotes_Path       = "/API/CMNotes/"
	CMNotes_PathList   = "/CMNotesList/"
	CMNotes_PathView   = "/CMNotesView/"
	CMNotes_PathEdit   = "/CMNotesEdit/"
	CMNotes_PathNew    = "/CMNotesNew/"
	CMNotes_PathSave   = "/CMNotesSave/"
	CMNotes_PathDelete = "/CMNotesDelete/"
	///
	///
	/// SQL Field Definitions
	///
CMNotes_NoteId_sql   = "noteId" // NoteId is a Int
CMNotes_StreamId_sql   = "streamId" // StreamId is a Int
CMNotes_Summary_sql   = "summary" // Summary is a String
CMNotes_Details_sql   = "details" // Details is a String
CMNotes_RecordState_sql   = "recordState" // RecordState is a Int
CMNotes_CreatedBy_sql   = "createdBy" // CreatedBy is a String
CMNotes_CreatedDateTime_sql   = "createdDateTime" // CreatedDateTime is a Time

	/// Definitions End

	/// Application Field Definitions
	///
CMNotes_NoteId_scrn   = "NoteId" // NoteId is a Int
CMNotes_StreamId_scrn   = "StreamId" // StreamId is a Int
CMNotes_Summary_scrn   = "Summary" // Summary is a String
CMNotes_Details_scrn   = "Details" // Details is a String
CMNotes_RecordState_scrn   = "RecordState" // RecordState is a Int
CMNotes_CreatedBy_scrn   = "CreatedBy" // CreatedBy is a String
CMNotes_CreatedDateTime_scrn   = "CreatedDateTime" // CreatedDateTime is a Time

	/// Definitions End
)
