package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	application "github.com/mt1976/mwt-go-dev/application"
	"github.com/mt1976/mwt-go-dev/core"
	globals "github.com/mt1976/mwt-go-dev/core"
	scheduler "github.com/mt1976/mwt-go-dev/jobs"
	siena "github.com/mt1976/mwt-go-dev/siena"
)

func main() {

	//ascii := figlet4go.NewAsciiRender()

	//wctProperties := application.GetProperties(APPCONFIG)

	// The underscore would be an error
	//	renderStr, _ := ascii.Render(wctProperties["appname"])
	tmpHostname, _ := os.Hostname()
	line := strings.Repeat("-", 100)
	log.Println(line)

	globals.LOG_header("Initialising ...")

	globals.Initialise()

	globals.LOG_success("Initialised")
	//	log.Println(line)
	globals.LOG_header("Caching ...")
	//application.InitialiseCache()
	globals.LOG_success("Cache Refreshed")
	//log.Println(line)

	globals.LOG_header("Scheduling Jobs")
	//log.Println("TEST>")
	//scheduler.RunJobCob("TEST")
	//log.Println("<TEST")
	scheduler.Start()
	globals.LOG_success("Jobs Scheduled")

	globals.LOG_header("Starting Handlers")
	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	mux.HandleFunc("/", core.LoginHandler)
	mux.HandleFunc("/login", core.ValidateLoginHandler)
	mux.HandleFunc("/logout", core.LogoutHandler)
	mux.HandleFunc("/home", core.HomePageHandler)
	//mux.HandleFunc("/srvServiceCatalog", application.ServiceCatalogHandler)
	mux.HandleFunc("/favicon.ico", application.FaviconHandler)
	mux.HandleFunc("/favicon-32x32.png", application.Favicon32Handler)
	mux.HandleFunc("/site.webmanifest", application.FaviconManifestHandler)
	mux.HandleFunc("/favicon-16x16.png", application.Favicon16Handler)
	mux.HandleFunc("/browserconfig.xml", application.FaviconBrowserConfigHandler)
	// mux.HandleFunc("/listResponses/", application.ListResponsesHandler)
	// mux.HandleFunc("/previewRequest/", application.PreviewRequestHandler)
	// mux.HandleFunc("/executeRequest/", application.ExecuteRequestHandler)
	// mux.HandleFunc("/viewResponse/", application.ViewResponseHandler)
	// mux.HandleFunc("/deleteResponse/", application.DeleteResponseHandler)
	mux.HandleFunc("/clearQueues/", clearQueuesHandler)
	// mux.HandleFunc("/viewSrvEnvironment/", application.ViewSrvEnvironmentHandler)
	// mux.HandleFunc("/listSrvConfiguration/", application.ListSrvConfigurationHandler)
	// mux.HandleFunc("/viewSrvConfiguration/", application.ViewSrvConfigurationHandler)
	// mux.HandleFunc("/editSrvConfiguration/", application.EditSrvConfigurationHandler)
	// mux.HandleFunc("/saveSrvConfiguration/", application.SaveSrvConfigurationHandler)
	// mux.HandleFunc("/viewAppConfiguration/", application.ViewAppConfigurationHandler)

	mux.HandleFunc("/listSvcDataMap/", siena.ListSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMap/", siena.ViewSvcDataMapHandler)
	mux.HandleFunc("/editSvcDataMap/", siena.EditSvcDataMapHandler)
	mux.HandleFunc("/saveSvcDataMap/", siena.SaveSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMapXML/", siena.ViewSvcDataMapXMLHandler)
	mux.HandleFunc("/editSvcDataMapXML/", siena.EditSvcDataMapXMLHandler)
	mux.HandleFunc("/saveSvcDataMapXML/", siena.SaveSvcDataMapXMLHandler)
	mux.HandleFunc("/newSvcDataMap/", siena.NewSvcDataMapHandler)
	mux.HandleFunc("/genSvcDataMap/", siena.GenSvcDataMapHandler)
	mux.HandleFunc("/deleteSvcDataMap/", siena.DeleteSvcDataMapHandler)
	mux.HandleFunc("/editDataMapColumns/", application.ListLoaderMapStoreHandler)
	mux.HandleFunc("/editLoaderMapStore/", application.EditLoaderMapStoreHandler)
	mux.HandleFunc("/saveLoaderMapStore/", application.SaveLoaderMapStoreHandler)
	mux.HandleFunc("/newLoaderMapStore/", application.NewLoaderMapStoreHandler)
	mux.HandleFunc("/runLoader/", siena.RunDataLoaderHandler)

	siena.Country_MUX(*mux)
	siena.Sector_MUX(*mux)
	siena.Firm_MUX(*mux)
	siena.Portfolio_MUX(*mux)

	mux.HandleFunc("/listSienaCentre/", siena.ListSienaCentreHandler)
	mux.HandleFunc("/viewSienaCentre/", siena.ViewSienaCentreHandler)
	mux.HandleFunc("/editSienaCentre/", siena.EditSienaCentreHandler)
	mux.HandleFunc("/saveSienaCentre/", siena.SaveSienaCentreHandler)
	mux.HandleFunc("/newSienaCentre/", siena.NewSienaCentreHandler)

	mux.HandleFunc("/listSienaBook/", siena.ListSienaBookHandler)
	mux.HandleFunc("/viewSienaBook/", siena.ViewSienaBookHandler)
	mux.HandleFunc("/editSienaBook/", siena.EditSienaBookHandler)
	mux.HandleFunc("/saveSienaBook/", siena.SaveSienaBookHandler)
	mux.HandleFunc("/newSienaBook/", siena.NewSienaBookHandler)

	mux.HandleFunc("/listSienaBroker/", siena.ListSienaBrokerHandler)
	mux.HandleFunc("/viewSienaBroker/", siena.ViewSienaBrokerHandler)
	mux.HandleFunc("/editSienaBroker/", siena.EditSienaBrokerHandler)
	mux.HandleFunc("/saveSienaBroker/", siena.SaveSienaBrokerHandler)
	mux.HandleFunc("/newSienaBroker/", siena.NewSienaBrokerHandler)

	mux.HandleFunc("/listSienaAccount/", siena.ListSienaAccountHandler)
	mux.HandleFunc("/viewSienaAccount/", siena.ViewSienaAccountHandler)

	mux.HandleFunc("/listSienaCurrency/", siena.ListSienaCurrencyHandler)
	mux.HandleFunc("/listSienaCurrencyPair/", siena.ListSienaCurrencyPairHandler)

	mux.HandleFunc("/listSienaMandatedUser/", siena.ListSienaMandatedUserHandler)
	mux.HandleFunc("/viewSienaMandatedUser/", siena.ViewSienaMandatedUserHandler)
	mux.HandleFunc("/editSienaMandatedUser/", siena.EditSienaMandatedUserHandler)
	mux.HandleFunc("/saveSienaMandatedUser/", siena.SaveSienaMandatedUserHandler)
	mux.HandleFunc("/newSienaMandatedUser/", siena.NewSienaMandatedUserHandler)

	mux.HandleFunc("/listSienaCounterparty/", siena.ListSienaCounterpartyHandler)
	mux.HandleFunc("/viewSienaCounterparty/", siena.ViewSienaCounterpartyHandler)
	mux.HandleFunc("/editSienaCounterparty/", siena.EditSienaCounterpartyHandler)
	mux.HandleFunc("/saveSienaCounterparty/", siena.SaveSienaCounterpartyHandler)
	mux.HandleFunc("/newSienaCounterparty/", siena.NewSienaCounterpartyHandler)

	mux.HandleFunc("/editSienaCounterpartyAddress/", siena.EditSienaCounterpartyAddressHandler)
	mux.HandleFunc("/saveSienaCounterpartyAddress/", siena.SaveSienaCounterpartyAddressHandler)
	mux.HandleFunc("/editSienaCounterpartyExtensions/", siena.EditSienaCounterpartyExtensionsHandler)
	mux.HandleFunc("/saveSienaCounterpartyExtensions/", siena.SaveSienaCounterpartyExtensionsHandler)

	mux.HandleFunc("/listSienaCounterpartyPayee/", siena.ListSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/viewSienaCounterpartyPayee/", siena.ViewSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/editSienaCounterpartyPayee/", siena.EditSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/saveSienaCounterpartyPayee/", siena.SaveSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/newSienaCounterpartyPayee/", siena.NewSienaCounterpartyPayeeHandler)

	mux.HandleFunc("/listSienaCounterpartyImportID/", siena.ListSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/viewSienaCounterpartyImportID/", siena.ViewSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/editSienaCounterpartyImportID/", siena.EditSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/saveSienaCounterpartyImportID/", siena.SaveSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/newSienaCounterpartyImportID/", siena.NewSienaCounterpartyImportIDHandler)

	mux.HandleFunc("/listSienaDealList/", siena.ListSienaDealListHandler)
	mux.HandleFunc("/viewSienaDealList/", siena.ViewSienaDealListHandler)
	mux.HandleFunc("/listSienaAccountLadder/", siena.ListSienaAccountLadderHandler)
	mux.HandleFunc("/listSienaAccountTransactions/", siena.ListSienaAccountTransactionsHandler)
	//mux.HandleFunc("/saveSienaDealList/", siena.SaveSienaDealListHandler)
	//mux.HandleFunc("/newSienaDealList/", siena.NewSienaDealListHandler)

	mux.HandleFunc("/listSienaCounterpartyGroup/", siena.ListSienaCounterpartyGroupHandler)
	mux.HandleFunc("/viewSienaCounterpartyGroup/", siena.ViewSienaCounterpartyGroupHandler)
	mux.HandleFunc("/editSienaCounterpartyGroup/", siena.EditSienaCounterpartyGroupHandler)
	mux.HandleFunc("/saveSienaCounterpartyGroup/", siena.SaveSienaCounterpartyGroupHandler)
	mux.HandleFunc("/newSienaCounterpartyGroup/", siena.NewSienaCounterpartyGroupHandler)

	mux.HandleFunc("/dashboard/", siena.SienaDashboardHandler)

	mux.HandleFunc("/listCredentialsStore/", application.ListCredentialsStoreHandler)
	mux.HandleFunc("/viewCredentialsStore/", application.ViewCredentialStoreHandler)
	mux.HandleFunc("/editCredentialsStore/", application.EditCredentialStoreHandler)
	mux.HandleFunc("/deleteCredentialsStore/", application.DeleteCredentialStoreHandler)
	mux.HandleFunc("/saveCredentialsStore/", application.SaveCredentialStoreHandler)
	mux.HandleFunc("/newCredentialsStore/", application.NewCredentialStoreHandler)
	mux.HandleFunc("/banCredentialsStore/", application.BanCredentialStoreHandler)
	mux.HandleFunc("/activateCredentialsStore/", application.ActivateCredentialStoreHandler)

	mux.HandleFunc("/listMessageStore/", application.ListMessageStoreHandler)
	mux.HandleFunc("/viewMessageStore/", application.ViewMessageStoreHandler)
	mux.HandleFunc("/editMessageStore/", application.EditMessageStoreHandler)
	mux.HandleFunc("/saveMessageStore/", application.SaveMessageStoreHandler)

	mux.HandleFunc("/listTranslationStore/", application.ListTranslationStoreHandler)
	mux.HandleFunc("/viewTranslationStore/", application.ViewTranslationStoreHandler)
	mux.HandleFunc("/editTranslationStore/", application.EditTranslationStoreHandler)
	mux.HandleFunc("/saveTranslationStore/", application.SaveTranslationStoreHandler)

	mux.HandleFunc("/listScheduleStore/", application.ListScheduleStoreHandler)
	mux.HandleFunc("/viewScheduleStore/", application.ViewScheduleStoreHandler)

	mux.HandleFunc("/listSessionStore/", application.ListSessionStoreHandler)
	mux.HandleFunc("/deleteSessionStore/", application.DeleteSessionStoreHandler)

	mux.HandleFunc("/listSystemStore/", application.ListSystemStoreHandler)
	mux.HandleFunc("/viewSystemStore/", application.ViewSystemStoreHandler)
	mux.HandleFunc("/editSystemStore/", application.EditSystemStoreHandler)
	mux.HandleFunc("/deleteSystemStore/", application.DeleteSystemStoreHandler)
	mux.HandleFunc("/saveSystemStore/", application.SaveSystemStoreHandler)
	mux.HandleFunc("/newSystemStore/", application.NewSystemStoreHandler)

	mux.HandleFunc("/listFundsCheck/", application.ListFundsCheckHandler)
	mux.HandleFunc("/viewFundsCheck/", application.ViewFundsCheckHandler)
	mux.HandleFunc("/actionFundsCheck/", application.ActionFundsCheckHandler)
	mux.HandleFunc("/submitFundsCheck/", application.SubmitFundsCheckHandler)

	mux.HandleFunc("/injectSQLViews/", application.SQLInjectionHandler)

	//	mux.HandleFunc("/refreshCache/", application.RefreshCacheHandler)
	application.DataCache_MUX(*mux)

	mux.HandleFunc("/listGiltsDataStore/", application.ListLSEGiltsDataStoreHandler)
	mux.HandleFunc("/viewLSEGiltsDataStore/", application.ViewLSEGiltsDataStoreHandler)
	mux.HandleFunc("/selectLSEGiltsDataStore/", application.SelectLSEGiltsDataStoreHandler)
	mux.HandleFunc("/deselectLSEGiltsDataStore/", application.DeselectLSEGiltsDataStoreHandler)

	mux.HandleFunc("/shutdown/", shutdownHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	globals.LOG_success("Handlers Started")
	log.Println(line)
	globals.LOG_header("Application Information")
	globals.LOG_header("Application")
	globals.LOG_info("Name", globals.ApplicationProperties["appname"])
	globals.LOG_info("Host Name", tmpHostname)
	globals.LOG_info("Server Release", fmt.Sprintf("%s [r%s-%s]", globals.ApplicationProperties["releaseid"], globals.ApplicationProperties["releaselevel"], globals.ApplicationProperties["releasenumber"]))
	globals.LOG_info("Server Date", time.Now().Format(globals.DATEFORMATUSER))
	if globals.IsChildInstance {
		globals.LOG_info("Server Mode", "Primary System")
	} else {
		globals.LOG_info("Server Mode", "Secondary System")
	}
	globals.LOG_info("Licence", globals.ApplicationProperties["licname"])
	globals.LOG_info("Lic URL", globals.ApplicationProperties["liclink"])

	globals.LOG_header("Application Database (MSSQL)")
	globals.LOG_info("Server", globals.ApplicationPropertiesDB["server"])
	globals.LOG_info("Database", globals.ApplicationPropertiesDB["database"])
	globals.LOG_info("Schema", globals.ApplicationPropertiesDB["schema"])
	globals.LOG_info("Parent Schema", globals.ApplicationPropertiesDB["parentschema"])

	globals.LOG_header("Siena")
	_, tempDate, _ := siena.GetBusinessDate(globals.SienaDB)
	globals.SienaSystemDate = tempDate
	globals.LOG_info("System", globals.SienaProperties["name"])
	globals.LOG_info("System Date", globals.SienaSystemDate.Internal.Format(globals.DATEFORMATUSER))

	globals.LOG_header("Siena Database (MSSQL)")
	globals.LOG_info("Server", globals.SienaPropertiesDB["server"])
	globals.LOG_info("Database", globals.SienaPropertiesDB["database"])
	globals.LOG_info("Schema", globals.SienaPropertiesDB["schema"])
	globals.LOG_info("Parent Schema", globals.SienaPropertiesDB["parentschema"])

	globals.LOG_header("Siena Connectivity")
	globals.LOG_info("TXNs Delivery", globals.SienaProperties["transactional_in"])
	globals.LOG_info("TXNs Response", globals.SienaProperties["transactional_out"])
	globals.LOG_info("Static Delivery", globals.SienaProperties["static_in"])
	globals.LOG_info("Static Response", globals.SienaProperties["static_out"])
	globals.LOG_info("Funds Check Request", globals.SienaProperties["funds_out"])
	globals.LOG_info("Funds Check Response", globals.SienaProperties["funds_in"])
	globals.LOG_info("Rates & Prices Delivery", globals.SienaProperties["rates_in"])

	globals.LOG_header("Sessions")
	globals.LOG_info("Session Life", globals.ApplicationProperties["sessionlife"])

	//scheduler.RunJobLSE("")
	//scheduler.RunJobFII("")

	globals.LOG_header("READY STEADY GO!!!")
	log.Println("URI            :", globals.ColorPurple+"http://localhost:"+globals.ApplicationProperties["port"]+globals.ColorReset)
	log.Println(line)

	httpPort := ":" + globals.ApplicationProperties["port"]
	//http.ListenAndServe(httpPort, nil)

	//scheduler.RunJobFII("")
	//s, _ := application.GLIEF_leiLookup("GB00BL68HJ26")
	//info("LEI", s)
	// Wrap your handlers with the LoadAndSave() middleware.
	http.ListenAndServe(httpPort, globals.SessionManager.LoadAndSave(mux))

	globals.Log_uptime()

}

//// TODO: migrage the following three functions to appsupport
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	//	wctProperties := application.GetProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//requestID := uuid.New()http.ListenAndServe(":8080", nil)

	log.Println("Servicing :", inUTL)
	//requestMessage := application.BuildRequestMessage(requestID.String(), "SHUTDOWN", line, line, line, wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	//application.SendRequest(requestMessage, requestID.String(), globals.ApplicationProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: globals.ApplicationProperties["port"], Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}
func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	//wctProperties := application.GetProperties(APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := core.RemoveContents(globals.ApplicationProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println("recPath", wctProperties["receivepath"])
	err2 := core.RemoveContents(globals.ApplicationProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])
	err3 := core.RemoveContents(globals.ApplicationProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}
	core.HomePageHandler(w, r)
}

// func clearResponsesHandler(w http.ResponseWriter, r *http.Request) {
// 	//var propertiesFileName = "config/properties.cfg"
// 	//	wctProperties := application.GetProperties(globals.APPCONFIG)
// 	//	tmpl := "viewResponse"
// 	inUTL := r.URL.Path
// 	//requestID := uuid.New()
// 	log.Println("Servicing :", inUTL)
// 	application.RemoveContents(globals.ApplicationProperties["receivepath"])
// 	application.HomePageHandler(w, r)
// }

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	globals.SessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := globals.SessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}
