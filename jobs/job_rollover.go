package jobs

import (
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/siena"
	cron "github.com/robfig/cron/v3"
)

func Rollover_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "ROLLOVER"
	j.Name = "ROLLOVER"
	j.Period = "10 1 * * *"
	j.Description = "Siena System Rollover Refresh"
	j.Type = core.Monitor
	return j
}

func Rollover_Register(c *cron.Cron) {
	//application.RegisterSchedule(Rollover_Job().ID, Rollover_Job().Name, Rollover_Job().Description, Rollover_Job().Period, Rollover_Job().Type)
	application.RegisterSchedule(Rollover_Job())
	c.AddFunc(Rollover_Job().Period, func() { Rollover_Run() })
}

// RunJobRollover is a Rollover function
func Rollover_Run() {
	logStart(Rollover_Job().Name)
	/// CONTENT STARTS
	core.Log_uptime()

	core.SienaDB = core.GlobalsDatabasePoke(core.SienaDB, core.SienaPropertiesDB)
	oldSysDate := core.SienaSystemDate
	_, tempDate, _ := siena.GetBusinessDate(core.SienaDB)
	core.SienaSystemDate = tempDate

	log.Printf("Old System Date: %v\n", oldSysDate)
	log.Printf("New System Date: %v\n", core.SienaSystemDate)

	message := fmt.Sprintf("Rolled from %v to %v", oldSysDate.Internal, core.SienaSystemDate.Internal)

	//application.UpdateSchedule("SRO", core.Monitor, message)

	/// CONTENT ENDS
	application.UpdateSchedule(Rollover_Job(), message)
	logEnd(Rollover_Job().Name)
}
