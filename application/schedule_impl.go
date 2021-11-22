package application

// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 12:24:49
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/lnquy/cron"
	hcron "github.com/lnquy/cron"
	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//Schedule_Publish annouces the endpoints available for this object
func schedule_PublishImpl(mux http.ServeMux) http.ServeMux {
	//No Overides this time
	return mux
}

func Schedule_Register(thisJob dm.JobDefinition) {
	var s dm.Schedule
	//s.Id = thisJob.ID + core.IDSep + thisJob.Type
	s.Name = thisJob.Name
	s.Description = thisJob.Description
	s.Schedule = thisJob.Period
	s.Started = time.Now().Format(core.DATETIMEFORMATUSER)
	s.Lastrun = ""
	s.Message = ""
	s.SYSCreated = time.Now().Format(core.DATETIMEFORMATUSER)
	currentUserID, _ := user.Current()
	host, _ := os.Hostname()
	s.SYSWho = currentUserID.Name
	s.SYSHost = host
	s.SYSUpdated = time.Now().Format(core.DATETIMEFORMATUSER)
	s.Type = thisJob.Type
	//log.Println("STORE", s)
	s.Id = dao.Schedule_NewID(s)
	registerIt := true
	s.Human = Schedule_GetCronHuman(s.Schedule)

	if core.IsChildInstance {
		if s.Type == core.Aquirer {
			registerIt = false
		}
	}
	if registerIt {
		dao.Schedule_Store(s)
		//desc := GetCronScheduleHuman(s.Schedule)

		icon := core.Character_Gears
		switch s.Type {
		case core.Aquirer:
			icon = core.Character_Aquirer
		case core.Monitor:
			icon = core.Character_Monitor + " "
		case core.Dispatcher:
			icon = core.Character_Dispatcher
		}
		op := fmt.Sprintf("%s %-18s %-18s %q", icon, s.Name, s.Schedule, Schedule_GetCronHuman(s.Schedule))
		logs.Schedule(op)
	}
}

func Schedule_Update(thisJob dm.JobDefinition, message string) {
	scheduleID := dao.Schedule_NewID(dm.Schedule{Name: thisJob.Name, Type: thisJob.Type})
	if len(scheduleID) > 1 {
		_, s, _ := dao.Schedule_GetByID(scheduleID)
		if len(s.Name) > 0 {
			s.Lastrun = time.Now().Format(core.DATETIMEFORMATUSER)
			s.Message = message
			thisMess := fmt.Sprintf("Ran Job - %-11s %-20s %q", thisJob.Type, s.Name, message)
			logs.Schedule(thisMess)
			dao.Schedule_Store(s)
		} else {
			thisMess := fmt.Sprintf("Update Schedule Schedule with '%s','%s','%s' ScheduleID = '%s'", thisJob.ID, thisJob.ID, message, scheduleID)
			logs.Schedule(thisMess)
			//spew.Dump(s)
		}
	} else {
		thisMess := fmt.Sprintf("Update Schedule Called with '%s','%s','%s'", thisJob.ID, thisJob.Type, message)
		logs.Schedule(thisMess)
	}
}

func Schedule_GetCronHuman(in string) string {
	desc := ""
	if len(in) != 0 {
		exprDesc, err := hcron.NewDescriptor(
			cron.Use24HourTimeFormat(true),
			cron.DayOfWeekStartsAtOne(false),
			cron.Verbose(true),
			cron.SetLogger(log.New(os.Stdout, "cron: ", 0)),
			cron.SetLocales(hcron.Locale_en),
		)
		if err != nil {
			logs.Error("failed to create CRON expression descriptor:", err)
		}
		desc, err = exprDesc.ToDescription(in, hcron.Locale_en)
		if err != nil {
			logs.Error("failed to convert CRON expression to human readable description:", err)
		}
	}

	return desc
}

func schedule_HandlerViewImpl(pageDetail Schedule_Page) Schedule_Page { return pageDetail }
func schedule_HandlerEditImpl(pageDetail Schedule_Page) Schedule_Page { return pageDetail }
func schedule_HandlerSaveImpl(item dm.Schedule) dm.Schedule {
	return item
}