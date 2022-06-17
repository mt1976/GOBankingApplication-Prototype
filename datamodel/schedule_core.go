package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Schedule defines the datamolde for the Schedule object
type Schedule struct {

SYSId       string
Id       string
Name       string
Description       string
Schedule       string
Started       string
Lastrun       string
Message       string
SYSCreated       string
SYSWho       string
SYSHost       string
SYSUpdated       string
Type       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdatedBy       string
SYSUpdatedHost       string
Human       string

}

const (
	Schedule_Title       = "Scheduler"
	Schedule_SQLTable    = "scheduleStore"
	Schedule_SQLSearchID = "id"
	Schedule_QueryString = "Schedule"
	///
	/// Handler Defintions
	///
	Schedule_Template     = "Schedule"
	Schedule_TemplateList = "/Schedule/Schedule_List"
	Schedule_TemplateView = "/Schedule/Schedule_View"
	Schedule_TemplateEdit = "/Schedule/Schedule_Edit"
	Schedule_TemplateNew  = "/Schedule/Schedule_New"
	///
	/// Handler Monitor Paths
	///
	Schedule_Path       = "/API/Schedule/"
	Schedule_PathList   = "/ScheduleList/"
	Schedule_PathView   = "/ScheduleView/"
	Schedule_PathEdit   = "/ScheduleEdit/"
	Schedule_PathNew    = "/ScheduleNew/"
	Schedule_PathSave   = "/ScheduleSave/"
	Schedule_PathDelete = "/ScheduleDelete/"
	///
	///
	/// SQL Field Definitions
	///
Schedule_SYSId_sql   = "_id" // SYSId is a Int
Schedule_Id_sql   = "id" // Id is a String
Schedule_Name_sql   = "name" // Name is a String
Schedule_Description_sql   = "description" // Description is a String
Schedule_Schedule_sql   = "schedule" // Schedule is a String
Schedule_Started_sql   = "started" // Started is a String
Schedule_Lastrun_sql   = "lastrun" // Lastrun is a String
Schedule_Message_sql   = "message" // Message is a String
Schedule_SYSCreated_sql   = "_created" // SYSCreated is a String
Schedule_SYSWho_sql   = "_who" // SYSWho is a String
Schedule_SYSHost_sql   = "_host" // SYSHost is a String
Schedule_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Schedule_Type_sql   = "type" // Type is a String
Schedule_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Schedule_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Schedule_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Schedule_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
Schedule_Human_sql   = "human" // Human is a String

	/// Definitions End

	/// Application Field Definitions
	///
Schedule_SYSId_scrn   = "SYSId" // SYSId is a Int
Schedule_Id_scrn   = "Id" // Id is a String
Schedule_Name_scrn   = "Name" // Name is a String
Schedule_Description_scrn   = "Description" // Description is a String
Schedule_Schedule_scrn   = "Schedule" // Schedule is a String
Schedule_Started_scrn   = "Started" // Started is a String
Schedule_Lastrun_scrn   = "Lastrun" // Lastrun is a String
Schedule_Message_scrn   = "Message" // Message is a String
Schedule_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Schedule_SYSWho_scrn   = "SYSWho" // SYSWho is a String
Schedule_SYSHost_scrn   = "SYSHost" // SYSHost is a String
Schedule_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Schedule_Type_scrn   = "Type" // Type is a String
Schedule_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Schedule_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Schedule_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Schedule_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
Schedule_Human_scrn   = "Human" // Human is a String

	/// Definitions End
)
