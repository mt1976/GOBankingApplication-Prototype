package adaptor

import (
	"strings"

	"github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/customeronboarding.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CustomerOnboarding (customeronboarding)
// Endpoint 	        : CustomerOnboarding (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 10:59:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in customeronboarding_impl.go

func CustomerOnboarding_GetList_impl() (int, []dm.CustomerOnboarding, error) { return 0, nil, nil }
func CustomerOnboarding_GetByID_impl(id string) (int, dm.CustomerOnboarding, error) {
	return 0, dm.CustomerOnboarding{}, nil
}

func CustomerOnboarding_NewID_impl(rec dm.CustomerOnboarding) string { return rec.ID }
func CustomerOnboarding_Delete_impl(id string) error                 { return nil }

func CustomerOnboarding_Update_impl(id string, rec dm.CustomerOnboarding, usr string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\customeronboarding_impl.go

func CustomerOnboarding_Simulator_ProcessResponse_impl(filename string) error {
	logs.Success("CustomerOnboarding_Simulator_ProcessResponse:" + filename)
	// tokenise the filename, get last element
	tokens := strings.Split(filename, "/")
	last := tokens[len(tokens)-1]
	core.Notification_Normal("New Customer Onboarding Tool Message Detected", last)
	return customeronboarding_Simulator_ProcessResponse_impl(tokens, last)
}

func customeronboarding_Simulator_ProcessResponse_impl(tokens []string, latestToken string) error {
	return nil
}
