package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/currency.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Currency
// Endpoint Root 	  : Currency
// Search QueryString : Code
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 18:05:22
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Currency struct {
	AppInternalID string  // Special field for internal use only
Code        string
Name        string
AmountDp        string
Country        string
CountryName        string
IntBase        string
KeydateBase        string
InterestRateTolerance        string
CheckPayTo        string
LatinAmericanSettlement        string
DefaultLayOffBookKey        string
CutOffTimeCutOffTime        string
CutOffTimeTimeZone        string
CutOffTimeDerivedDataUTCOffset        string
CutOffTimeDerivedDataHasDaylightSaving        string
CutOffTimeDerivedDataDaylightStart        string
CutOffTimeDerivedDataDaylightEnd        string
DealerInterventionQuoteTimeout        string
CutOffTimeCutOffPeriod        string
StripRateFutureExchangeCode        string
StripRateFutureCurrencyContractCurrencyIsoCode        string
StripRateFutureCurrencyContractFutureContractCode        string
OvernightFundingSpreadBid        string
OvernightFundingSpreadOffer        string

}

const (
	Currency_Title       = "Currencies"
	Currency_SQLTable    = "sienaCurrency"
	Currency_SQLSearchID = "Code"
	Currency_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Currency_TemplateList = "Currency_List"
	Currency_TemplateView = "Currency_View"
	Currency_TemplateEdit = "Currency_Edit"
	Currency_TemplateNew  = "Currency_New"
	///
	/// Handler Monitor Paths
	///
	Currency_PathList   = "/CurrencyList/"
	Currency_PathView   = "/CurrencyView/"
	Currency_PathEdit   = "/CurrencyEdit/"
	Currency_PathNew    = "/CurrencyNew/"
	Currency_PathSave   = "/CurrencySave/"
	Currency_PathDelete = "/CurrencyDelete/"
	///
	/// SQL Field Definitions
	///
	Currency_Code   = "Code" // Code is a String
	Currency_Name   = "Name" // Name is a String
	Currency_AmountDp   = "AmountDp" // AmountDp is a Int
	Currency_Country   = "Country" // Country is a String
	Currency_CountryName   = "CountryName" // CountryName is a String
	Currency_IntBase   = "IntBase" // IntBase is a String
	Currency_KeydateBase   = "KeydateBase" // KeydateBase is a String
	Currency_InterestRateTolerance   = "InterestRateTolerance" // InterestRateTolerance is a Float
	Currency_CheckPayTo   = "CheckPayTo" // CheckPayTo is a Bool
	Currency_LatinAmericanSettlement   = "LatinAmericanSettlement" // LatinAmericanSettlement is a Bool
	Currency_DefaultLayOffBookKey   = "DefaultLayOffBookKey" // DefaultLayOffBookKey is a String
	Currency_CutOffTimeCutOffTime   = "CutOffTimeCutOffTime" // CutOffTimeCutOffTime is a Time
	Currency_CutOffTimeTimeZone   = "CutOffTimeTimeZone" // CutOffTimeTimeZone is a String
	Currency_CutOffTimeDerivedDataUTCOffset   = "CutOffTimeDerivedDataUTCOffset" // CutOffTimeDerivedDataUTCOffset is a String
	Currency_CutOffTimeDerivedDataHasDaylightSaving   = "CutOffTimeDerivedDataHasDaylightSaving" // CutOffTimeDerivedDataHasDaylightSaving is a Bool
	Currency_CutOffTimeDerivedDataDaylightStart   = "CutOffTimeDerivedDataDaylightStart" // CutOffTimeDerivedDataDaylightStart is a Time
	Currency_CutOffTimeDerivedDataDaylightEnd   = "CutOffTimeDerivedDataDaylightEnd" // CutOffTimeDerivedDataDaylightEnd is a Time
	Currency_DealerInterventionQuoteTimeout   = "DealerInterventionQuoteTimeout" // DealerInterventionQuoteTimeout is a Int
	Currency_CutOffTimeCutOffPeriod   = "CutOffTimeCutOffPeriod" // CutOffTimeCutOffPeriod is a String
	Currency_StripRateFutureExchangeCode   = "StripRateFutureExchangeCode" // StripRateFutureExchangeCode is a String
	Currency_StripRateFutureCurrencyContractCurrencyIsoCode   = "StripRateFutureCurrencyContractCurrencyIsoCode" // StripRateFutureCurrencyContractCurrencyIsoCode is a String
	Currency_StripRateFutureCurrencyContractFutureContractCode   = "StripRateFutureCurrencyContractFutureContractCode" // StripRateFutureCurrencyContractFutureContractCode is a String
	Currency_OvernightFundingSpreadBid   = "OvernightFundingSpreadBid" // OvernightFundingSpreadBid is a Float
	Currency_OvernightFundingSpreadOffer   = "OvernightFundingSpreadOffer" // OvernightFundingSpreadOffer is a Float

	/// Definitions End
)
