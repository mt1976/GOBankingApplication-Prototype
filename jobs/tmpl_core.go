package jobs
// ----------------------------------------------------------------
// Automatically generated  "/jobs/tmpl.go"
// ----------------------------------------------------------------
// Package              : jobs
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/mwt-go-dev/application"
	dm          "github.com/mt1976/mwt-go-dev/datamodel"
	logs        "github.com/mt1976/mwt-go-dev/logs"
		core        "github.com/mt1976/mwt-go-dev/core"

	cron        "github.com/robfig/cron/v3"
)

// TmplJob defines the job properties, name, period etc..
func Tmpl_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// Create a new job definition in the Tmpl_Impl.go file called Tmpl_JobImpl()
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Tmpl"
	j.Name = "Tmpl"
	j.Period = "10 1 * * *"
	j.Description = "Tmpl processing"
	j.Type = core.General
	j = Tmpl_Job_impl(j)
	return j
}

func Tmpl_Register(c *cron.Cron) {
	application.Schedule_Register(Tmpl_Job())
	c.AddFunc(Tmpl_Job().Period, func() { Tmpl_Run() })
}

// Tmpl_Run initiates and runs the job as per the period.
func Tmpl_Run() {
	logs.StartJob(Tmpl_Job().Name)
	message := ""
	/// CONTENT STARTS
	// Create a func in the Tmpl_Impl.go file called Tmpl_Run_impl() that returns string,error	
	message,err := Tmpl_Run_impl()
	if err != nil {
		logs.Warning(err.Error())
	}
	/// CONTENT ENDS
	application.Schedule_Update(Tmpl_Job(), message)
	logs.EndJob(Tmpl_Job().Name)
}