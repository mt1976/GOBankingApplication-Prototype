package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/mbndr/figlet4go"
)

//CONST_CONFIG_FILE is cheese
const (
	CONST_CONFIG_FILE  = "config/properties.cfg"
	cSQL_CONFIG        = "config/mssql.cfg"
	cSIENACONFIG       = "config/siena.cfg"
	wctEpochDateFormat = "20060102T150405"
	sienaDateFormat    = "2006-01-02"
)

var gSessionToken = ""
var gUUID = "authorAdjust"
var gSecurityViolation = ""
var gDB *sql.DB

//HomePage ...
type HomePage struct {
	Title          string
	Body           string
	RequestPath    string
	ResponsePath   string
	ProcessedPath  string
	NoResponses    int
	Responses      []WctResponsePayload
	NoServices     int
	Services       string
	ServiceCatalog []ServiceCatalogItem
	Description    string
	ResponseType   string
}

func main() {

	ascii := figlet4go.NewAsciiRender()

	wctProperties := getProperties(CONST_CONFIG_FILE)

	// The underscore would be an error
	renderStr, _ := ascii.Render("MWT GO PROTO")

	fmt.Print(renderStr)

	fmt.Println("PATHS")
	fmt.Println("Delivery   :", wctProperties["deliverpath"])
	fmt.Println("Response   :", wctProperties["receivepath"])
	fmt.Println("Processed  :", wctProperties["processedpath"])
	fmt.Println("")
	fmt.Println("APPLICATION")
	fmt.Println("Name       :", wctProperties["appname"])
	fmt.Println("Msg Format :", wctProperties["responseformat"])
	fmt.Println("Licence    :", wctProperties["licname"])
	fmt.Println("Lic URL    :", wctProperties["liclink"])
	fmt.Println("")
	fmt.Println("RELEASE")
	fmt.Println("Release ID :", wctProperties["releaseid"])
	fmt.Println("Level      :", wctProperties["releaselevel"])
	fmt.Println("Number     :", wctProperties["releasenumber"])
	fmt.Println("")

	// Test Connection

	//
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", valLoginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/home", homePageHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/favicon-32x32.png", favicon32Handler)
	http.HandleFunc("/site.webmanifest", faviconManifestHandler)
	http.HandleFunc("/favicon-16x16.png", favicon16Handler)
	http.HandleFunc("/browserconfig.xml", faviconBrowserConfigHandler)
	http.HandleFunc("/listResponses/", listResponsesHandler)
	http.HandleFunc("/previewRequest/", previewRequestHandler)
	http.HandleFunc("/executeRequest/", executeRequestHandler)
	http.HandleFunc("/viewResponse/", viewResponseHandler)
	http.HandleFunc("/deleteResponse/", deleteResponseHandler)
	http.HandleFunc("/clearQueues/", clearQueuesHandler)
	http.HandleFunc("/viewSrvEnvironment/", viewSrvEnvironmentHandler)
	http.HandleFunc("/listSrvConfiguration/", listSrvConfigurationHandler)
	http.HandleFunc("/viewSrvConfiguration/", viewSrvConfigurationHandler)
	http.HandleFunc("/editSrvConfiguration/", editSrvConfigurationHandler)
	http.HandleFunc("/saveSrvConfiguration/", saveSrvConfigurationHandler)
	http.HandleFunc("/viewAppConfiguration/", viewAppConfigurationHandler)
	http.HandleFunc("/listSvcDataMap/", listSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMap/", viewSvcDataMapHandler)
	http.HandleFunc("/editSvcDataMap/", editSvcDataMapHandler)
	http.HandleFunc("/saveSvcDataMap/", saveSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMapXML/", viewSvcDataMapXMLHandler)
	http.HandleFunc("/editSvcDataMapXML/", editSvcDataMapXMLHandler)
	http.HandleFunc("/saveSvcDataMapXML/", saveSvcDataMapXMLHandler)
	http.HandleFunc("/newSvcDataMap/", newSvcDataMapHandler)
	http.HandleFunc("/genSvcDataMap/", genSvcDataMapHandler)
	http.HandleFunc("/deleteSvcDataMap/", deleteSvcDataMapHandler)
	http.HandleFunc("/listSienaCountry/", listSienaCountryHandler)
	http.HandleFunc("/viewSienaCountry/", viewSienaCountryHandler)
	http.HandleFunc("/editSienaCountry/", editSienaCountryHandler)
	http.HandleFunc("/saveSienaCountry/", saveSienaCountryHandler)
	http.HandleFunc("/newSienaCountry/", newSienaCountryHandler)

	http.HandleFunc("/listSienaSector/", listSienaSectorHandler)
	http.HandleFunc("/viewSienaSector/", viewSienaSectorHandler)
	http.HandleFunc("/editSienaSector/", editSienaSectorHandler)
	http.HandleFunc("/saveSienaSector/", saveSienaSectorHandler)
	http.HandleFunc("/newSienaSector/", newSienaSectorHandler)

	http.HandleFunc("/listSienaFirm/", listSienaFirmHandler)
	http.HandleFunc("/viewSienaFirm/", viewSienaFirmHandler)
	http.HandleFunc("/editSienaFirm/", editSienaFirmHandler)
	http.HandleFunc("/saveSienaFirm/", saveSienaFirmHandler)
	http.HandleFunc("/newSienaFirm/", newSienaFirmHandler)

	http.HandleFunc("/listSienaPortfolio/", listSienaPortfolioHandler)
	http.HandleFunc("/viewSienaPortfolio/", viewSienaPortfolioHandler)
	http.HandleFunc("/editSienaPortfolio/", editSienaPortfolioHandler)
	http.HandleFunc("/saveSienaPortfolio/", saveSienaPortfolioHandler)
	http.HandleFunc("/newSienaPortfolio/", newSienaPortfolioHandler)

	http.HandleFunc("/listSienaCentre/", listSienaCentreHandler)
	http.HandleFunc("/viewSienaCentre/", viewSienaCentreHandler)
	http.HandleFunc("/editSienaCentre/", editSienaCentreHandler)
	http.HandleFunc("/saveSienaCentre/", saveSienaCentreHandler)
	http.HandleFunc("/newSienaCentre/", newSienaCentreHandler)

	http.HandleFunc("/listSienaBook/", listSienaBookHandler)
	http.HandleFunc("/viewSienaBook/", viewSienaBookHandler)
	http.HandleFunc("/editSienaBook/", editSienaBookHandler)
	http.HandleFunc("/saveSienaBook/", saveSienaBookHandler)
	http.HandleFunc("/newSienaBook/", newSienaBookHandler)

	http.HandleFunc("/listSienaBroker/", listSienaBrokerHandler)
	http.HandleFunc("/viewSienaBroker/", viewSienaBrokerHandler)
	http.HandleFunc("/editSienaBroker/", editSienaBrokerHandler)
	http.HandleFunc("/saveSienaBroker/", saveSienaBrokerHandler)
	http.HandleFunc("/newSienaBroker/", newSienaBrokerHandler)

	http.HandleFunc("/listSienaAccount/", listSienaAccountHandler)
	http.HandleFunc("/viewSienaAccount/", viewSienaAccountHandler)

	http.HandleFunc("/shutdown/", shutdownHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("")
	fmt.Println("READY STEADY GO!!!")
	fmt.Println("URL        :", "http://localhost:"+wctProperties["port"])

	httpPort := ":" + wctProperties["port"]
	http.ListenAndServe(httpPort, nil)

}
