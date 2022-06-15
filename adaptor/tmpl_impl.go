package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/tmpl.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/06/2022 at 17:27:40
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in tmpl_impl.go

// If there are fields below, create the methods in adaptor\tmpl_impl.go

func Tmpl_ExtraField_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	return fieldval, nil
}

func Tmpl_ExtraField3_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	return fieldval, nil
}

func Tmpl_ExtraField_OnFetch_impl(rec dm.Tmpl) string { return "" }

func Tmpl_ExtraField3_OnFetch_impl(rec dm.Tmpl) string { return "" }
