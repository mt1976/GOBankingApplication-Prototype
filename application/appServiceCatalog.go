package application

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

//srvCatalogPage ...
type srvCatalogPage struct {
	UserMenu       []AppMenuItem
	UserRole       string
	UserNavi       string
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
	PageTitle      string
}

func ServiceCatalogHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {
		tmpl := "listSrvServiceCatalog"
		log.Println("Servicing :", inUTL)

		w.Header().Set("Content-Type", "text/html")

		title := globals.ApplicationProperties["appname"]

		//p, _ := loadsrvCatalogPage(title)

		noResp, _, respList := GetResponsesList(globals.ApplicationProperties, "json", w)

		noServices, servicesList, serviceCatalog := GetServices(globals.ApplicationProperties, "json", r)

		p := srvCatalogPage{Title: title,
			Body:           "",
			RequestPath:    globals.ApplicationProperties["deliverpath"],
			ResponsePath:   globals.ApplicationProperties["receivepath"],
			ProcessedPath:  globals.ApplicationProperties["processedpath"],
			NoResponses:    noResp,
			Responses:      respList,
			NoServices:     noServices,
			Services:       servicesList,
			ServiceCatalog: serviceCatalog,
			Description:    "A description of the srvCatalogPage.",
			ResponseType:   globals.ApplicationProperties["responseformat"],
			UserMenu:       GetAppMenuData(globals.UserRole),
			UserRole:       globals.UserRole,
			UserNavi:       globals.UserNavi,
			PageTitle:      "Service Catalog",
		}

		//	fmt.Println("serviceCatalog", serviceCatalog)
		//	fmt.Println("srvCatalogPage=", p.ServiceCatalog)
		fmt.Println("menu=", p.UserMenu)

		t, _ := template.ParseFiles(GetTemplateID(tmpl, globals.UserRole))
		t.Execute(w, p)
	}
}
